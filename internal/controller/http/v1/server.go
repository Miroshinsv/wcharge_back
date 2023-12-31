package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/config"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/httpserver"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	sessionName        = "wcharge" // в конфиге, а там из env
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	apiRouter    *mux.Router
	useCase      *usecase.UseCase
	sessionStore sessions.Store
	logger       logger.Interface
}

func NewServer(u *usecase.UseCase, sessionStore sessions.Store, l logger.Interface) *server {
	s := &server{
		router:       mux.NewRouter(),
		useCase:      u,
		sessionStore: sessionStore,
		logger:       l,
	}
	return s
}

func Start(cfg *config.Config, u *usecase.UseCase, l logger.Interface) {
	sessionsStore := sessions.NewCookieStore([]byte(cfg.SessionHttpKey))

	srv := NewServer(u, sessionsStore, l)
	srv.NewHttpRouter()
	httpServer := httpserver.New(srv.router, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil && code != http.StatusInternalServerError { // анти вечный цикл
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}
}
