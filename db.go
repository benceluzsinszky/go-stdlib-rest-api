package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func db() {
	connStr := "postgres://postgres:password@localhost:5432/playground?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	defer db.Close()

	fmt.Println("Connected to database")

}
