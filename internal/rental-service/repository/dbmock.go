package repository 

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func MockDataBase() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		panic("Failed to mock database: " + err.Error())
	}
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	return sqlxDb, mock
}