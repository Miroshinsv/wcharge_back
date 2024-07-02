package v1

import (
	"encoding/json"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (s *server) newPowerbankRoutes() {
	s.apiRouter.HandleFunc("/powerbanks", s.GetPowerbanks).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/powerbanks", s.CreatePowerbank).Methods(http.MethodPost)

	s.apiRouter.HandleFunc("/powerbanks/{id:[0-9]+}", s.GetPowerbank).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/powerbanks/{id:[0-9]+}", s.UpdatePowerbank).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/powerbanks/{id:[0-9]+}", s.DeletePowerbank).Methods(http.MethodDelete)
}

// GetPowerbanks godoc
// @Summary 	 Get info about all powerbanks
// @Success      200  {array}  	entity.Powerbank
// @Failure      500  {object}  map[string]string
// @Router       /powerbanks [get]
func (s *server) GetPowerbanks(w http.ResponseWriter, r *http.Request) {
	powerbanks, err := s.useCase.GetPowerbanks()
	if err != nil {
		log.Printf("Error - GetPowerbanks - useCase.GetPowerbanks() - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, powerbanks)
}

// GetPowerbank godoc
// @Summary 	 Get info about powerbank
// @Param        powerbankId   	path	int		true  	"Powerbank Id"
// @Success      200  {object}  entity.Powerbank
// @Failure      500  {object}  map[string]string
// @Router       /powerbanks/{powerbankId} [get]
func (s *server) GetPowerbank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - GetPowerbank - strconv.Atoi(vars[\"id\"]) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	p, err := s.useCase.GetPowerbank(id)
	if err != nil {
		log.Printf("Error - GetPowerbank - s.useCase.GetPowerbank(id) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, p)
}

// CreatePowerbank godoc
// @Summary 	 Create powerbank
// @Param        SerialNumber   body   	string	true	"Serial number of powerbank"
// @Param        Capacity   	body   	float64	true  	"Full capacity on powerbank"
// @Param        Used   		body   	bool	true  	"Status on use's powerbanks"
// @Success      200  {object}  entity.Powerbank
// @Failure      500  {object}  map[string]string
// @Router       /powerbanks [post]
func (s *server) CreatePowerbank(w http.ResponseWriter, r *http.Request) {
	type request struct {
		SerialNumber string  `json:"serial_number"`
		Capacity     float64 `json:"capacity"` // объем заряда
		Used         bool    `json:"used"`     // сколько уже использована банка в часах
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - CreatePowerbank - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	p := entity.Powerbank{
		SerialNumber: req.SerialNumber,
		Capacity:     req.Capacity,
		Used:         req.Used,
	}

	pb, err := s.useCase.CreatePowerbank(p)
	if err != nil {
		log.Printf("Error - CreatePowerbank - useCase.CreatePowerbank - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, pb)
}

// UpdatePowerbank godoc
// @Summary 	 Create powerbank
// @Param        powerbankId   	path	int		true  	"Powerbank Id"
// @Param        Used   		body   	bool	true  	"Status on use's powerbanks"
// @Success      200  {object}  entity.Powerbank
// @Failure      500  {object}  map[string]string
// @Router       /powerbanks/{powerbankId} [put]
func (s *server) UpdatePowerbank(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Used bool `json:"used"` // сколько уже использована банка в часах
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - UpdatePowerbank - strconv.Atoi(vars[\"id\"]) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - CreatePowerbank - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	p := entity.Powerbank{
		Used: req.Used,
	}

	pb, err := s.useCase.UpdatePowerbank(id, p)
	if err != nil {
		log.Printf("Error - UpdatePowerbank - usecase.UpdatePowerbank - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, pb)
}

// DeletePowerbank godoc
// @Summary 	 Create powerbank
// @Param        powerbankId   	path	int		true  	"Powerbank Id"
// @Success      200  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /powerbanks/{powerbankId} [delete]
func (s *server) DeletePowerbank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - DeletePowerbank - strconv.Atoi(vars[\"id\"]) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	err = s.useCase.DeletePowerbank(id)
	if err != nil {
		log.Printf("Error - DeletePowerbank - usecase.Powerbank.DeletePowerbank - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, nil)
}
