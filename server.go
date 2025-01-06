package main

import (
	"database/sql"
	"log"
	"net/http"
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

func (s *APIServer) Serve() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server online"))
	})

	itemsRouter := NewItemsRouter(mux, s.db)
	itemsRouter.InitRoutes()

	server := http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	log.Printf("Server running: http://localhost%s", s.addr)

	return server.ListenAndServe()
}
