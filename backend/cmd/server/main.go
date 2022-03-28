package main

import (
	"log"

	"github.com/hashfunc/project-milky-way/internal"
	"github.com/hashfunc/project-milky-way/internal/server"
)

func main() {
	srv, err := server.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	defer internal.CloseOrPanic(srv)

	log.Fatal(srv.Listen(":3000"))
}
