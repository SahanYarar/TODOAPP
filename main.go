package main

import (
	"todoapi/common"
	"todoapi/handler"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	Env := common.Get_Environment()
	Db := common.ConnectDatabase(Env.DatabaseUrl)
	ToDoRepository := repository.CreateRepositoryToDo(Db)

	ToDoHandler := handler.CreateHandler(ToDoRepository)

	r := gin.Default()
	r.POST("/todo/create", ToDoHandler.CreateToDo)
	r.GET("/todos/", ToDoHandler.GetAllToDos)
	r.GET("/todo/:id", ToDoHandler.GetToDo)
	r.PATCH("/todo/:id", ToDoHandler.UpdateToDo)
	r.DELETE("/todo/:id", ToDoHandler.DeleteToDo)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(Env.Port)
}
