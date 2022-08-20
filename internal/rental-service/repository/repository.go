package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/svilenkomitov/rental-service/internal/storage"
)

const (
	METER_TO_MILES = 0.000621371192
	MAX_MILES      = 100

	fetchRentalByIdQuery = `SELECT rentals.*, users.first_name, users.last_name, 
	rentals.price_per_day * rentals.sleeps AS price 
	FROM rentals JOIN users ON rentals.user_id=users.id
	WHERE rentals.id = $1`

	fetchRentalsSelectQuery = `SELECT rentals.*, users.first_name, users.last_name, 
	rentals.price_per_day * rentals.sleeps as price FROM rentals JOIN users ON rentals.user_id=users.id`
	priceMinConditionQuery = `rentals.price_per_day * rentals.sleeps >= %v`
	priceMaxConditionQuery = `rentals.price_per_day * rentals.sleeps <= %v`
	idsInConditionQuery    = "rentals.id IN (%v)"
	nearConditionQuery     = "st_distance(geography(st_makepoint(rentals.lat,rentals.lng)), geography(st_makepoint(%v))) * %f < %d"
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
	if err := r.db.DB.Select(&entities, buildFetchRentalsQuery(queries)); err != nil {
		return []*Entity{}, err
	}
	return entities, nil
}

func buildFetchRentalsQuery(queries map[QueryKey]interface{}) string {
	query := NewQueryBuilder().Select(fetchRentalsSelectQuery)

	for key, value := range queries {
		switch key {
		case PRICE_MIN_KEY:
			query.Where(fmt.Sprintf(priceMinConditionQuery, value))
		case PRICE_MAX_KEY:
			query.Where(fmt.Sprintf(priceMaxConditionQuery, value))
		case IDS_KEY:
			query.Where(fmt.Sprintf(idsInConditionQuery, joinIntArr(value.([]int))))
		case NEAR_KEY:
			query.Where(fmt.Sprintf(nearConditionQuery, joinFloatArr(value.([]float64)), METER_TO_MILES, MAX_MILES))
		case SORT_KEY:
			query.OrderBy(value.(string))
		case LIMIT_KEY:
			query.Limit(value.(int))
		case OFFSET_KEY:
			query.Offset(value.(int))
		}
	}
	return query.Build()
}

func joinIntArr(arr []int) string {
	var str []string
	for _, i := range arr {
		str = append(str, fmt.Sprintf("%v", i))
	}
	return strings.Join(str, ", ")
}

func joinFloatArr(arr []float64) string {
	var str []string
	for _, i := range arr {
		str = append(str, fmt.Sprintf("%v", i))
	}
	return strings.Join(str, ", ")
}
