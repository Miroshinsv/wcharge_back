package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) newUserActionsRoutes() {
	s.router.HandleFunc("/login", s.LoginWebAPI()).Methods(http.MethodPost)   // TODO
	s.router.HandleFunc("/logout", s.LogoutWebAPI()).Methods(http.MethodPost) // TODO
}

func (s *server) LoginWebAPI() http.HandlerFunc {
	type request struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - %w", err))
			return
		}
		u, err := s.useCase.GetUserByName(req.UserName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("login - GetUserByName - %w", err))
			return
		}

		if !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusBadRequest, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("erver - Login - s.sessionStore.Get - %w", err))
			return
		}

		session.Values["user_id"] = u.ID
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("server - Login - s.sessionStore.Save - %w", err))
			return
		}
	}
}

func (s *server) LogoutWebAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
