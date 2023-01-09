package adapters

import (
	"todoapi/entities"
	"todoapi/models"
	"todoapi/utils"
)

func CreateFromUserPasswordRequest(c *entities.User, r *models.UserPasswordRequest) *entities.User {
	hasedPassword := utils.HashPassword(r.Password)
	if &r.Password != nil {
		c.Password = hasedPassword
	}
	return c
}
