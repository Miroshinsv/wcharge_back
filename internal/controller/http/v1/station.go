package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
)

func (s *server) newStationRoutes() {
	s.apiRouter.HandleFunc("/station/all", s.GetStationsWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/station/get/{id:[0-9]+}", s.GetStationWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/station/create", s.CreateStationWebAPI()).Methods(http.MethodPost)
	s.apiRouter.HandleFunc("/station/update/{id:[0-9]+}", s.UpdateStationWebAPI()).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/station/delete/{id:[0-9]+}", s.DeleteStationWebAPI).Methods(http.MethodDelete)

	s.apiRouter.HandleFunc("/station/{id:[0-9]+}/get/all-powerbanks", s.GetAllPowerbanksInStation).Methods(http.MethodGet)

	s.apiRouter.HandleFunc(
		"/station/{station_id:[0-9]+}/take-powerbank/{powerbank_id:[0-9]+}",
		s.TakePowerbankWebAPI(),
	).Methods(http.MethodPost)

	s.apiRouter.HandleFunc(
		"/station/{station_id:[0-9]+}/put-powerbank/{powerbank_id:[0-9]+}",
		s.PutPowerbankWebAPI(),
	).Methods(http.MethodPost)

	s.apiRouter.HandleFunc(
		"/station/{station_id:[0-9]+}/add-powerbank/{powerbank_id:[0-9]+}",
		s.AddPowerbankToStationWebAPI(),
	).Methods(http.MethodPost)
}

func (s *server) GetStationsWebAPI(w http.ResponseWriter, r *http.Request) {
	stations, err := s.useCase.GetStations()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetStationsWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, stations)
}

func (s *server) GetStationWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetStationWebAPI - %w", err))
		return
	}

	if err != nil {
		return
	}
	st, err := s.useCase.GetStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetStationWebAPI - %w", err))
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
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreateStationWebAPI - %w", err))
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
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("CreateStationWebAPI - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
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
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateStationWebAPI - %w", err))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("UpdateStationWebAPI - %w", err))
			return
		}
		st := entity.Station{
			AddressId:    req.AddressId,
			FreeCapacity: req.FreeCapacity,
		}

		err = s.useCase.UpdateStation(id, st)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateStationWebAPI - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) DeleteStationWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteStationWebAPI - %w", err))
		return
	}

	err = s.useCase.DeleteStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteStationWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

func (s *server) GetAllPowerbanksInStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetAllPowerbanksInStation - %w", err))
		return
	}

	powerbanks, err := s.useCase.GetAllPowerbanksInStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetAllPowerbanksInStation - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, powerbanks)
}

func (s *server) TakePowerbankWebAPI() http.HandlerFunc {
	type request struct {
		UserId      int `json:"userId"`
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := r.Context().Value(ctxKeyUser).(entity.User).ID
		powerbankId, err := strconv.Atoi(vars["powerbank_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("TakePowerbankWebAPI - %w", err))
			return
		}
		stationId, err := strconv.Atoi(vars["station_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("TakePowerbankWebAPI - %w", err))
			return
		}
		req := &request{
			UserId:      userId,
			PowerbankId: powerbankId,
			StationId:   stationId,
		}

		rez, err := s.useCase.TakePowerbank(req.UserId, req.PowerbankId, req.StationId)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("TakePowerbankWebAPI - %w", err))
			return
		}
		data := "successful"
		if !rez {
			data = "failed"
		}
		s.respond(w, r, http.StatusOK, data)
	}

}

func (s *server) PutPowerbankWebAPI() http.HandlerFunc {
	type request struct {
		UserId      int `json:"userId"`
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := r.Context().Value(ctxKeyUser).(entity.User).ID
		powerbankId, err := strconv.Atoi(vars["powerbank_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("PutPowerbankWebAPI - %w", err))
			return
		}
		stationId, err := strconv.Atoi(vars["station_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("PutPowerbankWebAPI - %w", err))
			return
		}
		req := &request{
			UserId:      userId,
			PowerbankId: powerbankId,
			StationId:   stationId,
		}

		err = s.useCase.PutPowerbank(req.UserId, req.PowerbankId, req.StationId)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("PutPowerbankWebAPI - %w", err))
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}

}

func (s *server) AddPowerbankToStationWebAPI() http.HandlerFunc {
	type request struct {
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		powerbankId, err := strconv.Atoi(vars["powerbank_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("AddPowerbankToStationWebAPI - %w", err))
			return
		}
		stationId, err := strconv.Atoi(vars["station_id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("AddPowerbankToStationWebAPI - %w", err))
			return
		}
		req := &request{
			PowerbankId: powerbankId,
			StationId:   stationId,
		}

		err = s.useCase.AddPowerbankToStation(req.PowerbankId, req.StationId)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("AddPowerbankToStationWebAPI - %w", err))
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}
