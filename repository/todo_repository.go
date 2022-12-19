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
	Update(id uint64, u *entities.ToDo) error
	Delete(id uint64) error
}

func CreateRepositoryToDo(db *gorm.DB) *ToDoRepositoryPostgres {

	return &ToDoRepositoryPostgres{db}

}

func (ToDoRepository *ToDoRepositoryPostgres) Create(u *entities.ToDo) error {

	return ToDoRepository.db.Create(&u).Error
}

func (ToDoRepository *ToDoRepositoryPostgres) GetAll(u []*entities.ToDo) ([]*entities.ToDo, error) {

	err := ToDoRepository.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, err

}
func (ToDoRepository *ToDoRepositoryPostgres) Get(id uint64) (*entities.ToDo, error) {
	ToDo := &entities.ToDo{ID: id}
	err := ToDoRepository.db.First(&ToDo).Error

	if err != nil {
		return nil, err
	}
	return ToDo, err

}

func (ToDoRepository *ToDoRepositoryPostgres) Update(id uint64, u *entities.ToDo) error {
	ToDo := &entities.ToDo{ID: id}

	return ToDoRepository.db.Model(ToDo).Updates(&u).Error

}

func (ToDoRepository *ToDoRepositoryPostgres) Delete(id uint64) error {

	ToDo := &entities.ToDo{ID: id}

	return ToDoRepository.db.Delete(ToDo).Error
}
