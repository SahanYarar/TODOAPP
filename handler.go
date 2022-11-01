package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User_Handler struct {
	UserRepository UserRepository_interface
}

func New_Hand(UserRep UserRepository_interface) *User_Handler {
	return &User_Handler{UserRepository: UserRep}

}

func (hand *User_Handler) New_User(c *gin.Context) {
	payload := &UserRequest{}

	new_user := &user{
		Name:   payload.Name,
		Status: payload.Status,
	}

	c.JSON(http.StatusCreated, new_user)
}
