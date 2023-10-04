// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"encoding/json"
	"fmt"
	// Swagger docs.
	_ "github.com/Miroshinsv/wcharge_back/docs"
	"net/http"
)

// NewHttpRouter -.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func (s *server) NewHttpRouter() {
	s.router.HandleFunc("/", home)
	s.router.HandleFunc("/api/login", s.LoginWebAPI()).Methods(http.MethodPost)

	s.newUserRoutes()
	s.newStationRoutes()
	s.newPowerbankRoutes(s.router, s.useCase, s.logger)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func (s *server) LoginWebAPI() http.HandlerFunc {
	type request struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - "+err.Error()))
			return
		}
		u, err := s.useCase.GetUserByName(req.UserName)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("userRoutes - Login - invalid username or password"))
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("erver - Login - s.sessionStore.Get -"+err.Error()))
			return
		}

		session.Values["user_id"] = u.ID
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("server - Login - s.sessionStore.Save -"+err.Error()))
			return
		}
	}
}
