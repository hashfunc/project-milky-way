package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/hashfunc/project-milky-way/internal/db"
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

	filter := bson.D{
		{"point", bson.D{
			{"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", bson.A{longitude, latitude}},
				}},
				{"$maxDistance", 2000},
			}},
		}},
	}

	cursor, err := db.Connection.DB().
		Collection("stars").
		Find(context.TODO(), filter)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	stars := new([]*model.Star)
	if err := cursor.All(context.TODO(), stars); err != nil {
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
