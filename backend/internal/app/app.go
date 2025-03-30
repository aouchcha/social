package app

import (
	"io"
	"log"
	"os"

	"socialNetwork/internal/controller"
	"socialNetwork/internal/repository"
	"socialNetwork/internal/server"
	"socialNetwork/internal/service"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/database/sqlite"
)

func Init(config *config.Conf) {
	// Prepare logger
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Printf("cannot create log file: %v", err)
	}
	defer file.Close()
	logWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(logWriter)

	// Prepare database
	db, err := sqlite.InitDb(&config.Database)
	if err != nil {
		log.Fatalf("error occured while connecting database: %s", err.Error())
		return
	}
	// Close connection database
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal("can't close connection db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()

	// Prepare router <- -> service  <- -> repository
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := controller.NewHandler(service)
	server := new(server.Server)
	// Start listening server
	log.Fatalf("error occured while listening server: %s", server.Run(&config.API, handler.InitRoutes(config)))
}
