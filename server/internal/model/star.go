package model

type GeoJSON struct {
	Type        string        `bson:"type"`
	Coordinates []interface{} `bson:"coordinates"`
}

type Star struct {
	Code        string   `bson:"code"`
	Name        string   `bson:"name"`
	BCode       string   `bson:"b_code"`
	Address     string   `bson:"address"`
	RoadAddress string   `bson:"road_address"`
	Longitude   float64  `bson:"longitude"`
	Latitude    float64  `bson:"latitude"`
	Point       *GeoJSON `bson:"point"`
}
