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

func CreateToDoHandler(todoRepository repository.ToDoRepositoryInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: todoRepository}

}

func (handler *ToDoHandler) CreateToDo(c *gin.Context) {
	var payload = &models.ToDoRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createToDo := adapters.CreateFromToDoRequest(payload)
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
			"message": "ToDos not exists",
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

// 403 issue
//
//	ilk çözüm /user/:userid/todo/:todoid burdaki :todoid iyle işlem yap userid yi validate et
//	ikinci çözüm  middleware parçalara ayır tokendeki useridyi çek tododa verilenle karşılaştır
func (handler *ToDoHandler) UpdateToDo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userid"), 10, 64)
	todoID, err := strconv.ParseUint(c.Param("todoid"), 10, 64)
	checkToDo, err := handler.ToDoRepository.Get(todoID)
	if checkToDo == nil { //check todo güncelle
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if checkToDo.UserID != userID {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden!!",
		})
		return
	}

	var payload *models.ToDoPatchRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
		return
	}
	updatedToDo := adapters.CreateFromToDoPatchRequest(checkToDo, payload)

	err = handler.ToDoRepository.Update(updatedToDo)
	/*
		if err != nil {
			zap.S().Error("Error: ", zap.Error(err))
			return
		}

		if payloadToDo.Status == " " { //Trim space

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
		}*/

	c.JSON(http.StatusCreated, checkToDo)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userid"), 10, 64)
	todoID, err := strconv.ParseUint(c.Param("todoid"), 10, 64)
	checkToDo, err := handler.ToDoRepository.Get(todoID)
	if checkToDo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ToDo not exists!",
		})
		return
	}

	if checkToDo.UserID != userID {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden!!",
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
