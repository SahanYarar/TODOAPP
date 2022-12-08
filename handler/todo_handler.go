package handler

import (
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

	createToDo := &entities.ToDo{
		Description: payload.Description,
		Status:      payload.Status,
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
	ToDos, err := handler.ToDoRepository.GetAll(u)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if ToDos == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"massage": "ToDos not exists",
		})
		return
	}
	c.JSON(http.StatusOK, ToDos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	ToDoID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	ToDo, err := handler.ToDoRepository.Get(ToDoID)
	if ToDo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	c.JSON(http.StatusOK, ToDo)
}

func (handler *ToDoHandler) UpdateToDo(c *gin.Context) {
	var payload = &models.ToDoPatchRequest{}

	ToDoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	ValidateToDoId, err := handler.ToDoRepository.Get(ToDoID)
	if ValidateToDoId == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	UpdatedToDo := &entities.ToDo{
		Description: payload.Description,
		Status:      payload.Status,
	}

	ToDo, err := handler.ToDoRepository.Update(ToDoID, UpdatedToDo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}
	c.JSON(http.StatusCreated, ToDo)

}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	ToDoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	ValidateToDoId, err := handler.ToDoRepository.Get(ToDoID)
	if ValidateToDoId == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	err = handler.ToDoRepository.Delete(ToDoID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "ToDo deleted!",
	})

}
