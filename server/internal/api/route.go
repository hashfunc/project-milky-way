package api

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App) {
	routeApi(app)
}

func routeApi(app *fiber.App) {
	apiGroup := app.Group("/api")
	routeVersion1(&apiGroup)
}

func routeVersion1(router *fiber.Router) {
	v1Group := (*router).Group("/v1")
	v1Group.Get("/stars/:location", GetStars)
}
