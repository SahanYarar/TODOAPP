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

func CreateFromToDoPatchRequest(c *entities.ToDo, r *models.ToDoPatchRequest) *entities.ToDo {
	if &r.Status != nil {
		c.Status = r.Status
	}
	if &r.Description != nil {
		c.Description = r.Description
	}
	return c
}

/*

type ToDoAdapterInterface interface {
	CreateFromToDoRequest(r *models.ToDoRequest) *entities.ToDo
	CreateFromToDoPatchRequest(r *models.ToDoPatchRequest *AdapterToDo
}

type AdapterToDo struct{
	adapterToDo *entities.ToDo
}


func CreateFromToDoRequest(r *models.ToDoRequest) *entities.ToDo {
	payload := &entities.ToDo{
		Description: r.Description,
		Status:      r.Status,
	}
	return payload

}


func (c *AdapterToDo)CreateFromToDoPatchRequest(r *models.ToDoPatchRequest) *AdapterToDo{
	if &r.Status != nil {
		c.Status = r.Status
	}
	if &r.Description != nil {
		c.Description = r.Description
	}
	return c
}


**************************************************
func CreateFromToDoPatchRequest(r *entities.ToDo, p *models.ToDoPatchRequest) *entities.ToDo {
	payload := &entities.ToDo{
		ID:          todoID,
		Description: r.Description,
		Status:      r.Status,
	}
	return payload}



func CreateFromToDoPatchRequest(r *models.ToDoPatchRequest, todoID uint64) *entities.ToDo {
	payload := &entities.ToDo{
		ID:          todoID,
		Description: r.Description,
		Status:      r.Status,
	}
	return payload}

2 versions 1-Todo objesini pass by ref. alarak güncellemek
2-Adaptörü receiver func haline getirip güncelleme yapmak
İkisinide yap */
