package v1

import (
	"encoding/json"
	"errors"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

type stationRoutes struct {
	s usecase.StationAPI
	l logger.Interface
}

func newStationRoutes(router *mux.Router, s usecase.StationAPI, l logger.Interface) {

	sr := &stationRoutes{s, l}
	router.HandleFunc("/api/station/all", sr.GetStationWebAPI).Methods("GET")               // Получить список всех пользователей
	router.HandleFunc("/api/station/get/{id}", sr.GetStationWebAPI).Methods("GET")          // Получить информацию о конкретном пользователе
	router.HandleFunc("/api/station/create", sr.CreateStationWebAPI).Methods("POST")        // Создать нового пользователя
	router.HandleFunc("/api/station/update/{id}", sr.UpdateStationWebAPI).Methods("PUT")    // Обновить информацию о пользователе
	router.HandleFunc("/api/station/delete/{id}", sr.DeleteStationWebAPI).Methods("DELETE") // Удалить пользователя
}

func RequestToJSONStation(w http.ResponseWriter, r *http.Request) (entity.Station, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return entity.Station{}, errors.New("Content Type is not application/json")
	}
	var s entity.Station
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&s)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return entity.Station{}, errors.New("Bad Request - entity is not User")
	}

	return s, nil
}

func (ur *stationRoutes) GetStationsWebAPI(w http.ResponseWriter, r *http.Request) {
	stations, err := ur.s.GetStations()
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, s := range stations {
		json.NewEncoder(w).Encode(s)
	}
}

func (ur *stationRoutes) GetStationWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONStation(w, r)
	/*id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}*/
	if err != nil {
		return
	}
	user, err := ur.s.GetStation(s)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (ur *stationRoutes) CreateStationWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONStation(w, r)
	if err != nil {
		return
	}

	err = ur.s.CreateStation(s)
	if err != nil {
		errorResponse(w, "error - CreateUserWebAPI - usecase.User.CreateUser - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *stationRoutes) UpdateStationWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONStation(w, r)
	if err != nil {
		return
	}

	err = ur.s.UpdateStation(s)
	if err != nil {
		errorResponse(w, "error - UpdateUserWebAPI - usecase.User.UpdateUser - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *stationRoutes) DeleteStationWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONStation(w, r)
	if err != nil {
		return
	}

	user := ur.s.DeleteStation(s)
	if err != nil {
		errorResponse(w, "error - GetUsersWebAPI - usecase.User.GetUsers - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}
