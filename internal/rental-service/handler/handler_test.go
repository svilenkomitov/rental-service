package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	r "github.com/svilenkomitov/rental-service/internal/rental-service/repository"
	"github.com/svilenkomitov/rental-service/internal/rental-service/repository/repositoryfakes"
)

const (
	FetchRentalByIdEndpoint = "/rentals/{id}"
	FetchRentalsEndpoint    = "/rentals"
)

var (
	entity = r.Entity{
		Id:              1,
		Name:            "Maupin: Vanagon Camper",
		Type:            "camper-van",
		Description:     "fermentum nullam congue arcu sollicitudin lacus suspendisse nibh semper cursus sapien quis feugiat maecenas nec turpis viverra gravida risus phasellus tortor cras gravida varius scelerisque",
		UserId:          4,
		UserFirstName:   "John",
		UserLastName:    "Smith",
		Sleeps:          4,
		PricePerDay:     15000,
		Price:           60000,
		HomeCity:        "Portland",
		HomeState:       "OR",
		HomeZIP:         "97202",
		HomeCountry:     "US",
		VehicleMake:     "Volkswagen",
		VehicleModel:    "Vanagon Camper",
		VehicleYear:     1989,
		VehicleLength:   15,
		Created:         time.Now(),
		Updated:         time.Now(),
		Lat:             45.51,
		Lng:             -122.68,
		PrimaryImageURL: "https://res.cloudinary.com/outdoorsy/image/upload/v1498568017/p/rentals/11368/images/gmtye6p2eq61v0g7f7e7.jpg",
	}

	rental = _mapToDomain(&entity)
)

func TestHandler_FetchRentalById(t *testing.T) {
	repository := repositoryfakes.FakeRepository{}
	handler := NewHandler(&repository)
	router := mux.NewRouter()

	t.Run("success", func(t *testing.T) {
		repository.FetchRentalByIdReturns(&entity, nil)
		router.HandleFunc(FetchRentalByIdEndpoint, handler.FetchRentalById)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals/1")
		expectedBody, _ := json.Marshal(rental)

		Verify(t, resp, http.StatusOK, MimeApplicationJson, string(expectedBody))
	})

	t.Run("invalid id", func(t *testing.T) {
		router.HandleFunc(FetchRentalByIdEndpoint, handler.FetchRentalById)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals/s")

		Verify(t, resp, http.StatusBadRequest, MimeApplicationJson, "\"invalid id\"")
	})

	t.Run("not found", func(t *testing.T) {
		repository.FetchRentalByIdReturns(nil, r.NewRentalNotFoundError(1))
		router.HandleFunc(FetchRentalByIdEndpoint, handler.FetchRentalById)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals/1")

		Verify(t, resp, http.StatusNotFound, MimeApplicationJson, "\"rental with id [1] not found\"")
	})
}

func TestHandler_FetchRentals(t *testing.T) {
	repository := repositoryfakes.FakeRepository{}
	handler := NewHandler(&repository)
	router := mux.NewRouter()

	t.Run("success", func(t *testing.T) {
		repository.FetchRentalsReturns([]*r.Entity{&entity}, nil)
		router.HandleFunc(FetchRentalsEndpoint, handler.FetchRentals)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals")
		expectedBody, _ := json.Marshal([]*Rental{rental})

		Verify(t, resp, http.StatusOK, MimeApplicationJson, string(expectedBody))
	})

	t.Run("unsupported query", func(t *testing.T) {
		repository.FetchRentalsReturns([]*r.Entity{}, nil)
		router.HandleFunc(FetchRentalsEndpoint, handler.FetchRentals)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals?unsuported_query=1")

		Verify(t, resp, http.StatusBadRequest, MimeApplicationJson, string("\"unsupported query param [unsuported_query]\""))
	})

	t.Run("invalid query value", func(t *testing.T) {
		repository.FetchRentalsReturns([]*r.Entity{}, nil)
		router.HandleFunc(FetchRentalsEndpoint, handler.FetchRentals)
		server := httptest.NewServer(router)
		defer server.Close()

		resp, _ := http.Get(server.URL + "/rentals?price_min=-1&price_max=a")

		Verify(t, resp, http.StatusBadRequest, MimeApplicationJson, string("\"invalid queries values: [price_min price_max]\""))
	})
}

func Verify(t *testing.T, resp *http.Response, expectedStatusCode int, expectedContentType string, expectedBody string) {
	body, err := ioutil.ReadAll(resp.Body)
	respBody := string(body)
	if nil != err {
		t.Fatalf("Failed to read the response body: %v", err)
	}

	if status := resp.StatusCode; status != expectedStatusCode {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatusCode)
	}

	if contentType := resp.Header.Get(HeaderContentType); contentType != expectedContentType {
		t.Errorf("the response contains unexpected content type: got %s want %s",
			contentType, expectedContentType)
	}

	if respBody != expectedBody {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			respBody, expectedBody)
	}
}
