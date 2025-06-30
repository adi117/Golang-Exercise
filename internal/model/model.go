package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

type PaginatedResponse[T any] struct {
	Items       []T   `json:"items"`
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"totalPages"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
}
