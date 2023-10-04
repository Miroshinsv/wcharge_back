package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) newUserRoutes() {
	s.router.HandleFunc("/api/user/all", s.GetUsersWebAPI).Methods(http.MethodGet)                     // Получить список всех пользователей
	s.router.HandleFunc("/api/user/get/{id:[0-9]+}", s.GetUserWebAPI).Methods(http.MethodGet)          // Получить информацию о конкретном пользователе
	s.router.HandleFunc("/api/user/create", s.CreateUserWebAPI()).Methods(http.MethodPost)             // Создать нового пользователя
	s.router.HandleFunc("/api/user/update/{id:[0-9]+}", s.UpdateUserWebAPI).Methods(http.MethodPut)    // Обновить информацию о пользователе
	s.router.HandleFunc("/api/user/delete/{id:[0-9]+}", s.DeleteUserWebAPI).Methods(http.MethodDelete) // Удалить пользователя
}

func (s *server) GetUsersWebAPI(w http.ResponseWriter, r *http.Request) {
	users, err := s.useCase.GetUsers()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error()))
		return
	}
	for i, _ := range users {
		users[i].Sanitize()
	}
	s.respond(w, r, http.StatusOK, users)
}

func (s *server) GetUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}
	user, err := s.useCase.GetUser(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error()))
		return
	}
	user.Sanitize()
	s.respond(w, r, http.StatusOK, user)
}

func (s *server) CreateUserWebAPI() http.HandlerFunc {
	type request struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - "+err.Error()))
			return
		}
		u := entity.User{
			Username: req.UserName,
			Email:    req.Email,
			Password: req.Password,
		}

		err := s.useCase.CreateUser(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - CreateUserWebAPI - usecase.User.CreateUser - "+err.Error()))
			return
		}

		s.respond(w, r, http.StatusOK, "")
	}
}

func (s *server) UpdateUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}
	u := &entity.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - "+err.Error()))
		return
	}

	err = s.useCase.UpdateUser(id, *u)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - UpdateUserWebAPI - usecase.User.UpdateUser - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, "")
}

func (s *server) DeleteUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - DeleteUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}
	err = s.useCase.DeleteUser(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("\"error - GetUsersWebAPI - usecase.User.DeleteUser - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, "")
}
