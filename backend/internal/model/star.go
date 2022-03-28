package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/hashfunc/project-milky-way/internal/db"
)

type GeoJSON struct {
	Type        string        `bson:"type"`
	Coordinates []interface{} `bson:"coordinates"`
}

type Star struct {
	Code        string   `bson:"code" json:"-"`
	Name        string   `bson:"name" json:"name"`
	BCode       string   `bson:"b_code" json:"-"`
	Address     string   `bson:"address" json:"address"`
	RoadAddress string   `bson:"road_address" json:"road_address"`
	Longitude   float64  `bson:"longitude" json:"longitude"`
	Latitude    float64  `bson:"latitude" json:"latitude"`
	Point       *GeoJSON `bson:"point" json:"-"`
}

func QueryStars(longitude, latitude float64, distance int) ([]*Star, error) {
	filter := bson.D{{
		"point", bson.D{{
			"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", bson.A{longitude, latitude}},
				}},
				{"$maxDistance", distance},
			},
		}},
	}}

	cursor, err := db.Connection.DB().
		Collection("stars").
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	stars := new([]*Star)

	if err := cursor.All(context.TODO(), stars); err != nil {
		return nil, err
	}

	return *stars, nil
}
