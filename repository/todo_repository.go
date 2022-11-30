package repository

import (
	"todoapi/entities"

	"gorm.io/gorm"
)

type ToDoRepositoryPostgres struct {
	db *gorm.DB
}

type ToDoRepositoryInterface interface {
	Create(u *entities.ToDo) error
	GetAll(u []*entities.ToDo) ([]*entities.ToDo, error)
	Get(id uint64) (*entities.ToDo, error)
	Update(id uint64, u *entities.ToDo) (*entities.ToDo, error)
}

func CreateRepositoryToDo(db *gorm.DB) *ToDoRepositoryPostgres {

	return &ToDoRepositoryPostgres{db}

}

func (connect *ToDoRepositoryPostgres) Create(u *entities.ToDo) error {

	err := connect.db.Create(&u).Error

	if err != nil {
		return err
	}
	return nil
}

func (connect *ToDoRepositoryPostgres) GetAll(u []*entities.ToDo) ([]*entities.ToDo, error) {

	err := connect.db.Find(&u).Error

	if err != nil {
		return nil, err
	}
	return u, err

}
func (connect *ToDoRepositoryPostgres) Get(id uint64) (*entities.ToDo, error) {
	ToDo := &entities.ToDo{ID: id}
	err := connect.db.First(&ToDo).Error

	if err != nil {
		return nil, err
	}
	return ToDo, err

}

func (connect *ToDoRepositoryPostgres) Update(id uint64, u *entities.ToDo) (*entities.ToDo, error) {
	ToDo := &entities.ToDo{ID: id}
	err := connect.db.Model(ToDo).Updates(&u).Error

	if err != nil {
		return nil, err
	}
	return ToDo, err

}
