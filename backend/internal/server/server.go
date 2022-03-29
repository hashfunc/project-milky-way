package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/internal/api"
	"github.com/hashfunc/project-milky-way/internal/config"
	"github.com/hashfunc/project-milky-way/internal/db"
)

type Server struct {
	application *fiber.App
	config      *config.Config
}

func (server *Server) Start() error {
	bind := server.config.Bind
	if bind == "" {
		bind = ":8080"
	}
	return server.application.Listen(bind)
}

func (server *Server) Close() error {
	if err := db.Connection.Close(); err != nil {
		return err
	}
	return nil
}

func NewServer() (*Server, error) {
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
		config:      serverConfig,
	}

	return server, nil
}
