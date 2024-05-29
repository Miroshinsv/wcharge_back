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
	s.apiRouter.HandleFunc("/stations", s.GetStationsWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/stations", s.CreateStationWebAPI).Methods(http.MethodPost)

	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.GetStationWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.UpdateStationWebAPI).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.DeleteStationWebAPI).Methods(http.MethodDelete)

	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}/powerbanks", s.GetAllPowerbanksInStation).Methods(http.MethodGet)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks",
		s.TakePowerbankWebAPI,
	).Methods(http.MethodPost)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks/{powerbank_id:[0-9]+}",
		s.PutPowerbankWebAPI,
	).Methods(http.MethodPut)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks/{powerbank_id:[0-9]+}",
		s.AddPowerbankToStationWebAPI,
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

	st, err := s.useCase.GetStation(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetStationWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, st)
}

func (s *server) CreateStationWebAPI(w http.ResponseWriter, r *http.Request) {
	type request struct {
		SerialNumber string `json:"serial_number"`
		AddressId    int    `json:"address"`
		Capacity     int    `json:"capacity"`
		FreeCapacity int    `json:"free_capacity"`
	}

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

	_, err := s.useCase.CreateStation(st) // TODO
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("CreateStationWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

func (s *server) UpdateStationWebAPI(w http.ResponseWriter, r *http.Request) {
	type request struct {
		AddressId    int `json:"address"`
		FreeCapacity int `json:"free_capacity"`
	}

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

func (s *server) TakePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {

	type request struct {
		UserId    int `json:"userId"`
		StationId int `json:"stationId"`
	}

	vars := mux.Vars(r)

	userId := r.Context().Value(ctxKeyUser).(entity.User).ID

	stationId, err := strconv.Atoi(vars["station_id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("TakePowerbankWebAPI - %w", err))
		return
	}
	req := &request{
		UserId:    userId,
		StationId: stationId,
	}

	rez, err := s.useCase.TakePowerbank(req.UserId, req.StationId)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		powerbankId int `json:"powerbank_id"`
	}

	data := response{rez.ID}
	//if !rez {
	//	data = "failed"
	//}
	s.respond(w, r, http.StatusOK, data)
}

func (s *server) PutPowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	type request struct {
		UserId      int `json:"userId"`
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
		Position    int `json:"position"`
	}

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
	position, err := strconv.Atoi(vars["position"])
	req := &request{
		UserId:      userId,
		PowerbankId: powerbankId,
		StationId:   stationId,
		Position:    position,
	}

	err = s.useCase.PutPowerbank(req.UserId, req.PowerbankId, req.StationId, req.Position)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("PutPowerbankWebAPI - %w", err))
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}

func (s *server) AddPowerbankToStationWebAPI(w http.ResponseWriter, r *http.Request) {

	type request struct {
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
		Position    int `json:"position"`
	}

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

	err = s.useCase.AddPowerbankToStation(req.PowerbankId, req.StationId, req.Position)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("AddPowerbankToStationWebAPI - %w", err))
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}
