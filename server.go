package main

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	port string
	db   *sql.DB
}

func NewServer(port string, db *sql.DB) *Server {
	return &Server{
		port: port,
		db:   db,
	}
}

func (s *Server) Serve() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server online"))
	})

	itemsRouter := NewItemsRouter(mux, s.db)
	itemsRouter.InitRoutes()

	loggedMux := NewLogger(mux)

	server := http.Server{
		Addr:    s.port,
		Handler: loggedMux,
	}

	log.Printf("Server running: http://localhost%s", s.port)
	return server.ListenAndServe()
}
