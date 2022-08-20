package storage

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	POSTGIS_EXTENSION      string = "postgis"
	validateExtensionQuery string = "SELECT pg_extension.extname FROM pg_extension where pg_extension.extname='%s'"
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

	if err := installExtension(db, POSTGIS_EXTENSION); err != nil {
		return nil, err
	}

	if err := _validateExtensionInstalled(db, POSTGIS_EXTENSION); err != nil {
		return nil, err
	}

	log.Infof("Connected to database [Host: %s], [Port: %d], [Name: %s]",
		config.Host, config.Port, config.DbName)

	return &Database{db}, err
}

func installExtension(db *sqlx.DB, ext string) error {
	if _, err := db.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %s;", ext)); err != nil {
		return err
	}
	log.Infof("Extension [%s] installed.", ext)
	return nil
}

func _validateExtensionInstalled(db *sqlx.DB, ext string) error {
	var e interface{}
	if err := db.Get(&e, fmt.Sprintf(validateExtensionQuery, ext)); err != nil {
		log.Errorf("missing required extension: [%v]", ext)
		return err
	}
	return nil
}
