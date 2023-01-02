package handler

import (
	"net/http"
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

	newUser := adapters.CreateFromUserRequest(payload)
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
			"massage": "Users not exists",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) Login(c *gin.Context) {
	//request iste payloada bak
	var payload *models.UserLoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	//response get yaz onla var mı yok mu kontrol et
	user, err := handler.UserRepository.GetUserByEmail(payload.Email)
	if user == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not exists!",
		})
		return
	}
	if user.Email != payload.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Invalid email",
		})
		return
	}
	// yollanan şifreyi haslı şifre ile karşılaştır

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Invalid  password",
		})
		return
	}
	//JWT token yarat
	user_jwt, err := utils.GenerateJWTToken(*user)
	//JWT tokenı geri yolla
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", user_jwt, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *UserHandler) ValidateToken(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Massage": "its working",
	})

}
