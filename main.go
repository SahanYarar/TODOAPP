package main

import (
	"todoapi/common"
	"todoapi/handler"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	env := common.GetEnvironment()
	db := common.ConnectDatabase(env.DatabaseUrl)
	userRepository := repository.CreateRepositoryUser(db)
	userHandler := handler.CreateUserHandler(userRepository)
	todoRepository := repository.CreateRepositoryToDo(db)
	todoHandler := handler.CreateToDoHandler(todoRepository)

	r := gin.Default()
	r.POST("/user/create", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)

	r.POST("/todo/create", todoHandler.CreateToDo)
	r.GET("/todos/", todoHandler.GetAllToDos)
	r.GET("/todo/:id", todoHandler.GetToDo)
	r.PATCH("/todo/:id", todoHandler.UpdateToDo)
	r.DELETE("/todo/:id", todoHandler.DeleteToDo)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(env.Port)
}
