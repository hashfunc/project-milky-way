package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/internal/config"
	"github.com/hashfunc/project-milky-way/internal/db"
)

func main() {
	serverConfig, err := config.LoadConfigFile()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Connect(&serverConfig.Database); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Connection.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := fiber.New(
		fiber.Config{
			AppName:      serverConfig.Name,
			ServerHeader: serverConfig.Name,
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(serverConfig.Name)
	})

	log.Fatal(app.Listen(":3000"))
}
