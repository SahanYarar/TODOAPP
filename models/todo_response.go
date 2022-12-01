package models

type ToDoResponse struct {
	ID          int    `json:"Id"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
	Deleted_at  string `json:"DeletedAt"`
}
