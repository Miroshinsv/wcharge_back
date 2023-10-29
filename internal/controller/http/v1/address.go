package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) newAddressRoutes() {
	s.apiRouter.HandleFunc("/address/all", s.GetAddressesWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/address/get/{id:[0-9]+}", s.GetAddressWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/address/create", s.CreateAddressWebAPI()).Methods(http.MethodPost)
	s.apiRouter.HandleFunc("/address/update/{id:[0-9]+}", s.UpdateAddressWebAPI()).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/address/delete/{id:[0-9]+}", s.DeleteAddressWebAPI).Methods(http.MethodDelete)
}

func (s *server) GetAddressesWebAPI(w http.ResponseWriter, r *http.Request) {
	addresses, err := s.useCase.GetAddresses()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetAddressesWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, addresses)
}

func (s *server) GetAddressWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetAddressWebAPI - %w", err))
		return
	}

	if err != nil {
		return
	}
	st, err := s.useCase.GetAddress(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetAddressWebAPI - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, st)
}

func (s *server) CreateAddressWebAPI() http.HandlerFunc {
	type request struct {
		Country string  `json:"country"`
		City    string  `json:"city"`
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreateAddressWebAPI - %w", err))
			return
		}
		st := entity.Address{
			Country: req.Country,
			City:    req.City,
			Address: req.Address,
			Lat:     req.Lat,
			Lng:     req.Lng,
		}

		err := s.useCase.CreateAddress(st)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("CreateAddressWebAPI - usecase.CreateAddress - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) UpdateAddressWebAPI() http.HandlerFunc {
	type request struct {
		Country string  `json:"country"`
		City    string  `json:"city"`
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateAddressWebAPI - strconv.Atoi(vars[\"id\"]) - %w", err))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("UpdateAddressWebAPI - %w", err))
			return
		}
		st := entity.Address{
			Country: req.Country,
			City:    req.City,
			Address: req.Address,
			Lat:     req.Lat,
			Lng:     req.Lng,
		}

		err = s.useCase.UpdateAddress(id, st)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdateAddressWebAPI - useCase.UpdateAddress - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) DeleteAddressWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteAddressWebAPI - strconv.Atoi(vars[\"id\"]) - %w", err))
		return
	}

	err = s.useCase.DeleteAddress(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeleteAddressWebAPI - usecase.User.GetUsers - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}
