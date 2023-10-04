package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) newStationRoutes() {
	s.router.HandleFunc("/api/station/all", s.GetStationWebAPI).Methods(http.MethodGet)                      // Получить список всех пользователей
	s.router.HandleFunc("/api/station/get/{id:[0-9]+}", s.GetStationWebAPI).Methods(http.MethodGet)          // Получить информацию о конкретном пользователе
	s.router.HandleFunc("/api/station/create", s.CreateStationWebAPI()).Methods(http.MethodPost)             // Создать нового пользователя
	s.router.HandleFunc("/api/station/update/{id:[0-9]+}", s.UpdateStationWebAPI()).Methods(http.MethodPut)  // Обновить информацию о пользователе
	s.router.HandleFunc("/api/station/delete/{id:[0-9]+}", s.DeleteStationWebAPI).Methods(http.MethodDelete) // Удалить пользователя
}

func (s *server) GetStationsWebAPI(w http.ResponseWriter, r *http.Request) {
	stations, err := s.useCase.GetStations()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, stations)
}

func (s *server) GetStationWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		return
	}
	st, err := s.useCase.GetStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, st)
}

func (s *server) CreateStationWebAPI() http.HandlerFunc {
	type request struct {
		SerialNumber string `json:"serial_number"`
		AddressId    int    `json:"address"`
		Capacity     int    `json:"capacity"`
		FreeCapacity int    `json:"free_capacity"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - "+err.Error()))
			return
		}
		st := entity.Station{
			SerialNumber: req.SerialNumber,
			AddressId:    req.AddressId,
			Capacity:     req.Capacity,
			FreeCapacity: req.FreeCapacity,
		}

		err := s.useCase.CreateStation(st)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - CreateStationWebAPI - usecase.CreateStation - "+err.Error()))
			return
		}

		s.respond(w, r, http.StatusOK, "")
	}
}

func (s *server) UpdateStationWebAPI() http.HandlerFunc {
	type request struct {
		AddressId    int `json:"address"`
		FreeCapacity int `json:"free_capacity"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - UpdateStationWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("userRoutes - Login - "+err.Error()))
			return
		}
		st := entity.Station{
			ID:           id,
			AddressId:    req.AddressId,
			FreeCapacity: req.FreeCapacity,
		}

		err = s.useCase.UpdateStation(id, st)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - CreateStationWebAPI - usecase.CreateStation - "+err.Error()))
			return
		}

		s.respond(w, r, http.StatusOK, "")
	}
}

func (s *server) DeleteStationWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - DeleteStationWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}

	err = s.useCase.DeleteStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - DeleteStationWebAPI - usecase.User.GetUsers - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, "")
}
