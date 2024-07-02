package v1

import (
	"encoding/json"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (s *server) newAddressRoutes() {
	s.apiRouter.HandleFunc("/addresses", s.GetAddresses).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/addresses", s.CreateAddress).Methods(http.MethodPost)

	s.apiRouter.HandleFunc("/addresses/{id:[0-9]+}", s.GetAddress).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/addresses/{id:[0-9]+}", s.UpdateAddress).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/addresses/{id:[0-9]+}", s.DeleteAddress).Methods(http.MethodDelete)
}

// GetAddresses godoc
// @Summary 	 Get info about all addresses
// @Success      200  {array}  	entity.Address
// @Failure      500  {object}  map[string]string
// @Router       /addresses [get]
func (s *server) GetAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := s.useCase.GetAddresses()
	if err != nil {
		log.Printf("Error - GetAddresses - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, addresses)
}

// GetAddress godoc
// @Summary 	 Get info about address
// @Param        addressId   	path	int		true  	"Address Id"
// @Success      200  {object}  entity.Address
// @Failure      500  {object}  map[string]string
// @Router       /addresses/{addressId} [get]
func (s *server) GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - GetAddress - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	st, err := s.useCase.GetAddress(id)
	if err != nil {
		log.Printf("Error - GetAddress - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, st)
}

// CreateAddress godoc
// @Summary 	 Create address
// @Param        Country   		body   	string	true	"Country"
// @Param        City   		body   	string	true  	"City"
// @Param        Address   		body   	string	true  	"Address"
// @Param        Lat   			body   	float64	true  	"Latitude"
// @Param        Lng   			body   	float64	true  	"Longitude"
// @Success      200  {object}  entity.Address
// @Failure      500  {object}  map[string]string
// @Router       /addresses [post]
func (s *server) CreateAddress(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Country string  `json:"country"`
		City    string  `json:"city"`
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - CreateAddress - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	st := entity.Address{
		Country: req.Country,
		City:    req.City,
		Address: req.Address,
		Lat:     req.Lat,
		Lng:     req.Lng,
	}

	adr, err := s.useCase.CreateAddress(st)
	if err != nil {
		log.Printf("Error - CreateAddress - usecase.CreateAddress - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, adr)
}

// UpdateAddress godoc
// @Summary 	 Update address
// @Param        addressId   	path	int		true  	"Address Id"
// @Param        Country   		body   	string	true	"Country"
// @Param        City   		body   	string	true  	"City"
// @Param        Address   		body   	string	true  	"Address"
// @Param        Lat   			body   	float64	true  	"Latitude"
// @Param        Lng   			body   	float64	true  	"Longitude"
// @Success      200  {object}  entity.Address
// @Failure      500  {object}  map[string]string
// @Router       /addresses/{addressId} [put]
func (s *server) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Country string  `json:"country"`
		City    string  `json:"city"`
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - UpdateAddress - strconv.Atoi(vars[\"id\"]) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("Error - UpdateAddress - %s", err)
		s.error(w, http.StatusBadRequest, err)
		return
	}
	st := entity.Address{
		Country: req.Country,
		City:    req.City,
		Address: req.Address,
		Lat:     req.Lat,
		Lng:     req.Lng,
	}

	adr, err := s.useCase.UpdateAddress(id, st)
	if err != nil {
		log.Printf("Error - UpdateAddress - useCase.UpdateAddress - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, adr)
}

// DeleteAddress godoc
// @Summary 	 Delete address
// @Param        addressId   	path	int		true  	"Address Id"
// @Success      200  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /addresses/{addressId} [delete]
func (s *server) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error - DeleteAddress - strconv.Atoi(vars[\"id\"]) - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	err = s.useCase.DeleteAddress(id)
	if err != nil {
		log.Printf("Error - DeleteAddress - usecase.User.GetUsers - %s", err)
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, nil)
}
