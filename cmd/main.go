package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

func main() {
	dbConfig := storage.LoadConfig()
	_, err := storage.Connect(dbConfig)

	if err != nil {
		log.Fatalf("Connecting to database failed: %v", err)
	}
}