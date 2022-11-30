package entities

import "time"

type ToDo struct {
	ID        uint64    `gorm:"primaryKey"`
	Details   string    `gorm:"column:details"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
