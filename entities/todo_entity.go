package entities

type User struct {
	Id     uint64 `gorm:"primaryKey"`
	Name   string
	Status string
}
