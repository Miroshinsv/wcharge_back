package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) newPowerbankRoutes() {
	s.apiRouter.HandleFunc("/powerbank/all", s.GetPowerbanksWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/powerbank/get/{id:[0-9]+}", s.GetPowerbankWebAPI).Methods(http.MethodGet)
	s.apiRouter.HandleFunc("/powerbank/create", s.CreatePowerbankWebAPI()).Methods(http.MethodPost)
	s.apiRouter.HandleFunc("/powerbank/update/{id:[0-9]+}", s.UpdatePowerbankWebAPI()).Methods(http.MethodPut)
	s.apiRouter.HandleFunc("/powerbank/delete/{id:[0-9]+}", s.DeletePowerbankWebAPI).Methods(http.MethodDelete)
}

func (s *server) GetPowerbanksWebAPI(w http.ResponseWriter, r *http.Request) {
	powerbanks, err := s.useCase.GetPowerbanks()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetPowerbanksWebAPI - useCase.GetPowerbanks() - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, powerbanks)
}

func (s *server) GetPowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetPowerbankWebAPI - strconv.Atoi(vars[\"id\"]) - %w", err))
		return
	}

	p, err := s.useCase.GetPowerbank(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("GetPowerbankWebAPI - s.useCase.GetPowerbank(id) - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, p)
}

func (s *server) CreatePowerbankWebAPI() http.HandlerFunc {
	type request struct {
		SerialNumber string `json:"serial_number"`
		Capacity     int    `json:"capacity"` // объем заряда
		Used         int    `json:"used"`     // сколько уже использована банка в часах
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreatePowerbankWebAPI - %w", err))
			return
		}
		p := entity.Powerbank{
			SerialNumber: req.SerialNumber,
			Capacity:     req.Capacity,
			Used:         req.Used,
		}

		err := s.useCase.CreatePowerbank(p)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("CreatePowerbankWebAPI - useCase.CreatePowerbank - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) UpdatePowerbankWebAPI() http.HandlerFunc {
	type request struct {
		Used int `json:"used"` // сколько уже использована банка в часах
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdatePowerbankWebAPI - strconv.Atoi(vars[\"id\"]) - %w", err))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreatePowerbankWebAPI - %w", err))
			return
		}
		p := entity.Powerbank{
			Used: req.Used,
		}

		err = s.useCase.UpdatePowerbank(id, p)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("UpdatePowerbankWebAPI - usecase.UpdatePowerbank - %w", err))
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}

}

func (s *server) DeletePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeletePowerbankWebAPI - strconv.Atoi(vars[\"id\"]) - %w", err))
		return
	}

	err = s.useCase.DeletePowerbank(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("DeletePowerbankWebAPI - usecase.Powerbank.DeletePowerbank - %w", err))
		return
	}

	s.respond(w, r, http.StatusOK, nil)
}
