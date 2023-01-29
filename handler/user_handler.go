package handler

import (
	"net/http"
	"strconv"
	"todoapi/adapters"
	"todoapi/entities"
	"todoapi/models"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func CreateUserHandler(userRepository repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

func (handler *UserHandler) SignUser(c *gin.Context) {
	var payload = &models.UserSignRequest{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Validate() == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!!!",
		})
		return
	}
	newUser := adapters.CreateFromUserSignRequest(payload)
	err := handler.UserRepository.CreateUser(newUser)
	userResponse := adapters.CreateFromUserEntities(newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (handler *UserHandler) GetAllUsers(c *gin.Context) {

	var u []*entities.User
	user, err := handler.UserRepository.GetAllUsers(u)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if user == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users not exists",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	user, err := handler.UserRepository.GetUserByID(userID)
	if user == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not exists!",
		})
		return
	}
	userResponse := adapters.CreateFromUserEntities(user)
	c.JSON(http.StatusOK, userResponse)
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkUser, err := handler.UserRepository.GetUserByID(userID)
	if checkUser == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not exists!",
		})
		return
	}

	err = handler.UserRepository.DeleteUser(userID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "User deleted!",
	})

}

func (handler *UserHandler) UpdateUserPassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkToDo, err := handler.UserRepository.GetUserByID(userID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var payload *models.UserPasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	if payload.Validate() == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password did not match or it is  nil!",
		})
		return
	}

	updatedUser := adapters.CreateFromUserPasswordRequest(checkToDo, payload)
	err = handler.UserRepository.UpdateUserPassword(updatedUser)
	c.JSON(http.StatusCreated, checkToDo)
}
