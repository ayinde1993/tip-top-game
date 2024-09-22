package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ayinde1993/the-tip-top-game/service/user"
	"github.com/gorilla/mux"
)

// fisrt  to be parameter in the project

type APIServer struct {
	addr string
	db   *sql.DB
}

// create api server instance
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// create Run Methods
func (s *APIServer) Run() error {
	//initialise a router
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//register user http server // end points
	// userStore := user.NewStore(s.db)
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
