package models

type UserResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
