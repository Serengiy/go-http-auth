package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	Message    string      `json:"message"`
	Code       int         `json:"code"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalPages int64 `json:"total_pages"`
}

func PaginatedJSON(w http.ResponseWriter, code int, data interface{}, message string, pagination *Pagination) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := Response{
		Data: data,
		Meta: Meta{
			Message:    message,
			Code:       code,
			Pagination: pagination,
		},
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func JSON(w http.ResponseWriter, code int, data interface{}, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := Response{
		Data: data,
		Meta: Meta{
			Message: message,
			Code:    code,
		},
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}
