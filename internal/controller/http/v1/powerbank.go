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

type powerbankRoutes struct {
	p usecase.PowerbankAPI
	l logger.Interface
}

func newPowerbankRoutes(router *mux.Router, p usecase.PowerbankAPI, l logger.Interface) {

	sr := &powerbankRoutes{p, l}
	router.HandleFunc("/api/powerbank/all", sr.GetPowerbankWebAPI).Methods("GET")               // Получить список всех Powerbank
	router.HandleFunc("/api/powerbank/get/{id}", sr.GetPowerbankWebAPI).Methods("GET")          // Получить информацию о конкретном Powerbank
	router.HandleFunc("/api/powerbank/create", sr.CreatePowerbankWebAPI).Methods("POST")        // Создать новогый Powerbank
	router.HandleFunc("/api/powerbank/update/{id}", sr.UpdatePowerbankWebAPI).Methods("PUT")    // Обновить информацию о Powerbank
	router.HandleFunc("/api/powerbank/delete/{id}", sr.DeletePowerbankWebAPI).Methods("DELETE") // Удалить Powerbank
}

func RequestToJSONPowerbank(w http.ResponseWriter, r *http.Request) (entity.Powerbank, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return entity.Powerbank{}, errors.New("Content Type is not application/json")
	}
	var p entity.Powerbank
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return entity.Powerbank{}, errors.New("Bad Request - entity is not Powerbank")
	}

	return p, nil
}

func (ur *powerbankRoutes) GetPowerbanksWebAPI(w http.ResponseWriter, r *http.Request) {
	powerbanks, err := ur.p.GetPowerbanks()
	if err != nil {
		errorResponse(w, "error - GetPowerbanksWebAPI - usecase.Powerbank.GetPowerbanks - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, s := range powerbanks {
		json.NewEncoder(w).Encode(s)
	}
}

func (ur *powerbankRoutes) GetPowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONPowerbank(w, r)
	/*id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}*/
	if err != nil {
		return
	}
	user, err := ur.p.GetPowerbank(s)
	if err != nil {
		errorResponse(w, "error - GetPowerbanksWebAPI - usecase.Powerbank.GetPowerbanks - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (ur *powerbankRoutes) CreatePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONPowerbank(w, r)
	if err != nil {
		return
	}

	err = ur.p.CreatePowerbank(s)
	if err != nil {
		errorResponse(w, "error - CreatePowerbankWebAPI - usecase.Powerbank.CreatePowerbank - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *powerbankRoutes) UpdatePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONPowerbank(w, r)
	if err != nil {
		return
	}

	err = ur.p.UpdatePowerbank(s)
	if err != nil {
		errorResponse(w, "error - UpdatePowerbankWebAPI - usecase.Powerbank.UpdatePowerbank - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ur *powerbankRoutes) DeletePowerbankWebAPI(w http.ResponseWriter, r *http.Request) {
	s, err := RequestToJSONPowerbank(w, r)
	if err != nil {
		return
	}

	user := ur.p.DeletePowerbank(s)
	if err != nil {
		errorResponse(w, "error - GetPowerbanksWebAPI - usecase.Powerbank.GetPowerbanks - "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}
