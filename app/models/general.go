package models

type StandardGetRequest struct {
	Name   string `json:"name" query:"name"`
	Status string `json:"status" query:"name"`
	Page   int    `json:"page" query:"page"`
	Limit  int    `json:"limit" query:"limit"`
}

type StandardResponse struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type StandardResponseWithPaginate struct {
	StandardResponse
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
