package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *server) newPowerbankRoutes(router *mux.Router, p usecase.PowerbankAPI, l logger.Interface) {
	s.router.HandleFunc("/api/powerbank/all", s.GetPowerbankWebAPI).Methods(http.MethodGet)                      // Получить список всех Powerbank
	s.router.HandleFunc("/api/powerbank/get/{id:[0-9]+}", s.GetPowerbankWebAPI).Methods(http.MethodGet)          // Получить информацию о конкретном Powerbank
	s.router.HandleFunc("/api/powerbank/create", s.CreatePowerbankWebAPI()).Methods(http.MethodPost)             // Создать новогый Powerbank
	s.router.HandleFunc("/api/powerbank/update/{id:[0-9]+}", s.UpdatePowerbankWebAPI()).Methods(http.MethodPut)  // Обновить информацию о Powerbank
	s.router.HandleFunc("/api/powerbank/delete/{id:[0-9]+}", s.DeletePowerbankWebAPI).Methods(http.MethodDelete) // Удалить Powerbank
}

func (s *server) GetPowerbanksWebAPI(w http.ResponseWriter, r *http.Request) {
	powerbanks, err := s.useCase.GetPowerbanks()
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetPowerbanksWebAPI - usecase.Powerbank.GetPowerbanks - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, powerbanks)
}

func (s *server) GetPowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetUserWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}

	p, err := s.useCase.GetPowerbank(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - GetPowerbanksWebAPI - usecase.Powerbank.GetPowerbanks - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, p)
}

func (s *server) CreatePowerbankWebAPI() http.HandlerFunc {
	type request struct {
		Used int `json:"used"` // сколько уже использована банка в часах
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreatePowerbankWebAPI - "+err.Error()))
			return
		}
		p := entity.Powerbank{
			Used: req.Used,
		}

		err := s.useCase.CreatePowerbank(p)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - CreatePowerbankWebAPI - usecase.Powerbank.CreatePowerbank - "+err.Error()))
			return
		}

		s.respond(w, r, http.StatusOK, "")
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
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - UpdatePowerbankWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("CreatePowerbankWebAPI - "+err.Error()))
			return
		}
		p := entity.Powerbank{
			Used: req.Used,
		}

		err = s.useCase.UpdatePowerbank(id, p)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - UpdatePowerbankWebAPI - usecase.Powerbank.UpdatePowerbank - "+err.Error()))
			return
		}

		s.respond(w, r, http.StatusOK, "")
	}

}

func (s *server) DeletePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - DeletePowerbankWebAPI - strconv.Atoi(vars[\"id\"]) - "+err.Error()))
		return
	}

	err = s.useCase.DeletePowerbank(id)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error - DeletePowerbankWebAPI - usecase.Powerbank.DeletePowerbank - "+err.Error()))
		return
	}

	s.respond(w, r, http.StatusOK, "")
}
