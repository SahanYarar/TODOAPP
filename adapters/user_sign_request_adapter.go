package adapters

import (
	"todoapi/entities"
	"todoapi/models"
	"todoapi/utils"
)

func CreateFromUserSignRequest(r *models.UserSignRequest) *entities.User {
	hasedPassword := utils.HashPassword(r.Password)
	payload := &entities.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: hasedPassword,
	}
	return payload

}
