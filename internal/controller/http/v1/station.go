package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
)

func (s *server) newStationRoutes() {
	s.apiRouter.HandleFunc("/stations", s.GetStations).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/stations", s.CreateStation).Methods(http.MethodPost)

	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.GetStation).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.UpdateStation).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}", s.DeleteStation).Methods(http.MethodDelete)

	s.apiRouter.HandleFunc("/stations/{id:[0-9]+}/powerbanks", s.GetAllPowerbanksInStation).Methods(http.MethodGet)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks",
		s.TakePowerbank,
	).Methods(http.MethodPost)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks/{powerbank_id:[0-9]+}",
		s.PutPowerbank,
	).Methods(http.MethodPut)
	s.apiRouter.HandleFunc(
		"/stations/{station_id:[0-9]+}/powerbanks/{powerbank_id:[0-9]+}",
		s.AddPowerbankToStation,
	).Methods(http.MethodPost)
}

// GetStations godoc
// @Summary 	 Get info about all stations
// @Success      200  {array}  entity.Station
// @Failure      500  {object}  error
// @Router       /stations [get]
func (s *server) GetStations(w http.ResponseWriter, r *http.Request) {
	stations, err := s.useCase.GetStations()
	if err != nil {
		log.Printf("Error - GetStations - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, stations)
}

// GetStation godoc
// @Summary 	 Get info about station
// @Param        id   path      int  true  "Station ID"
// @Success      200  {object}  entity.Station
// @Failure      500  {object}  error
// @Router       /stations/{id} [get]
func (s *server) GetStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - GetStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	st, err := s.useCase.GetStation(id)
	if err != nil {
		log.Printf("Error - GetStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, st)
}

// CreateStation godoc
// @Summary 	 Create station
// @Param        SerialNumber   body   string	true  "Serial number of station"
// @Param        Address   		body   int		true  "Address ID"
// @Param        Capacity   	body   float64	true  "Full capacity on station"
// @Param        FreeCapacity   body   float64	true  "Free capacity on station"
// @Success      200  {object}  entity.Station
// @Failure      500  {object}  error
// @Router       /stations [post]
func (s *server) CreateStation(w http.ResponseWriter, r *http.Request) {
	type request struct {
		SerialNumber string  `json:"serial_number"`
		Address      int     `json:"address"`
		Capacity     float64 `json:"capacity"`
		FreeCapacity float64 `json:"free_capacity"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - CreateStation - %s", err)
		s.error(w, r, http.StatusBadRequest, err)
		return
	}
	st := entity.Station{
		SerialNumber: req.SerialNumber,
		Address:      req.Address,
		Capacity:     req.Capacity,
		FreeCapacity: req.FreeCapacity,
	}

	_, err := s.useCase.CreateStation(st) // TODO
	if err != nil {
		log.Printf("Error - CreateStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

// UpdateStation godoc
// @Summary 	 Update station
// @Param        id   			path      int  		true  "Station ID"
// @Param        Address   		body   	  int 		true  "Address ID"
// @Param        FreeCapacity   body      float64	true  "Free capacity on station"
// @Success      200  {object}  entity.Station
// @Failure      500  {object}  error
// @Router       /stations/{id} [put]
func (s *server) UpdateStation(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Address      int     `json:"address"`
		FreeCapacity float64 `json:"free_capacity"`
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - UpdateStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	req := &request{}
	if err = json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - UpdateStation - %s", err)
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	st := entity.Station{
		Address:      req.Address,
		FreeCapacity: req.FreeCapacity,
	}
	err = s.useCase.UpdateStation(id, st)
	if err != nil {
		log.Printf("Error - UpdateStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

// DeleteStation godoc
// @Summary 	 Delete station
// @Param        id   			path      int  		true  "Station ID"
// @Success      200  {object}  nil
// @Failure      500  {object}  error
// @Router       /stations/{id} [delete]
func (s *server) DeleteStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - DeleteStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	err = s.useCase.DeleteStation(id)
	if err != nil {
		log.Printf("Error - DeleteStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}

// GetAllPowerbanksInStation godoc
// @Summary 	 Get powerbanks in station
// @Param        id   			path      int  		true  "Station ID"
// @Success      200  {array}  entity.Powerbank
// @Failure      500  {object}  error
// @Router       /stations/{id}/powerbanks [get]
func (s *server) GetAllPowerbanksInStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - GetAllPowerbanksInStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	powerbanks, err := s.useCase.GetAllPowerbanksInStation(id)
	if err != nil {
		log.Printf("Error - GetAllPowerbanksInStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, powerbanks)
}

// TakePowerbank godoc
// @Summary 	 Take random powerbank from station
// @Param        id   			path      int  		true  "Station ID"
// @Success      200  {integer} int
// @Failure      500  {object}  error
// @Router       /stations/{id}/powerbanks [post]
func (s *server) TakePowerbank(w http.ResponseWriter, r *http.Request) {

	type request struct {
		UserId    int `json:"userId"`
		StationId int `json:"stationId"`
	}

	vars := mux.Vars(r)

	userId := r.Context().Value(ctxKeyUser).(entity.User).ID

	stationId, err := strconv.Atoi(vars["station_id"])
	if err != nil {
		log.Printf("Error - TakePowerbank - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	req := &request{
		UserId:    userId,
		StationId: stationId,
	}

	rez, err := s.useCase.TakePowerbank(req.UserId, req.StationId)
	if err != nil {
		log.Printf("Error - TakePowerbank - TakePowerbank - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	type response struct {
		PowerbankId int `json:"powerbank_id"`
	}

	data := response{rez.ID}
	//if !rez {
	//	data = "failed"
	//}
	s.respond(w, r, http.StatusOK, data)
}

// PutPowerbank godoc
// @Summary 	 Return powerbank to station
// @Param        stationId   			path      int  		true  "Station ID"
// @Param        powerbankId   			path      int  		true  "Powerbank ID"
// @Success      200  {object} 	nil
// @Failure      500  {object}  error
// @Router       /stations/{stationId}/powerbanks/{powerbankId} [put]
func (s *server) PutPowerbank(w http.ResponseWriter, r *http.Request) { // TODO ???
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
		log.Printf("Error - PutPowerbank - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	stationId, err := strconv.Atoi(vars["station_id"])
	if err != nil {
		log.Printf("Error - PutPowerbank - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
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
		log.Printf("Error - PutPowerbank - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}

// AddPowerbankToStation godoc
// @Summary 	 Add powerbank to station
// @Param        stationId   			path      int  		true  "Station ID"
// @Param        powerbankId   			path      int  		true  "Powerbank ID"
// @Param        Position   			body      int		true  "Powerbank's position in station"
// @Success      200  {object} 	nil
// @Failure      500  {object}  error
// @Router       /stations/{stationId}/powerbanks/{powerbankId} [post]
func (s *server) AddPowerbankToStation(w http.ResponseWriter, r *http.Request) { // TODO ???

	type request struct {
		PowerbankId int `json:"powerbankId"`
		StationId   int `json:"stationId"`
		Position    int `json:"position"`
	}

	vars := mux.Vars(r)
	powerbankId, err := strconv.Atoi(vars["powerbank_id"])
	if err != nil {
		log.Printf("Error - AddPowerbankToStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	stationId, err := strconv.Atoi(vars["station_id"])
	if err != nil {
		log.Printf("Error - AddPowerbankToStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	req := &request{
		PowerbankId: powerbankId,
		StationId:   stationId,
	}

	err = s.useCase.AddPowerbankToStation(req.PowerbankId, req.StationId, req.Position)
	if err != nil {
		log.Printf("Error - AddPowerbankToStation - %s", err)
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}
