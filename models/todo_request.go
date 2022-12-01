package models

type ToDoRequest struct {
	Description string `json:"description"`
	Status      string `json:"status"`
}
