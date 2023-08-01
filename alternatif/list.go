package alternatif

import (
	"encoding/json"
	"fmt"
	"macellan-gate-user-sync/helper"
	"os"
)

var BaseUrl = "https://altpay-backend.test/integration"

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type Paginate struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

type Meta struct {
	Paginate Paginate `json:"paginate"`
}

type Response struct {
	Data []User `json:"data"`
	Meta Meta   `json:"meta"`
}

func GetUserList(groupId string) ([]User, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", BaseUrl, groupId)
	var users []User

	body, err := helper.SendAPIRequest("GET", url, os.Getenv("ALTERNATIF_TOKEN"), nil)
	if err != nil {
		return nil, fmt.Errorf("alternatif api request failed => %w", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response %w", err)
	}

	users = append(users, response.Data...)

	for page := 2; page <= response.Meta.Paginate.TotalPages; page++ {
		pageURL := fmt.Sprintf("%s?page=%d", url, page)

		body, err = helper.SendAPIRequest("GET", pageURL, os.Getenv("ALTERNATIF_TOKEN"), nil)
		if err != nil {
			return nil, fmt.Errorf("failed Alternatif request %w", err)
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed Alternatif body parse %w", err)
		}

		users = append(users, response.Data...)
	}

	return users, nil
}
