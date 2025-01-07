package main

import (
	"main/internal/db"
	"main/internal/server"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	db := db.InitDb()
	defer db.Close()

	server := server.NewServer(port, db)
	server.Serve()

}
