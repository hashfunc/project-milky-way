package api

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/hashfunc/project-milky-way/internal/model"
)

func GetStars(ctx *fiber.Ctx) error {
	location := ctx.Params("location")
	if location == "" {
		return fiber.ErrBadRequest
	}

	axis := strings.Split(location, ",")
	if len(axis) != 2 {
		return fiber.ErrBadRequest
	}

	longitude, err := strconv.ParseFloat(axis[0], 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	latitude, err := strconv.ParseFloat(axis[1], 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	stars, err := model.QueryStars(longitude, latitude, 2000)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(
		&DefaultResponse{
			Code:    "OK",
			Message: "OK",
			Data:    &stars,
		},
	)
}
