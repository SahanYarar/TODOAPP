package main

import (
	"todoapi/common"
	"todoapi/handler"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	env := common.Get_Environment()
	db := common.ConnectDatabase(env.DatabaseUrl)
	ToDoRepository := repository.CreateRepositoryToDo(db)

	ToDoHandler := handler.CreateHandler(ToDoRepository)

	r := gin.Default()
	r.POST("/todo/create", ToDoHandler.CreateToDo)
	r.GET("/todo/getall", ToDoHandler.GetAllToDos)
	r.GET("/todo/:id", ToDoHandler.GetToDo)
	r.PUT("/todo/update/:id", ToDoHandler.UpdateToDo)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(":9920")
}
