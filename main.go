package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	defer db.Close()

	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	server := NewAPIServer(port, db)

	server.Serve()

}
