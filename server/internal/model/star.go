package model

type Star struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	BCode       string `json:"b_code"`
	Address     string `json:"address"`
	RoadAddress string `json:"road_address"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
}
