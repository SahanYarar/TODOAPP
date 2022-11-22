package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect_Database(DatabaseUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(DatabaseUrl), &gorm.Config{})

	if err != nil {
		return nil
	}
	return db
}
