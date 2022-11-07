package handler

import (
	"fmt"
	"net/http"

	"todoapi/entities"
	"todoapi/models"

	"todoapi/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type User_Handler struct {
	UserRepository repository.UserRepository_Interface
}

func New_Handler(UserRepo repository.UserRepository_Interface) *User_Handler {
	return &User_Handler{UserRepository: UserRepo}

}

func (handler *User_Handler) Create_User(c *gin.Context) {
	var payload = &models.UserRequest{}
	fmt.Println(payload.Name)

	new_user := &entities.User{
		Name:   payload.Name,
		Status: payload.Status,
	}

	err := handler.UserRepository.Create_User_Db(new_user)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}

	c.JSON(http.StatusCreated, new_user)
}
