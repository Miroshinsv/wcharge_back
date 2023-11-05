package v1

import (
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"testing"
)

type testServer struct {
	router           *mux.Router
	afterLoginRouter *mux.Router
	useCase          *usecase.UseCase
	sessionStore     sessions.Store
	logger           logger.Interface
}

type testUseCase interface {
}

func newTestServer() *server {

	s := &server{
		router:       mux.NewRouter(),
		useCase:      nil,
		sessionStore: nil,
		logger:       nil,
	}
	return s
}

func TestAPIServer_HandleHello(t *testing.T) {

}
