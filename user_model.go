package main

type user struct {
	Id     uint64 `gorm:"primaryKey"`
	Name   string
	Status string
}
