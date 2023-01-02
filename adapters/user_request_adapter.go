package adapters

import (
	"todoapi/entities"
	"todoapi/models"
)

func CreateFromUserRequest(r *models.UserRequest) *entities.User {
	payload := &entities.User{
		Name:  r.Name,
		Email: r.Email,
	}
	return payload

}
