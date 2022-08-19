package repository

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

const (
	fetchRentalByIdQuery = `SELECT rentals.*, users.first_name, users.last_name, 
	rentals.price_per_day * rentals.sleeps AS price 
	FROM rentals JOIN users ON rentals.user_id=users.id
	WHERE rentals.id = $1`
)

const (
	rentalNotFoundMsgFormat = "rental with id [%v] does not exist"
)

type RentalNotFoundError struct {
	msg string
}

func NewRentalNotFoundError(id int) *RentalNotFoundError {
	return &RentalNotFoundError{msg: fmt.Sprintf(rentalNotFoundMsgFormat, id)}
}

func (e *RentalNotFoundError) Error() string {
	return e.msg
}

type Repository interface {
	FetchRental(id int) (*Entity, error)
}

type defaultRepository struct {
	db *storage.Database
}

func NewRepository(db *storage.Database) Repository {
	return &defaultRepository{
		db: db,
	}
}

func (r *defaultRepository) FetchRental(id int) (*Entity, error) {
	var entity Entity
	if err := r.db.DB.Get(&entity, fetchRentalByIdQuery, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, NewRentalNotFoundError(id)
		} else {
			log.Errorf("Error occurred while fetching rental with id [%d]: %v", id, err)
			return nil, err
		}
	}
	return &entity, nil
}
