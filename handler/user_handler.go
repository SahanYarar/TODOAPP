package handler

import (
	"net/http"
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

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var payload = &models.UserRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createUser := adapters.CreateFromUserRequest(payload)
	err := handler.UserRepository.CreateUser(createUser)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}

	c.JSON(http.StatusCreated, createUser)
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

/*func Register(c *gin.Context){

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!"})

}*/
