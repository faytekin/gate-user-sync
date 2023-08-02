package alternatif

import (
	"encoding/json"
	"fmt"
	"gate-user-sync/helper"
	"gate-user-sync/models"
	"strings"
)

var apiConfig *helper.AlternatifAPI

func SetConfig(config *helper.Config) {
	apiConfig = &config.AlternatifApi
}

func AddNewUsers(phones []string) error {
	response, err := addNewUsersRequest(phones)
	if err != nil {
		return err
	}

	if response.Added > 0 {
		helper.Log(fmt.Sprintf("Out of %d records, %d were added. %d have already been added", response.Total, response.Added, response.PreviouslyAdded))
	}

	if response.Failed > 0 {
		failNumbers := strings.Join(response.FailNumbers, ", ")
		helper.Log(fmt.Sprintf("%d numbers couldn't be added. Unadded numbers are: %s", response.Failed, failNumbers))
	}

	return nil
}

func RemoveUsers(phones []string) error {
	response, err := removeUsersRequest(phones)
	if err != nil {
		return err
	}

	if response {
		helper.Log(fmt.Sprintf("%d users have been successfully removed from the group.", len(phones)))
	} else {
		helper.Log("User removal from group failed.")
	}

	return nil
}

func addNewUsersRequest(phones []string) (models.AddUserResponseData, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", apiConfig.Url, apiConfig.GroupId)

	defaultReturn := models.AddUserResponseData{}
	data := models.RequestData{Phones: phones}
	body, err := helper.SendAPIRequest("POST", url, apiConfig.BearerToken, data)
	if err != nil {
		return defaultReturn, err
	}

	var response models.AddUserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return defaultReturn, err
	}

	return response.Data, nil
}

func removeUsersRequest(phones []string) (bool, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", apiConfig.Url, apiConfig.GroupId)

	data := models.RequestData{Phones: phones}
	body, err := helper.SendAPIRequest("DELETE", url, apiConfig.BearerToken, data)
	if err != nil {
		return false, err
	}

	var response models.GeneralSuccessResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	return response.Data.Success, nil
}
