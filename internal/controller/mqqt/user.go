package mqqt

import (
	//"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/rabbitmq/rmq_rpc/server"
)

type userRoutes struct {
	userUseCase usecase.UserAPI
}

func newUserRoutes(routes map[string]server.CallHandler, t usecase.UserAPI) {

}
