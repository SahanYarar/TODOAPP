package main

import (
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

func (con *UserRepositoryPostgres) Create_User(user *user) error {

	err := con.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}
