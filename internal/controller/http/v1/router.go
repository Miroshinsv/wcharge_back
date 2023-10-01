// Package v1 implements routing paths. Each services in own file.
package v1

import (
	// Swagger docs.
	_ "github.com/Miroshinsv/wcharge_back/docs"
	"net/http"
)

// NewHttpRouter -.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func (s *server) NewHttpRouter() {
	s.router.HandleFunc("/", home)

	s.newUserRoutes()
	newStationRoutes(s.router, s.useCase, s.logger)
	newPowerbankRoutes(s.router, s.useCase, s.logger)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
