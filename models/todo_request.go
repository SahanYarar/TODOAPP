package models

type ToDoRequest struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UserID      uint64 `json:"user_id"`
}
