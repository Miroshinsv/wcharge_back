package v1

import (
	"encoding/json"
	"errors"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var sessionName = "userSession"

func (s *server) newUserRoutes() {
	s.router.HandleFunc("/api/user/login", s.LoginWebAPI()).Methods("POST")
	s.router.HandleFunc("/api/user/all", s.GetUsersWebAPI).Methods("GET")                              // Получить список всех пользователей
	s.router.HandleFunc("/api/user/get/{id:[0-9]+}", s.GetUserWebAPI).Methods("GET")                   // Получить информацию о конкретном пользователе
	s.router.HandleFunc("/api/user/create", s.CreateUserWebAPI()).Methods("POST")                      // Создать нового пользователя
	s.router.HandleFunc("/api/user/update/{id:[0-9]+}", s.UpdateUserWebAPI).Methods("PUT")             // Обновить информацию о пользователе
	s.router.HandleFunc("/api/user/delete/{id:[0-9]+}", s.DeleteUserWebAPI).Methods(http.MethodDelete) // Удалить пользователя
}

func RequestToJSONUser(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "content type is not application/json", http.StatusUnsupportedMediaType)
		return entity.User{}, errors.New("content type is not application/json")
	}
	var u entity.User
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&u)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "bad request - wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "bad request "+err.Error(), http.StatusBadRequest)
		}
		return entity.User{}, errors.New("bad request - entity is not User")
	}

	return u, nil
}

func (s *server) LoginWebAPI() http.HandlerFunc {
	type request struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errorResponse(w, "userRoutes - Login - "+err.Error(), http.StatusBadRequest)
			return
		}
		u, err := s.useCase.GetUserByName(req.UserName)
		if err != nil || !u.ComparePassword(req.Password) {
			errorResponse(w, "userRoutes - Login - invalid username or password", http.StatusInternalServerError)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			errorResponse(w, "server - Login - s.sessionStore.Get -"+err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user_id"] = u.ID
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			errorResponse(w, "server - Login - s.sessionStore.Save -"+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) GetUsersWebAPI(w http.ResponseWriter, r *http.Request) {
	users, err := s.useCase.GetUsers()
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for i, _ := range users {
		users[i].Sanitize()
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - Encode(users) - "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) GetUserWebAPI(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse(w, "error - GetUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error(), 0)
		return
	}
	user, err := s.useCase.GetUser(id)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user.Sanitize()
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		errorResponse(w, "error - GetUserWebAPI - Encode(user) - "+err.Error(), http.StatusInternalServerError)
		return
	}
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
			errorResponse(w, "userRoutes - Login - "+err.Error(), http.StatusBadRequest)
			return
		}
		u := entity.User{
			Username: req.UserName,
			Email:    req.Email,
			Password: req.Password,
		}
		err := s.useCase.CreateUser(u)
		if err != nil {
			errorResponse(w, "error - CreateUserWebAPI - usecase.User.CreateUser - "+err.Error(), 0)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) UpdateUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse(w, "error - GetUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error(), 0)
		return
	}

	u, err := RequestToJSONUser(w, r)
	if err != nil {
		return
	}

	err = s.useCase.UpdateUser(id, u)
	if err != nil {
		errorResponse(w, "error - UpdateUserWebAPI - usecase.User.UpdateUser - "+err.Error(), 0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *server) DeleteUserWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse(w, "error - DeleteUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error(), 0)
		return
	}
	user := s.useCase.DeleteUser(id)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.DeleteUser - "+err.Error(), 0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - json.NewEncoder(w).Encode(user) - "+err.Error(), 0)
		return
	}
}
