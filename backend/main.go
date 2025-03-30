package main

import (
	"log"
	"socialNetwork/internal/app"
	"socialNetwork/pkg/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Init(cfg)
}
