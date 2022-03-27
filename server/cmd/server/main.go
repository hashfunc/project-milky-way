package main

import (
	"github.com/hashfunc/project-milky-way/internal/server"
	"log"

	"github.com/hashfunc/project-milky-way/internal"
)

func main() {
	srv, err := server.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	defer internal.CloseOrPanic(srv)

	log.Fatal(srv.Listen(":3000"))
}
