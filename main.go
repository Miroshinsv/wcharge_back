package main

import (
	"github.com/Miroshinsv/wcharge_back/config"
	"github.com/Miroshinsv/wcharge_back/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
