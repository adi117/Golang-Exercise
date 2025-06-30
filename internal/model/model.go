package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}
