package repository

import (
	"todoapi/entities"

	"gorm.io/gorm"
)

type UserRepositoryDatabase struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(u *entities.User) error
	GetAllUsers(u []*entities.User) ([]*entities.User, error)
	GetUser(id uint64) (*entities.User, error)
	UpdateUser(u *entities.User) error
	DeleteUser(id uint64) error
}

func CreateRepositoryUser(db *gorm.DB) *UserRepositoryDatabase {
	return &UserRepositoryDatabase{db}
}

func (userRepository *UserRepositoryDatabase) CreateUser(u *entities.User) error {
	return userRepository.db.Create(&u).Error
}

func (userRepository *UserRepositoryDatabase) GetAllUsers(u []*entities.User) ([]*entities.User, error) {

	err := userRepository.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, err
}

func (userRepository *UserRepositoryDatabase) GetUser(id uint64) (*entities.User, error) {
	user := &entities.User{ID: id}
	err := userRepository.db.First(&user).Error

	if err != nil {
		return nil, err
	}
	return user, err
}

func (userRepository *UserRepositoryDatabase) UpdateUser(u *entities.User) error {
	return userRepository.db.Model(&u).Where("id = ?", &u.ID).Updates(&u).Error
}

func (userRepository *UserRepositoryDatabase) DeleteUser(id uint64) error {

	todoID := &entities.User{ID: id}

	return userRepository.db.Delete(todoID).Error
}