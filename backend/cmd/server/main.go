package main

import (
	"log"

	"github.com/hashfunc/project-milky-way/internal/server"
)

func main() {
	srv, err := server.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	srv.Run()
}
