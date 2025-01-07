package server

import (
	"database/sql"
	"log"
	"main/internal/middlewares"
	"main/internal/routers"
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

	itemsRouter := routers.NewItemsRouter(mux, s.db)
	itemsRouter.InitRoutes()

	loggedMux := middlewares.NewLogger(mux)

	server := http.Server{
		Addr:    s.port,
		Handler: loggedMux,
	}

	log.Printf("Server running: http://localhost%s", s.port)
	return server.ListenAndServe()
}
