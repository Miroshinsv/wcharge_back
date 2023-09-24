// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Miroshinsv/wcharge_back/config"
	v1 "github.com/Miroshinsv/wcharge_back/internal/controller/http/v1"

	//mqtt "github.com/Miroshinsv/wcharge_back/internal/controller/mqtt"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	repo "github.com/Miroshinsv/wcharge_back/internal/usecase/repo/postgres"

	//"github.com/Miroshinsv/wcharge_back/internal/usecase/webapi"
	"github.com/Miroshinsv/wcharge_back/pkg/httpserver"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/Miroshinsv/wcharge_back/pkg/postgres"

	//"github.com/Miroshinsv/wcharge_back/pkg/rabbitmq/rmq_rpc/server"
	"github.com/gorilla/mux"
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
	UseCase := usecase.New(
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
	handler := mux.NewRouter()
	v1.NewRouter(handler, UseCase, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
