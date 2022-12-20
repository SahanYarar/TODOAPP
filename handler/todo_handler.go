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
	todos, err := handler.ToDoRepository.GetAll(u)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if todos == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"massage": "ToDos not exists",
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	todo, err := handler.ToDoRepository.Get(todoID)
	if todo == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (handler *ToDoHandler) UpdateToDo(c *gin.Context) {
	var payload = &models.ToDoPatchRequest{}

	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkToDo, err := handler.ToDoRepository.Get(todoID)
	if checkToDo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	payloadToDo := adapters.CreateFromRequest(payload, todoID)

	if payloadToDo.Status == " " {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
			"Status field cannot be empty if it's given in json!!": payloadToDo.Status,
		})
		return
	}

	if payloadToDo.Description == " " {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
			"Description field cannot be empty if given in json!!": payloadToDo.Description,
		})
		return
	}

	err = handler.ToDoRepository.Update(payloadToDo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}

	updatedToDo, err := handler.ToDoRepository.Get(todoID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, updatedToDo)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	checkToDo, err := handler.ToDoRepository.Get(todoID) //isimde sıkıntı
	if checkToDo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	err = handler.ToDoRepository.Delete(todoID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "ToDo deleted!",
	})

}
