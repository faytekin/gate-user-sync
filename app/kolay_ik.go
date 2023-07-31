package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var kolayIKBaseUrl = "https://api.kolayik.com/v2/"

type PersonIds struct {
	Id string `json:"id"`
}

type PersonListData struct {
	Total       int         `json:"total"`
	PerPage     int         `json:"perPage"`
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	Items       []PersonIds `json:"items"`
}

type PersonListResponse struct {
	Error bool           `json:"error"`
	Data  PersonListData `json:"data"`
}

type BulkViewRequest struct {
	PersonIDs []string `json:"person_ids"`
}

type Persons struct {
	FirstName   string `json:"firstName"`
	ID          string `json:"id"`
	LastName    string `json:"lastName"`
	MobilePhone string `json:"mobilePhone"`
	WorkPhone   string `json:"workPhone"`
}

type BulkViewResponseData struct {
	Persons []Persons `json:"persons"`
}

type BulkViewResponse struct {
	Error bool                 `json:"error"`
	Data  BulkViewResponseData `json:"data"`
}

func getKolayIkPersonPhoneList(status string) ([]string, error) {
	activePersons, err := getKolayIKPersonList(status)
	if err != nil {
		return nil, fmt.Errorf("kolay IK PersonIds List Failed %w", err)
	}

	var phoneNumbers []string
	for _, person := range activePersons {
		formattedPhoneNumber := person.getFormattedPhone()

		if formattedPhoneNumber != "" {
			phoneNumbers = append(phoneNumbers, formattedPhoneNumber)
		}
	}

	//for _, person := range activePersons {
	//	phoneNumber := person.GetFormattedPhoneNumber()
	//
	//	if phoneNumber == "" || len(phoneNumber) < 11 {
	//		fmt.Println("Name:", person.FirstName, person.LastName)
	//		fmt.Println("Phone:", phoneNumber)
	//		fmt.Println("--------------------------------")
	//	}
	//}

	return phoneNumbers, nil
}

func getKolayIKPersonList(status string) ([]Persons, error) {
	if os.Getenv("KOLAY_IK_TOKEN") == "" {
		panic("Please put KOLAY_IK_TOKEN to .env file")
	}

	personIds, err := getPersonIds(status)
	if err != nil {
		return nil, fmt.Errorf("person ids List Failed %w", err)
	}

	personList, err := getPersons(personIds)
	if err != nil {
		return nil, fmt.Errorf("failed to get person phone numbers %w", err)
	}

	return personList, nil
}

func getPersonIds(status string) ([]PersonIds, error) {
	url := kolayIKBaseUrl + "person/list?status=" + status
	var allPeople []PersonIds

	body, err := sendAPIRequest("POST", url, os.Getenv("KOLAY_IK_TOKEN"), nil)
	if err != nil {
		return nil, fmt.Errorf("api request failed => %w", err)
	}

	var response PersonListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response %w", err)
	}

	allPeople = append(allPeople, response.Data.Items...)

	for page := 2; page <= response.Data.LastPage; page++ {
		pageURL := fmt.Sprintf("%s&page=%d", url, page)

		body, err = sendAPIRequest("POST", pageURL, os.Getenv("KOLAY_IK_TOKEN"), nil)
		if err != nil {
			return nil, fmt.Errorf("failed KolayIK request %w", err)
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed KolayIK body parse %w", err)
		}

		allPeople = append(allPeople, response.Data.Items...)
	}

	return allPeople, nil
}

func getPersons(personIds []PersonIds) ([]Persons, error) {
	url := kolayIKBaseUrl + "person/bulk-view"

	var ids []string
	for _, person := range personIds {
		ids = append(ids, person.Id)
	}

	requestData := BulkViewRequest{PersonIDs: ids}
	body, err := sendAPIRequest("POST", url, os.Getenv("KOLAY_IK_TOKEN"), requestData)
	if err != nil {
		return nil, err
	}

	var response BulkViewResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Persons, nil
}

func (p *Persons) getFormattedPhone() string {
	var phone string

	if p.WorkPhone != "" && len(p.WorkPhone) > 4 {
		phone = p.WorkPhone
	} else if p.MobilePhone != "" {
		phone = p.MobilePhone
	}

	formattedPhoneNumber := strings.ReplaceAll(phone, " ", "")
	formattedPhoneNumber = strings.ReplaceAll(formattedPhoneNumber, "+", "")

	return formattedPhoneNumber
}
