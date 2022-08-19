package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/svilenkomitov/rental-service/internal/server"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

func main() {
	dbConfig := storage.LoadConfig()
	db, err := storage.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Connecting to database failed: %v", err)
	}

	serverConfig := server.LoadConfig()
	server := server.NewServer(serverConfig, db)
	if err := server.Start(); err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}
