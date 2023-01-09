package models

import (
	"time"
	"todoapi/entities"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID         int             `json:"Id"`
	Name       string          `json:"name"`
	Email      string          `json:"email"`
	Todos      []entities.ToDo `json:"Todos"`
	CreatedAt  time.Time       `json:"CreatedAt"`
	UpdatedAt  time.Time       `json:"UpdatedAt"`
	Deleted_at gorm.DeletedAt  `json:"DeletedAt"`
}
