package server

import (
	"github.com/gofiber/fiber/v2"
)

type GetStarsQueryParams struct {
	Longitude float64 `validate:"required"`
	Latitude  float64 `validate:"required"`
	Distance  int     `validate:"required,min=100,max=5000"`
}

func GetStars(server *Server) Handler {
	return func(ctx *fiber.Ctx) error {
		params := new(GetStarsQueryParams)

		if err := ctx.QueryParser(params); err != nil {
			return BadRequest(ctx, err)
		}

		if err := server.validate.Struct(params); err != nil {
			return BadRequest(ctx, err)
		}

		stars, err := server.db.QueryStars(params.Longitude, params.Latitude, params.Distance)

		if err != nil {
			return InternalServerError(ctx, err)
		}

		return OK(ctx, stars)
	}
}
