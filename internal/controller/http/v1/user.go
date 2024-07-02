package v1

import (
	"encoding/json"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (s *server) newUserRoutes() {
	// prefix /api
	s.apiRouter.HandleFunc("/users", s.GetUsers).Methods(http.MethodGet)    // Получить список всех пользователей
	s.apiRouter.HandleFunc("/users", s.CreateUser).Methods(http.MethodPost) // Создать нового пользователя

	s.apiRouter.HandleFunc("/users/{id:[0-9]+}", s.GetUser).Methods(http.MethodGet)       // Получить информацию о конкретном пользователе
	s.apiRouter.HandleFunc("/users/{id:[0-9]+}", s.UpdateUser).Methods(http.MethodPut)    // Обновить информацию о пользователе
	s.apiRouter.HandleFunc("/users/{id:[0-9]+}", s.DeleteUser).Methods(http.MethodDelete) // Удалить пользователя
	s.apiRouter.HandleFunc("/users/{id:[0-9]+}/powerbanks", s.GetUserPowerbanks).Methods(http.MethodGet)
}

// GetUsers godoc
// @Summary 	 Get info about all users
// @Success      200  {array}  	entity.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.useCase.GetUsers()
	if err != nil {
		log.Printf("Error - GetUsers - usecase.User.GetUsers - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	for i, _ := range *users {
		(*users)[i].Sanitize()
	}
	s.respond(w, http.StatusOK, users)
}

// GetUser godoc
// @Summary 	 Get info about user
// @Param        userId   		path	int		true  	"User Id"
// @Success      200  {object}  entity.User
// @Failure      500  {object}  map[string]string
// @Router       /users/{userId} [get]
func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - GetUser - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	user, err := s.useCase.GetUser(id)
	if err != nil {
		log.Printf("Error - GetUsers - usecase.GetUsers - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	user.Sanitize()
	s.respond(w, http.StatusOK, user)
}

// CreateUser godoc
// @Summary 	 Create station
// @Param        UserName   	body   	string	true	"Username"
// @Param        Email   		body   	string	true  	"Email"
// @Param        Password   	body   	string	true  	"Password"
// @Success      200  {object}  entity.User
// @Failure      500  {object}  map[string]string
// @Router       /users [post]
func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - CreateUser - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	u := entity.User{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	}

	uu, err := s.useCase.CreateUser(u)
	if err != nil {
		log.Printf("Error - CreateUser - usecase.CreateUser - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, uu)
}

// UpdateUser godoc
// @Summary 	 Update station
// @Param        userId   		path	int		true  	"User Id"
// @Param        UserName   	body   	string	true	"Username"
// @Param        Email   		body   	string	true  	"Email"
// @Param        Password   	body   	string	true  	"Password"
// @Param        Address   		body   	int		true  	"Address Id"
// @Success      200  {object}  entity.User
// @Failure      500  {object}  map[string]string
// @Router       /users/{userId} [put]
func (s *server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Address  int    `json:"address"`
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - UpdateUser - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - UpdateUser - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	u := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	uu, err := s.useCase.UpdateUser(*u, id)
	if err != nil {
		log.Printf("Error - UpdateUser - usecase.UpdateUser - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, uu)
}

// DeleteUser godoc
// @Summary 	 Delete station
// @Param        userId   		path	int		true	"User Id"
// @Success      200  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /users/{userId} [delete]
func (s *server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - DeleteUser -  %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	err = s.useCase.DeleteUser(id)
	if err != nil {
		log.Printf("Error - DeleteUser - usecase.DeleteUser -  %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, nil)
}

// GetUserPowerbanks godoc
// @Summary 	 Get info about all powerbanks of users
// @Param        userId   		path	int		true	"User Id"
// @Success      200  {array}  	entity.Powerbank
// @Failure      500  {object}  map[string]string
// @Router       /users/{userId}/powerbanks [get]
func (s *server) GetUserPowerbanks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		log.Printf("Error - GetUserPowerbanks - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	powerbanks, err := s.useCase.GetUserPowerbanks(userId)
	if err != nil {
		log.Printf("Error - GetUserPowerbanks - usecase.GetUserPowerbanks - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, powerbanks)
}
