package alternatif

import (
	"encoding/json"
	"fmt"
	"gate-user-sync/helper"
	"os"
	"strings"
)

type RequestData struct {
	Phones []string `json:"phones"`
}

type AddUserResponseData struct {
	Total           int      `json:"total"`
	Added           int      `json:"added"`
	PreviouslyAdded int      `json:"previously_added"`
	Failed          int      `json:"failed"`
	FailNumbers     []string `json:"fail_numbers"`
}

type AddUserResponse struct {
	Data AddUserResponseData `json:"data"`
}

type GeneralSuccessResponse struct {
	Data struct {
		Success bool `json:"success"`
	} `json:"data"`
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

func addNewUsersRequest(phones []string) (AddUserResponseData, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", BaseUrl, os.Getenv("ALTERNATIF_USER_GROUP_ID"))

	defaultReturn := AddUserResponseData{}
	data := RequestData{Phones: phones}
	body, err := helper.SendAPIRequest("POST", url, os.Getenv("ALTERNATIF_TOKEN"), data)
	if err != nil {
		return defaultReturn, err
	}

	var response AddUserResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return defaultReturn, err
	}

	return response.Data, nil
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

func removeUsersRequest(phones []string) (bool, error) {
	url := fmt.Sprintf("%s/user/group/%s/users", BaseUrl, os.Getenv("ALTERNATIF_USER_GROUP_ID"))

	data := RequestData{Phones: phones}
	body, err := helper.SendAPIRequest("DELETE", url, os.Getenv("ALTERNATIF_TOKEN"), data)
	if err != nil {
		return false, err
	}

	var response GeneralSuccessResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	return response.Data.Success, nil
}
