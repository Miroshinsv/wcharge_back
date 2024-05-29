// Package app configures and runs application.
package app

import (
	"encoding/json"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/streadway/amqp"

	"github.com/Miroshinsv/wcharge_back/config"
	v1 "github.com/Miroshinsv/wcharge_back/internal/controller/http/v1"
	grpcclient "github.com/Miroshinsv/wcharge_back/internal/usecase/repo/grpc"

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

	m, err := grpcclient.NewMqttV1Client(cfg, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// Use case
	useCase := usecase.New(
		repo.New(pg),
		m,
	)

	test(useCase)

	// HTTP Server
	v1.Start(cfg, useCase, l)
}

func test(u *usecase.UseCase) {
	cfg, _ := config.NewConfig()
	conn, _ := amqp.Dial(cfg.Rabbit.URL)
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"mqtt_test", // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {

	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	err = ch.QueueBind(
		q.Name,      // queue name
		"",          // routing key
		"mqtt_test", // exchange
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	type Data struct {
		T int
		D string
	}

	var forever chan struct{}

	go func() {
		type Powerbank struct {
			Position     int
			SerialNumber string
			Capacity     int
			Used         int
		}

		type FullStation struct {
			SerialNumber string
			Capacity     int
			Powerbanks   []Powerbank
		}

		var msg FullStation

		for d := range msgs {
			err := json.Unmarshal(d.Body, &msg)

			if err != nil {
				continue
			}

			if msg.SerialNumber != "" {
				station, err := u.CreateStation(
					entity.Station{
						SerialNumber: msg.SerialNumber,
						Capacity:     msg.Capacity,
						AddressId:    1,
					},
				)

				if err != nil {
					continue
				}

				for _, p := range msg.Powerbanks {
					powerbank, err := u.CreatePowerbank(
						entity.Powerbank{
							Position:     p.Position,
							SerialNumber: p.SerialNumber,
						},
					)

					if err != nil {
						break
					}

					err = u.AddPowerbankToStation(powerbank.ID, station.ID, p.Position)

					if err != nil {
						break
					}

				}
			}
		}

		//if msg.SerialNumber != "" {
		//
		//}
	}()

	<-forever
}
