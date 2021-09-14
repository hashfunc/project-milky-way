package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const projectName = "Project Milky Way"

func main() {
	config := fiber.Config{
		AppName:      projectName,
		ServerHeader: projectName,
	}

	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(projectName)
	})

	log.Fatal(app.Listen(":3000"))
}
