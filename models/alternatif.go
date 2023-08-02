package models

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
