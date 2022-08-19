package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/svilenkomitov/rental-service/internal/rental-service/repository"
)

const (
	HeaderContentType   = "Content-Type"
	HeaderLocation      = "Location"
	MimeApplicationJson = "application/json"
)

type Handler struct {
	repository repository.Repository
}

func NewHandler(repository repository.Repository) Handler {
	return Handler{
		repository: repository,
	}
}

func (handler *Handler) FetchRentalById(resp http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if nil != err {
		_writeJsonResponse(resp, http.StatusBadRequest, "invalid id")
		return
	}

	rental, err := handler.repository.FetchRentalById(id)

	if err != nil {
		log.Errorf("Error occurred while fetching rental with id [%d]: %v", id, err)
		if _, ok := err.(*repository.RentalNotFoundError); ok {
			_writeJsonResponse(resp, http.StatusNotFound, "rental not found")
			return
		}
		_writeJsonResponse(resp, http.StatusInternalServerError, "internal server error")
		return
	}

	_writeJsonResponse(resp, http.StatusOK, _mapToDomain(rental))
}

func _writeJsonResponse(w http.ResponseWriter, code int, resp interface{}) {
	w.Header().Set(HeaderContentType, MimeApplicationJson)
	w.WriteHeader(code)

	data, err := json.Marshal(resp)
	if nil != err {
		panic("Failed to serialize response payload to json: " + err.Error())
	}

	_, _ = w.Write(data)
}
