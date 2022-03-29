package database

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

type GeoJSON struct {
	Type        string        `bson:"type"`
	Coordinates []interface{} `bson:"coordinates"`
}
