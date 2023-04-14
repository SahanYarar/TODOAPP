package adapters

import (
	"todoapi/entities"
	"todoapi/models"
)

func CreateFromUserEntities(u *entities.User) *models.UserResponse {

	var response = &models.UserResponse{
		ID:            int(u.ID),
		Name:          u.Name,
		Email:         u.Email,
		Todos:         u.Todos,
		IsEmailActive: u.IsEmailActive,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		Deleted_at:    u.DeletedAt,
	}
	return response
}
