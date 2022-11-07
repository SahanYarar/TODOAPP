package models

type UserRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
