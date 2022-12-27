package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password"`
	TodoID    uint64         `gorm:"column:todo_id"`
	Todo      []ToDo         `gorm:"foreignKey:TodoID"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
