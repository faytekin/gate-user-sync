package ik

import (
	"encoding/json"
	"fmt"
	"gate-user-sync/helper"
	"gate-user-sync/models"
)

var apiConfig *helper.IKAPI

func SetConfig(config *helper.Config) {
	apiConfig = &config.IKApi
}

func GetPhoneList(status string) ([]string, error) {
	activePersons, err := getPersonList(status)
	if err != nil {
		return nil, err
	}

	var phoneNumbers []string
	for _, person := range activePersons {
		formattedPhoneNumber := person.GetFormattedPhone()

		if formattedPhoneNumber != "" {
			phoneNumbers = append(phoneNumbers, formattedPhoneNumber)
		}
	}

	return phoneNumbers, nil
}

func getPersonList(status string) ([]models.Persons, error) {
	personIds, err := getPersonIds(status)
	if err != nil {
		return nil, err
	}

	personList, err := getPersons(personIds)
	if err != nil {
		return nil, err
	}

	return personList, nil
}

func getPersonIds(status string) ([]models.PersonIds, error) {
	url := fmt.Sprintf("%s/person/list?status=%s", apiConfig.Url, status)
	var allPeople []models.PersonIds

	body, err := helper.SendAPIRequest("POST", url, apiConfig.BearerToken, nil)
	if err != nil {
		return nil, err
	}

	var response models.PersonListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response %w", err)
	}

	allPeople = append(allPeople, response.Data.Items...)

	for page := 2; page <= response.Data.LastPage; page++ {
		pageURL := fmt.Sprintf("%s&page=%d", url, page)

		body, err = helper.SendAPIRequest("POST", pageURL, apiConfig.BearerToken, nil)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed KolayIK body parse %w", err)
		}

		allPeople = append(allPeople, response.Data.Items...)
	}

	return allPeople, nil
}

func getPersons(personIds []models.PersonIds) ([]models.Persons, error) {
	url := fmt.Sprintf("%s/person/bulk-view", apiConfig.Url)

	var ids []string
	for _, person := range personIds {
		ids = append(ids, person.Id)
	}

	requestData := models.BulkViewRequest{PersonIDs: ids}
	body, err := helper.SendAPIRequest("POST", url, apiConfig.BearerToken, requestData)
	if err != nil {
		return nil, err
	}

	var response models.BulkViewResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Persons, nil
}
