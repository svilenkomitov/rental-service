package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

func TestRepository_FetchRentalById(t *testing.T) {

	sqlxDb, mock := MockDataBase()
	repository := NewRepository(&storage.Database{DB: sqlxDb})

	entityRows := sqlmock.NewRows([]string{"user_id", "name", "type",
		"description", "sleeps", "price_per_day", "home_city", "home_state", "home_zip", "home_country", "vehicle_make",
		"vehicle_model", "vehicle_year", "vehicle_length", "created", "updated", "lat", "lng", "primary_image_url",
		"price", "first_name", "last_name",
	})

	fetchRentalByIdQuery := regexp.QuoteMeta(`SELECT rentals.*, users.first_name, users.last_name, rentals.price_per_day * rentals.sleeps AS price FROM rentals JOIN users ON rentals.user_id=users.id WHERE rentals.id = $1`)

	t.Run("success", func(t *testing.T) {
		id := 1
		mock.ExpectQuery(fetchRentalByIdQuery).
			WithArgs(id).
			WillReturnRows(entityRows.
				AddRow(id, "Maupin: Vanagon Camper", "camper-van", "fermentum nullam congue arcu sollicitudin lacus suspendisse nibh semper cursus sapien quis feugiat maecenas nec turpis viverra gravida risus phasellus tortor cras gravida varius scelerisque",
					4, 15000, "Portland", "OR", "97202", "US", "Volkswagen", "Vanagon Camper", 1989, 15, time.Now(), time.Now(), 45.51, -122.68, "https://res.cloudinary.com/outdoorsy/image/upload/v1498568017/p/rentals/11368/images/gmtye6p2eq61v0g7f7e7.jpg", 234, "John", "Smith"),
			)
		if _, err := repository.FetchRentalById(id); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("record not found", func(t *testing.T) {
		id := 1
		expectedError := NewRentalNotFoundError(id)
		mock.ExpectQuery(fetchRentalByIdQuery).
			WithArgs(id).
			WillReturnRows(entityRows)

		_, err := repository.FetchRentalById(id)
		if err == nil {
			t.Errorf("expected error \"%v\"; received: \"%v\"", expectedError, err)
		}
		if _, ok := err.(*RentalNotFoundError); !ok {
			t.Errorf("expected error \"%v\"; received: \"%v\"", expectedError, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
