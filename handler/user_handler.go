package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todoapi/adapters"
	"todoapi/entities"
	"todoapi/models"
	"todoapi/repository"
	"todoapi/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository  repository.UserRepositoryInterface
	RedisRepository repository.RedisRepositoryInterface
}

func CreateUserHandler(userRepository repository.UserRepositoryInterface, redisRepository repository.RedisRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepository: userRepository,
		RedisRepository: redisRepository,
	}
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

func (handler *UserHandler) Login(c *gin.Context) {
	var payload *models.UserLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	user, err := handler.UserRepository.GetUserByEmail(payload.Email)
	if user == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not exists!",
		})
		return
	}
	if user.Email != payload.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid  password",
		})
		return
	}
	user_jwt, err := utils.GenerateJWTToken(*user)
	stringUserId := strconv.FormatInt(int64(user.ID), 10)
	err = handler.RedisRepository.Set(c, stringUserId, user_jwt, time.Hour)

	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": user_jwt,
	})

}

func (handler *UserHandler) Logout(c *gin.Context) {
	var payload *models.UserLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	user, err := handler.UserRepository.GetUserByEmail(payload.Email)
	if user == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not exists!",
		})
		return
	}
	if user.Email != payload.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid  password",
		})
		return
	}

	stringUserId := strconv.FormatInt(int64(user.ID), 10)
	err = handler.RedisRepository.Delete(c, stringUserId)
	if err != nil {
		fmt.Println("Patlayan delete")
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Logout",
	})
	return
}

func (handler *UserHandler) ValidateToken(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Token is valid",
	})

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
