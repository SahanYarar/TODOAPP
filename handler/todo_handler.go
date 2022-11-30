package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"todoapi/entities"
	"todoapi/models"

	"todoapi/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ToDoHandler struct {
	ToDoRepository repository.ToDoRepositoryInterface
}

func CreateHandler(ToDoRepo repository.ToDoRepositoryInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: ToDoRepo}

}

func (handler *ToDoHandler) CreateToDo(c *gin.Context) {
	var payload = &models.ToDoRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(payload.Details)

	createToDo := &entities.ToDo{
		Details: payload.Details,
		Status:  payload.Status,
	}

	err := handler.ToDoRepository.Create(createToDo)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}

	c.JSON(http.StatusCreated, createToDo)
}

func (handler *ToDoHandler) GetAllToDos(c *gin.Context) {

	var u []*entities.ToDo
	/*if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}*/

	ToDos, err := handler.ToDoRepository.GetAll(u)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, ToDos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	ToDoID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	ToDo, err := handler.ToDoRepository.Get(ToDoID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, ToDo)
}

func (handler *ToDoHandler) UpdateToDo(c *gin.Context) {
	var payload = &models.ToDoRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UpdatedToDo := &entities.ToDo{
		Details: payload.Details,
		Status:  payload.Status,
	}

	ToDoID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	ToDo, err := handler.ToDoRepository.Update(ToDoID, UpdatedToDo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}
	c.JSON(http.StatusCreated, ToDo)

}
