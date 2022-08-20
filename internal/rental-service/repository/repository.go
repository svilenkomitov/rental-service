package repository

import (
	"database/sql"
	"fmt"

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

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Repository
type Repository interface {
	FetchRentalById(id int) (*Entity, error)
	FetchRentals(queries map[QueryKey]interface{}) ([]*Entity, error)
}

type defaultRepository struct {
	db *storage.Database
}

func NewRepository(db *storage.Database) Repository {
	return &defaultRepository{
		db: db,
	}
}

func (r *defaultRepository) FetchRentalById(id int) (*Entity, error) {
	var entity Entity
	if err := r.db.DB.Get(&entity, fetchRentalByIdQuery, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, NewRentalNotFoundError(id)
		} else {
			return nil, err
		}
	}
	return &entity, nil
}

func (r *defaultRepository) FetchRentals(queries map[QueryKey]interface{}) ([]*Entity, error) {
	var entities []*Entity
	if err := r.db.DB.Select(&entities, buildQuery(queries)); err != nil {
		return []*Entity{}, err
	}
	return entities, nil
}
