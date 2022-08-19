package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/svilenkomitov/rental-service/internal/rental-service/handler"
	"github.com/svilenkomitov/rental-service/internal/rental-service/repository"
	"github.com/svilenkomitov/rental-service/internal/storage"
)

const (
	FetchRentalByIdEndpoint = "/rentals/{id}"
)

type Server struct {
	server *http.Server
}

func NewServer(c *Config, db *storage.Database) *Server {
	r := setUpRouting(db)
	server := setUpServer(c, r)
	return &Server{
		server: server,
	}
}

func setUpRouting(db *storage.Database) *mux.Router {
	router := mux.NewRouter()

	repository := repository.NewRepository(db)
	handler := handler.NewHandler(repository)

	router.NewRoute().
		Methods(http.MethodGet).
		Path(FetchRentalByIdEndpoint).
		HandlerFunc(handler.FetchRentalById)

	return router
}

func setUpServer(c *Config, router *mux.Router) *http.Server {
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(int(c.Port)),
		Handler: router,
	}
	return server
}

func (s *Server) Start() error {
	log.Info("Starting the HTTP server at addr: ", s.server.Addr)
	if err := s.server.ListenAndServe(); nil != err && err != http.ErrServerClosed {
		log.Errorf("Failed to start server: %v", err)
	}
	return nil
}
