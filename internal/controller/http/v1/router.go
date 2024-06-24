// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"context"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//httpSwagger "github.com/swaggo/http-swagger"

	//"github.com/rs/zerolog/log"
	"log"

	// Swagger docs.
	//_ "github.com/Miroshinsv/wcharge_back/docs"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// NewHttpRouter -.
// @version     2.0
// @host        localhost:8080
// @BasePath    /api/v1
func (s *server) NewHttpRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(s.commonMiddleware)

	s.apiRouter = s.router.PathPrefix("/api/v1").Subrouter()

	// TODO
	//s.apiRouter.Use(s.authenticateUser)

	s.apiRouter.HandleFunc("/whoami", s.handleWhoAmI()).Methods(http.MethodGet)

	s.newUserRoutes()
	s.newStationRoutes()
	s.newPowerbankRoutes()
	s.newAddressRoutes()
	s.newUserActionsRoutes()
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(entity.User))
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		switch {
		case rw.code >= 500:
			log.Printf("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
		case rw.code >= 400:
			log.Printf("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
		default:
			log.Printf("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
		}
	})
}

func (s *server) commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			fmt.Printf("completed with %d %s\n", http.StatusInternalServerError, "s.sessionStore.Get")
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			fmt.Printf("completed with %d %s\n", http.StatusInternalServerError, "session.Values")
			return
		}

		u, err := s.useCase.GetUser(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}
