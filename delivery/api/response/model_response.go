package response

import "github.com/jutionck/golang-db-sinar-harapan-makmur/model/dto"

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponseCreated struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type PagedResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data,omitempty"`
	Paging dto.Paging    `json:"paging,omitempty"`
}
