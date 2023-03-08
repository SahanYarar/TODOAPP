package handler

import (
	"net/http"
	"strconv"
	"todoapi/adapters"
	"todoapi/entities"
	"todoapi/kafka"
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

// @Summary SignUser
// @Description Signs user and starts kafka.produce
// @Produce json
// @Param body body models.UserSignRequest true "User name,email and password"
// @Success      201  {object} models.UserResponse
// @Failure 400
// @Router /sign_up [post]
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

	go kafka.Produce("e-mail", userResponse)

	c.JSON(http.StatusCreated, userResponse)
}

// @Summary GetAllUsers
// @Description Retrieves all users
// @Produce json
// @Success      200  {object} []entities.User
// @Failure 404
// @Router /users [get]
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

// @Summary GetUser
// @Description Retrieves user by id
// @Produce json
// @Success      200 {object} entities.User
// @Failure 404
// @Router /user/{userid} [get]
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

// @Summary DeleteUser
// @Description Deletes user
// @Produce json
// @Success      204
// @Failure 404
// @Router /user/delete/{userid} [delete]
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

// @Summary UpdateUserPassword
// @Description Updates user password by id
// @Security ApiKeyAuth
// @Produce json
// @Param body body models.UserPasswordRequest true "User password,confirmpassword"
// @Success      201  {object} entities.User
// @Failure 400
// @Failure 500
// @Router /user/update/{userid} [patch]
func (handler *UserHandler) UpdateUserPassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkUser, err := handler.UserRepository.GetUserByID(userID)
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

	updatedUser := adapters.CreateFromUserPasswordRequest(checkUser, payload)
	err = handler.UserRepository.UpdateUserPassword(updatedUser)
	c.JSON(http.StatusCreated, checkUser)
}

// @Summary ActivateEmail
// @Description Actives email
// @Produce json
// @Success 201
// @Failure 500
// @Router /activation/{userid} [patch]
func (handler *UserHandler) ActivateEmail(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkUser, err := handler.UserRepository.GetUserByID(userID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	checkUser.IsEmailActive = true
	err = handler.UserRepository.UpdateIsEmailActive(checkUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Your Account is now active!!!",
		"User":    checkUser.IsEmailActive,
	})

}

// @Summary ResetUserPassword
// @Description Resets user password by email
// @Produce json
// @Param body body models.UserResetPasswordRequest true "User email"
// @Success 201
// @Failure 400
// @Failure 500
// @Failure 404
// @Router /resetpassword [patch]
func (handler *UserHandler) UserResetPassword(c *gin.Context) {
	var payload *models.UserResetPasswordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	user, err := handler.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
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
	if user.IsEmailActive == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is not active!!",
		})
		return
	}
	userResponse := adapters.CreateFromUserEntities(user)
	go kafka.Produce("e-mail", userResponse)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Produced !!",
	})
}
