package api

import (
	"database/sql"
	"net/http"

	"github.com/claudineyveloso/rest-api.git/services/healthy"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()
	healthy.RegisterRoutes(r)
	return http.ListenAndServe("localhost:8080", r)
}
