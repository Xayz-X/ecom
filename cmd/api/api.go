package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Xayz-X/ecom/service/user"
	"github.com/gorilla/mux"
)

// Holds the API server address and database connection
type APIServer struct {
	addr string
	db   *sql.DB
}

// create a new API server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// this is new
func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	// user handler register
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
