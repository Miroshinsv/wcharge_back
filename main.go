package main

import (
	"github.com/Miroshinsv/wcharge_back/config"
	"github.com/Miroshinsv/wcharge_back/internal/app"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	graylogAddr := cfg.Graylog.URL // Укажите адрес вашего сервера Graylog
	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewUDPWriter: %s", err)
	}

	log.SetOutput(gelfWriter)
	//log.Println("Это информационное сообщение")
	//log.Printf("Это сообщение с форматированием: %s", "пример")
	//err = someFunction()
	//if err != nil {
	//	log.Printf("Ошибка: %s", err)
	//}

	//// Создание нового логгера
	//logger := log.New(gelfWriter, "", 0)
	//// Пример отправки сообщения об ошибке
	//err = someFunction()
	//if err != nil {
	//	logger.Printf("short_message: %s, full_message: %s, level: %d", "Ошибка произошла", err.Error(), 3)
	//}

	/**
	graylogAddr := "localhost:12201" // Укажите адрес вашего сервера Graylog
	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewUDPWriter: %s", err)
	}
	logger := logrus.New()
	logger.SetOutput(gelfWriter)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"type": "example",
	}).Info("Это информационное сообщение")

	logger.WithFields(logrus.Fields{
		"type": "example",
	}).Error("Это сообщение об ошибке")

	err = someFunction()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"type": "error",
		}).Errorf("Ошибка: %s", err)
	}

	*/

	// Run
	app.Run(cfg)
}

//func someFunction() error {
//	// Ваша функция, которая может вернуть ошибку
//	return errors.New("1234")
//}
