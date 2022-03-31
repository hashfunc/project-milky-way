package server

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/ext/kakao"
	"github.com/hashfunc/project-milky-way/internal"
	"github.com/hashfunc/project-milky-way/internal/config"
	"github.com/hashfunc/project-milky-way/internal/database"
)

type Server struct {
	app      *fiber.App
	db       *database.Connection
	config   *config.Config
	validate *validator.Validate
	kakao    *kakao.Client
}

type (
	Handler  func(*fiber.Ctx) error
	Register func(*Server) Handler
)

func (server *Server) RegisterAPI(register Register) Handler {
	return register(server)
}

func (server *Server) Register() {
	group := server.app.Group("api/v1")
	group.Get("/stars", server.RegisterAPI(GetStars))
	group.Get("/search/keyword", server.RegisterAPI(SearchKeyword))
}

func (server *Server) Run() {
	defer internal.CloseOrPanic(server.db)

	db, err := database.New(&server.config.Database)

	if err != nil {
		log.Fatal(err)
	}

	server.db = db

	bind := server.config.Bind

	if bind == "" {
		bind = ":8080"
	}

	if err := server.app.Listen(bind); err != nil {
		log.Fatal(err)
	}
}

func NewServer() (*Server, error) {
	cfg, err := config.LoadConfigFile()

	if err != nil {
		return nil, err
	}

	app := fiber.New(
		fiber.Config{
			AppName: cfg.Name,
		},
	)

	server := &Server{
		app:      app,
		config:   cfg,
		validate: validator.New(),
		kakao:    kakao.NewClient(cfg.Secret.Kakao),
	}

	server.Register()

	return server, nil
}
