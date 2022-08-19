package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/svilenkomitov/rental-service/internal/rental-service/repository"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

func main() {
	dbConfig := storage.LoadConfig()
	db, err := storage.Connect(dbConfig)

	if err != nil {
		log.Fatalf("Connecting to database failed: %v", err)
	}

	rental, err := repository.NewRepository(db).FetchRental(2)
	log.Info(rental)
}