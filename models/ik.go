package models

import "strings"

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

func (p *Persons) GetFormattedPhone() string {
	var phone string

	if p.MobilePhone != "" && len(p.MobilePhone) > 4 {
		phone = p.MobilePhone
	} else if p.WorkPhone != "" && len(p.WorkPhone) > 4 {
		phone = p.WorkPhone
	}

	formattedPhoneNumber := strings.ReplaceAll(phone, " ", "")
	formattedPhoneNumber = strings.ReplaceAll(formattedPhoneNumber, "+", "")

	return formattedPhoneNumber
}
