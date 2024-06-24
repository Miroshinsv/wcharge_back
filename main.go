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
	log.SetPrefix("BACK__")
	log.SetOutput(gelfWriter)

	//docs.SwaggerInfo.Title = "Swagger WeCharge API"
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = cfg.Swagger.URL
	//docs.SwaggerInfo.BasePath = "/api/v1"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//r := gin.New()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//err = r.Run(":8989")
	//if err != nil {
	//	log.Printf("Error - Swagger - Gin - Run: %s", err)
	//}

	app.Run(cfg)
}
