package storage

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	*sqlx.DB
}

func Connect(config *Config) (*Database, error) {

	log.Infof("Connecting to database [Host: %s], [Port: %d], [Name: %s]",
		config.Host, config.Port, config.DbName)

	db, err := sqlx.Open(config.DbDriver, config.GetDataSourceName())

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Infof("Connected to database [Host: %s], [Port: %d], [Name: %s]",
		config.Host, config.Port, config.DbName)

	return &Database{db}, err
}