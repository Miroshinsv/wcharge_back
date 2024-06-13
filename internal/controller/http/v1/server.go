package v1

import (
	"encoding/json"
	"errors"
	"log"
	//"github.com/rs/zerolog/log"
	//"github.com/rs/zerolog/log"
	//"github.com/slukits/graylog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Miroshinsv/wcharge_back/config"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/httpserver"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	//logger       *graylog.Logger
}

func NewServer(u *usecase.UseCase, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		useCase:      u,
		sessionStore: sessionStore,
		//logger:       l,
	}
	return s
}

func Start(cfg *config.Config, u *usecase.UseCase) {

	sessionsStore := sessions.NewCookieStore([]byte(cfg.SessionHttpKey))

	srv := NewServer(u, sessionsStore)
	srv.NewHttpRouter()
	httpServer := httpserver.New(srv.router, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Printf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Printf("app - Run - httpServer.Shutdown: %w", err)
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
