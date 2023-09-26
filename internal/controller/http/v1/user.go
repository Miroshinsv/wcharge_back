package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
)

type userRoutes struct {
	u usecase.UserAPI
	l logger.Interface
}

func newUserRoutes(router *mux.Router, u usecase.UserAPI, l logger.Interface) {

	ur := &userRoutes{u, l}
	router.HandleFunc("/api/user/all", ur.GetUsersWebAPI).Methods("GET")              // Получить список всех пользователей
	router.HandleFunc("/api/user/get/{id}", ur.GetUserWebAPI).Methods("GET")          // Получить информацию о конкретном пользователе
	router.HandleFunc("/api/user/create", ur.CreateUserWebAPI).Methods("POST")        // Создать нового пользователя
	router.HandleFunc("/api/user/update/{id}", ur.UpdateUserWebAPI).Methods("PUT")    // Обновить информацию о пользователе
	router.HandleFunc("/api/user/delete/{id}", ur.DeleteUserWebAPI).Methods("DELETE") // Удалить пользователя
}

type GetUsersResponse struct {
	Users []entity.User `json:"history"`
}

type UpdateRequest struct {
	ID       int          `json:"id"       binding:"required"`
	Username string       `json:"username" binding:"required"`
	Email    string       `json:"email"    binding:"required"`
	Role     *entity.Role `json:"role"     binding:"required"`
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

func (ur *userRoutes) GetUsersWebAPI(w http.ResponseWriter, r *http.Request) {
	users, err := ur.u.GetUsers()
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, u := range users {
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (ur *userRoutes) GetUserWebAPI(w http.ResponseWriter, r *http.Request) {
	u, err := RequestToJSONUser(w, r)
	if err != nil {
		return
	}
	user, err := ur.u.GetUser(u)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (ur *userRoutes) CreateUserWebAPI(w http.ResponseWriter, r *http.Request) {
	u, err := RequestToJSONUser(w, r)
	if err != nil {
		return
	}

	err = ur.u.CreateUser(u)
	if err != nil {
		errorResponse(w, "error - CreateUserWebAPI - usecase.User.CreateUser - "+err.Error(), 0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *userRoutes) UpdateUserWebAPI(w http.ResponseWriter, r *http.Request) {
	u, err := RequestToJSONUser(w, r)
	if err != nil {
		return
	}

	err = ur.u.UpdateUser(u)
	if err != nil {
		errorResponse(w, "error - UpdateUserWebAPI - usecase.User.UpdateUser - "+err.Error(), 0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *userRoutes) DeleteUserWebAPI(w http.ResponseWriter, r *http.Request) {
	u, err := RequestToJSONUser(w, r)
	if err != nil {
		return
	}

	user := ur.u.DeleteUser(u)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), 0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}
