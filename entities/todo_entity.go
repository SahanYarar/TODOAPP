package entities

import (
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	ID          uint64         `gorm:"primaryKey"`
	Description string         `gorm:"column:description"`
	Status      string         `gorm:"column:status"`
	UserID      uint64         `gorm:"column:user_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
