package repository

import (
	"todoapi/entities"

	"gorm.io/gorm"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

type UserRepository_Interface interface {
	Create_User_Db(user *entities.User) error
}

func New_Repo_User(db *gorm.DB) *UserRepositoryPostgres {

	return &UserRepositoryPostgres{db}

}

func (connect *UserRepositoryPostgres) Create_User_Db(user *entities.User) error {

	err := connect.db.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}
