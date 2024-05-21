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
	// prefix /api
	s.apiRouter.HandleFunc("/user", s.GetUsersWebAPI).Methods(http.MethodGet)      // Получить список всех пользователей
	s.apiRouter.HandleFunc("/user", s.CreateUserWebAPI()).Methods(http.MethodPost) // Создать нового пользователя // TODO

	s.apiRouter.HandleFunc("/user/{id:[0-9]+}", s.GetUserWebAPI).Methods(http.MethodGet)       // Получить информацию о конкретном пользователе
	s.apiRouter.HandleFunc("/user/{id:[0-9]+}", s.UpdateUserWebAPI()).Methods(http.MethodPut)  // Обновить информацию о пользователе // TODO
	s.apiRouter.HandleFunc("/user/{id:[0-9]+}", s.DeleteUserWebAPI).Methods(http.MethodDelete) // Удалить пользователя
	s.apiRouter.HandleFunc("/user/{id:[0-9]+}/powerbanks", s.GetUserPowerbanksWebAPI).Methods(http.MethodGet)
}

func (s *server) GetUsersWebAPI(w http.ResponseWriter, r *http.Request) {
	users, err := s.useCase.GetUsers()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetUsersWebAPI - usecase.User.GetUsers - %w", err))
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
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetUserWebAPI - %w", err))
		return
	}
	user, err := s.useCase.GetUser(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetUsersWebAPI - usecase.GetUsers - %w", err))
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
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreateUserWebAPI - %w", err))
			return
		}
		u := entity.User{
			Username: req.UserName,
			Email:    req.Email,
			Password: req.Password,
		}

		err := s.useCase.CreateUser(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("CreateUserWebAPI - usecase.CreateUser - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) UpdateUserWebAPI() http.HandlerFunc {
	type request struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		AddressID int    `json:"address_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateUserWebAPI - %w", err))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("UpdateUserWebAPI - %w", err))
			return
		}
		u := &entity.User{
			Username:  req.Username,
			Email:     req.Email,
			Phone:     req.Phone,
			AddressID: req.AddressID,
		}
		err = s.useCase.UpdateUser(id, *u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateUserWebAPI - usecase.UpdateUser - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) DeleteUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteUserWebAPI -  %w", err))
		return
	}
	err = s.useCase.DeleteUser(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteUserWebAPI - usecase.DeleteUser -  %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

func (s *server) GetUserPowerbanksWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetUserPowerbanksWebAPI - %w", err))
		return
	}

	powerbanks, err := s.useCase.GetUserPowerbanks(userId)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetUserPowerbanksWebAPI - usecase.GetUserPowerbanks - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, powerbanks)
}
