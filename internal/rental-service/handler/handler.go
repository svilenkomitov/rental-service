package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
	id, isValid := _toInt(mux.Vars(req)["id"])
	if !isValid {
		_writeJsonResponse(resp, http.StatusBadRequest, "invalid id")
		return
	}

	rental, err := handler.repository.FetchRentalById(id.(int))

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

func (handler *Handler) FetchRentals(resp http.ResponseWriter, req *http.Request) {

	// TODO: validate supported queries

	validQueries := make(map[repository.QueryKey]interface{})
	invalidQueries := make([]repository.QueryKey, 0)

	_validate(validQueries, &invalidQueries, req, repository.PRICE_MIN_KEY, _toInt)
	_validate(validQueries, &invalidQueries, req, repository.PRICE_MAX_KEY, _toInt)
	_validate(validQueries, &invalidQueries, req, repository.LIMIT_KEY, _toInt)
	_validate(validQueries, &invalidQueries, req, repository.OFFSET_KEY, _toInt)
	_validate(validQueries, &invalidQueries, req, repository.IDS_KEY, _toIntArray)
	_validate(validQueries, &invalidQueries, req, repository.NEAR_KEY, _toFloatArray)
	_validate(validQueries, &invalidQueries, req, repository.SORT_KEY, _toString) // TODO: add ASC DSC

	if len(invalidQueries) > 0 {
		log.Errorf("queries validation failed. invalid queries: %v", invalidQueries)
		_writeJsonResponse(resp, http.StatusBadRequest, fmt.Sprintf("invalid queries: %v", invalidQueries))
		return
	}

	entities, err := handler.repository.FetchRentals(validQueries)

	if err != nil {
		log.Errorf("Error occurred while fetching rentals with queries [%v]: %v", validQueries, err)
		_writeJsonResponse(resp, http.StatusInternalServerError, "internal server error")
		return
	}

	rentals := make([]*Rental, 0)
	for _, entity := range entities {
		rentals = append(rentals, _mapToDomain(entity))
	}

	_writeJsonResponse(resp, http.StatusOK, rentals)
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

type toFunc func(string) (interface{}, bool)

func _validate(validQieries map[repository.QueryKey]interface{},
	invalidQueries *[]repository.QueryKey,
	req *http.Request,
	key repository.QueryKey,
	to toFunc) {

	queryStr := _getQuery(string(key), req)
	if queryStr != "" {
		if value, isValid := to(queryStr); isValid {
			validQieries[key] = value
		} else {
			*invalidQueries = append(*invalidQueries, key)
		}
	}
}

func _toIntArray(str string) (interface{}, bool) {
	strArr := _parseArr(str)
	intArr := make([]int, 0, len(strArr))
	for _, str := range strArr {
		i, isValid := _toInt(str)
		if !isValid {
			log.Debugf("Error occurred while parsing toIntArray [%v].", strArr)
			return nil, false
		}
		intArr = append(intArr, i.(int))
	}
	return intArr, true
}

func _toFloatArray(str string) (interface{}, bool) {
	strArr := _parseArr(str)
	floatArr := make([]float64, 0, len(strArr))
	for _, str := range strArr {
		i, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Debugf("Error occurred while parsing toFloatArray [%v]: %v", strArr, err)
			return nil, false
		}
		floatArr = append(floatArr, i)
	}
	return floatArr, true
}

func _toString(str string) (interface{}, bool) {
	regexStr := "^[A-Za-z0-9_]+$"
	validStr := regexp.MustCompile(regexStr)
	if !validStr.MatchString(str) {
		log.Debugf("Error occurred while parsing toString [%s]: %v", str,
			errors.New(fmt.Sprintf("invalid string format: %s", regexStr)))
		return nil, false
	}
	return str, true
}

func _toInt(str string) (interface{}, bool) {
	//TODO: validate negative numbers
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Debugf("Error occurred while parsing toInt [%s]: %v", str, err)
		return nil, false
	}
	return i, true
}

func _getQuery(key string, req *http.Request) string {
	return req.URL.Query().Get(key)
}

func _parseArr(str string) []string {
	return strings.Split(str, ",")
}
