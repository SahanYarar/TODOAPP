package models

type ToDoResponse struct {
	ID        int    `json:"Id"`
	Details   string `json:"details"`
	Status    string `json:"status"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}
