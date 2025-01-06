package main

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	port string
	db   *sql.DB
}

func NewAPIServer(port string, db *sql.DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

func (s *APIServer) Serve() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server online"))
	})

	itemsRouter := NewItemsRouter(mux, s.db)
	itemsRouter.InitRoutes()

	server := http.Server{
		Addr:    s.port,
		Handler: mux,
	}

	log.Printf("Server running: http://localhost%s", s.port)

	return server.ListenAndServe()
}
