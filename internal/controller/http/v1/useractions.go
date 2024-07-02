package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) newUserActionsRoutes() {
	s.router.HandleFunc("/login", s.Login).Methods(http.MethodPost)   // TODO
	s.router.HandleFunc("/logout", s.Logout).Methods(http.MethodPost) // TODO
}

func (s *server) Login(w http.ResponseWriter, r *http.Request) {
	type request struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - %w", err))
		return
	}
	u, err := s.useCase.GetUserByName(req.UserName)
	if err != nil {
		s.error(w, http.StatusInternalServerError, fmt.Errorf("login - GetUserByName - %w", err))
		return
	}

	if !u.ComparePassword(req.Password) {
		s.error(w, http.StatusBadRequest, errIncorrectEmailOrPassword)
		return
	}

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, http.StatusInternalServerError, fmt.Errorf("erver - Login - s.sessionStore.Get - %w", err))
		return
	}

	session.Values["user_id"] = u.ID
	err = s.sessionStore.Save(r, w, session)
	if err != nil {
		s.error(w, http.StatusInternalServerError, fmt.Errorf("server - Login - s.sessionStore.Save - %w", err))
		return
	}
}

func (s *server) Logout(w http.ResponseWriter, r *http.Request) {

}
