package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

type UserRepository_interface interface {
	Create_User(user *user)
}

func New_Repo_User(db *gorm.DB) *UserRepositoryPostgres {

	return &UserRepositoryPostgres{db}

}

func (con *UserRepositoryPostgres) Create_User(user *user) {

	con.db.Create(&user)

	fmt.Println("Succesfully created")
	return

}
