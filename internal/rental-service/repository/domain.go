package repository

import (
	"time"
)

type Entity struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	Description string `db:"description"`

	UserId        int    `db:"user_id"`
	UserFirstName string `db:"first_name"`
	UserLastName  string `db:"last_name"`

	Sleeps      int     `db:"sleeps"`
	PricePerDay int64   `db:"price_per_day"`
	Price       float64 `db:"price"`

	HomeCity    string `db:"home_city"`
	HomeState   string `db:"home_state"`
	HomeZIP     string `db:"home_zip"`
	HomeCountry string `db:"home_country"`

	VehicleMake   string  `db:"vehicle_make"`
	VehicleModel  string  `db:"vehicle_model"`
	VehicleYear   int     `db:"vehicle_year"`
	VehicleLength float64 `db:"vehicle_length"`

	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`

	Lat float64 `db:"lat"`
	Lng float64 `db:"lng"`

	PrimaryImageURL string `db:"primary_image_url"`
}
