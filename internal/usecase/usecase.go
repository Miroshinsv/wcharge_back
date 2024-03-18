package usecase

import grpcclient "github.com/Miroshinsv/wcharge_back/internal/usecase/repo/grpc"

// UseCase controller -> UseCase -> repo -> entity
type UseCase struct {
	postgres PostgresRepo
	//webapi  WebAPIRepo
	mqtt *grpcclient.MqttV1Client
}

func New(r PostgresRepo, m *grpcclient.MqttV1Client) *UseCase {
	return &UseCase{
		postgres: r,
		mqtt:     m,
	}
}
