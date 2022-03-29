package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (connection *Connection) QueryStars(longitude, latitude float64, distance int) ([]*Star, error) {
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

	cursor, err := connection.DB().
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
