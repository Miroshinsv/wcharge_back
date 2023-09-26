// Package v1 implements routing paths. Each services in own file.
package v1

import (
	// Swagger docs.
	_ "github.com/Miroshinsv/wcharge_back/docs"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter -.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(router *mux.Router, useCase *usecase.UseCase, l logger.Interface) {
	router.HandleFunc("/", home).Methods("GET")
	newUserRoutes(router, useCase, l)
	newStationRoutes(router, useCase, l)
	newPowerbankRoutes(router, useCase, l)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
