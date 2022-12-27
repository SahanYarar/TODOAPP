package adapters

import (
	"todoapi/entities"
	"todoapi/models"
)

func CreateFromToDoRequest(r *models.ToDoRequest) *entities.ToDo {
	payload := &entities.ToDo{
		Description: r.Description,
		Status:      r.Status,
	}
	return payload

}

func CreateFromToDoPatchRequest(r *entities.ToDo, p *models.ToDoPatchRequest) *entities.ToDo {
	payload := &entities.ToDo{
		ID:          todoID,
		Description: r.Description,
		Status:      r.Status,
	}
	return payload
}

/*
func CreateFromToDoPatchRequest(r *models.ToDoPatchRequest, todoID uint64) *entities.ToDo {
	payload := &entities.ToDo{
		ID:          todoID,
		Description: r.Description,
		Status:      r.Status,
	}
	return payload

} 2 versions 1-Todo objesini pass by ref. alarak güncellemek
2-Adaptörü receiver func haline getirip güncelleme yapmak
İkisinide yap */
