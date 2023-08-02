package alternatif

import (
	"encoding/json"
	"fmt"
	"gate-user-sync/helper"
	"gate-user-sync/models"
)

func GetPhoneList() ([]string, error) {
	users, err := getUserList()
	if err != nil {
		return nil, err
	}

	var phoneNumbers []string
	for _, user := range users {
		phoneNumbers = append(phoneNumbers, user.Phone)
	}

	return phoneNumbers, nil
}

func getUserList() ([]models.User, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", apiConfig.Url, apiConfig.GroupId)
	var users []models.User

	body, err := helper.SendAPIRequest("GET", url, apiConfig.BearerToken, nil)
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	users = append(users, response.Data...)

	for page := 2; page <= response.Meta.Paginate.TotalPages; page++ {
		pageURL := fmt.Sprintf("%s?page=%d", url, page)

		body, err = helper.SendAPIRequest("GET", pageURL, apiConfig.BearerToken, nil)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		users = append(users, response.Data...)
	}

	return users, nil
}
