package models

type Response[T any] struct {
	Status  string `json:"status"`
	Massage string `json:"massage"`
	Data    T      `json:"data,omitempty"`
}
