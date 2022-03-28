package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/internal/model"
)

type GetStarsQueryParams struct {
	Longitude float64 `validate:"required"`
	Latitude  float64 `validate:"required"`
	Distance  int     `validate:"required,min=100,max=5000"`
}

func GetStars(ctx *fiber.Ctx) error {
	params := new(GetStarsQueryParams)

	if err := ctx.QueryParser(params); err != nil {
		return BadRequest(ctx, err)
	}

	if err := validate.Struct(params); err != nil {
		return BadRequest(ctx, err)
	}

	stars, err := model.QueryStars(params.Longitude, params.Latitude, params.Distance)

	if err != nil {
		return InternalServerError(ctx, err)
	}

	return OK(ctx, stars)
}
