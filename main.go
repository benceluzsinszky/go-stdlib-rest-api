package main

import (
	"os"

	_ "github.com/lib/pq"
)

func main() {

	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	db := InitDb()
	defer db.Close()

	server := NewServer(port, db)
	server.Serve()

}
