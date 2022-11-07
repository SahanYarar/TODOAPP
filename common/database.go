package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect_Database() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=8085 user=your_user dbname=deneme_01 password=toor")

	if err != nil {
		fmt.Println(err)
	}
	return db
}
