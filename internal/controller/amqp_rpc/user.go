package amqprpc

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/rabbitmq/rmq_rpc/server"
)

type userRoutes struct {
	translationUseCase usecase.User
}

func newUserRoutes(routes map[string]server.CallHandler, t usecase.User) {
	r := &userRoutes{t}
	{
		routes["getHistory"] = r.getHistory()
	}
}

type historyResponse struct {
	History []entity.User `json:"history"`
}

func (r *userRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		translations, err := r.UserUseCase.History(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - translationRoutes - getHistory - r.translationUseCase.History: %w", err)
		}

		response := historyResponse{translations}

		return response, nil
	}
}
