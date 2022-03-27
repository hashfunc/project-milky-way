package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/internal/api"
	"github.com/hashfunc/project-milky-way/internal/config"
	"github.com/hashfunc/project-milky-way/internal/db"
)

type Server struct {
	application *fiber.App
}

func (server *Server) Listen(address string) error {
	return server.application.Listen(address)
}

func (server *Server) Close() error {
	if err := db.Connection.Close(); err != nil {
		return err
	}
	return nil
}

func New() (*Server, error) {
	serverConfig, err := config.LoadConfigFile()

	if err != nil {
		return nil, err
	}

	err = db.Connect(&serverConfig.Database)

	if err != nil {
		return nil, err
	}

	application := fiber.New(
		fiber.Config{
			AppName:      serverConfig.Name,
			ServerHeader: serverConfig.Name,
		})

	api.Route(application)

	server := &Server{
		application: application,
	}

	return server, nil
}
