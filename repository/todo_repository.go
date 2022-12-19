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
	Update(u *entities.ToDo) error
	Delete(id uint64) error
}

func CreateRepositoryToDo(db *gorm.DB) *ToDoRepositoryPostgres {

	return &ToDoRepositoryPostgres{db}

}

func (todoRepository *ToDoRepositoryPostgres) Create(u *entities.ToDo) error {

	return todoRepository.db.Create(&u).Error
}

func (todoRepository *ToDoRepositoryPostgres) GetAll(u []*entities.ToDo) ([]*entities.ToDo, error) {

	err := todoRepository.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, err

}
func (todoRepository *ToDoRepositoryPostgres) Get(id uint64) (*entities.ToDo, error) {
	todo := &entities.ToDo{ID: id}
	err := todoRepository.db.First(&todo).Error

	if err != nil {
		return nil, err
	}
	return todo, err

}

func (todoRepository *ToDoRepositoryPostgres) Update(u *entities.ToDo) error {
	return todoRepository.db.Model(&u).Where("id = ?", &u.ID).Save(&u).Error
}

func (todoRepository *ToDoRepositoryPostgres) Delete(id uint64) error {

	todoID := &entities.ToDo{ID: id}

	return todoRepository.db.Delete(todoID).Error
}
