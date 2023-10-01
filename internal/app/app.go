// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/Miroshinsv/wcharge_back/config"
	v1 "github.com/Miroshinsv/wcharge_back/internal/controller/http/v1"

	//mqtt "github.com/Miroshinsv/wcharge_back/internal/controller/mqtt"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	repo "github.com/Miroshinsv/wcharge_back/internal/usecase/repo/postgres"

	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	fmt.Println(cfg)
	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	useCase := usecase.New(
		repo.New(pg),
	)

	// MQTT Server
	/*
		mqttRouter := mqtt.NewRouter(UserUseCase)

		rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
		}
	*/

	// HTTP Server
	v1.Start(cfg, useCase, l)
}
