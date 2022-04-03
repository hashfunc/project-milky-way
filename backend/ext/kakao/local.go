package kakao

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type LocalKeywordDocument struct {
	ID                string `json:"id"`
	PlaceName         string `json:"place_name"`
	CategoryName      string `json:"category_name"`
	CategoryGroupCode string `json:"category_group_code"`
	CategoryGroupName string `json:"category_group_name"`
	Phone             string `json:"phone"`
	AddressName       string `json:"address_name"`
	RoadAddressName   string `json:"road_address_name"`
	X                 string `json:"x"`
	Y                 string `json:"y"`
	PlaceURL          string `json:"place_url"`
	Distance          string `json:"distance"`
}

type LocalKeywordResponse struct {
	Meta      Meta                   `json:"meta"`
	Documents []LocalKeywordDocument `json:"documents"`
}

func (client *Client) LocalKeyword(query string) (*LocalKeywordResponse, error) {
	agent := fiber.AcquireAgent()

	request := agent.Request()
	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(client.config.URL + "/v2/local/search/keyword.json")
	request.Header.Set("Authorization", "KakaoAK "+client.config.Key)

	agent.QueryString("query=" + url.QueryEscape(query))

	if err := agent.Parse(); err != nil {
		return nil, err
	}

	status, body, err := agent.Bytes()

	if err != nil {
		return nil, err[0]
	}

	if status != 200 {
		println(string(body))
		return nil, errors.New("카카오 키워드 검색 오류")
	}

	response := new(LocalKeywordResponse)

	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
