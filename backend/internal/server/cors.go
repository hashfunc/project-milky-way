package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/hashfunc/project-milky-way/internal/config"
)

func corsHandler(cfg *config.CORSConfig) fiber.Handler {
	corsConfig := cors.Config{
		AllowHeaders:     cfg.AllowHeaders,
		AllowCredentials: cfg.AllowCredentials,
		ExposeHeaders:    cfg.ExposeHeaders,
		MaxAge:           cfg.MaxAge,
	}

	if cfg.AllowOrigins != "" {
		corsConfig.AllowOrigins = cfg.AllowOrigins
	}

	if cfg.AllowMethods != "" {
		corsConfig.AllowMethods = cfg.AllowMethods
	}

	return cors.New(corsConfig)
}
