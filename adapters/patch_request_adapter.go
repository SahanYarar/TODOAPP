package adapters

import (
	"todoapi/entities"
	"todoapi/models"
)

func CreateFromRequest(r *models.ToDoPatchRequest, todoID uint64) *entities.ToDo {
	payload := &entities.ToDo{
		ID:          todoID,
		Description: r.Description,
		Status:      r.Status,
	}
	return payload

}
