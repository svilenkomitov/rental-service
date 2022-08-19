package handler

import "github.com/svilenkomitov/rental-service/internal/rental-service/repository"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Location struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	ZIP     string  `json:"zip"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type Price struct {
	Day int `json:"day"`
}

type Rental struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	Year            int      `json:"year"`
	Length          float64  `json:"length"`
	Sleeps          int      `json:"sleeps"`
	PrimaryImageURL string   `json:"primary_image_url"`
	Price           Price    `json:"price"`
	Location        Location `json:"location"`
	User            User     `json:"user"`
}

func _mapToDomain(entity *repository.Entity) *Rental {
	return &Rental{
		Id:              entity.Id,
		Name:            entity.Name,
		Description:     entity.Description,
		Type:            entity.Type,
		Make:            entity.VehicleMake,
		Model:           entity.VehicleModel,
		Year:            entity.VehicleYear,
		Length:          entity.VehicleLength,
		Sleeps:          entity.Sleeps,
		PrimaryImageURL: entity.PrimaryImageURL,
		Price: Price{
			Day: int(entity.PricePerDay),
		},
		Location: Location{
			City:    entity.HomeCity,
			State:   entity.HomeState,
			ZIP:     entity.HomeZIP,
			Country: entity.HomeCountry,
			Lat:     entity.Lat,
			Lng:     entity.Lng,
		},
		User: User{
			Id:        entity.UserId,
			FirstName: entity.UserFirstName,
			LastName:  entity.UserLastName,
		},
	}
}
