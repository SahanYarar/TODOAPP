package main

import (
	"todoapi/common"
	"todoapi/handler"
	"todoapi/middleware"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	env := common.GetEnvironment()
	db := common.ConnectDatabase(env.DatabaseUrl)
	userRepository := repository.CreateRepositoryUser(db)
	todoRepository := repository.CreateRepositoryToDo(db)

	userHandler := handler.CreateUserHandler(userRepository)
	todoHandler := handler.CreateToDoHandler(todoRepository)

	r := gin.Default()
	//Get all users
	r.GET("/users", userHandler.GetAllUsers)
	//Login user
	r.POST("/login", userHandler.Login)
	//Sign user
	r.POST("/sign_up", userHandler.SignUser)
	r.GET("/validate", middleware.RequireAuth, userHandler.ValidateToken)

	//Create Todo
	r.POST("/todo/create", todoHandler.CreateToDo)
	//Get all Todos
	r.GET("/todos/", todoHandler.GetAllToDos)
	//Get Todo by id
	r.GET("/todo/:id", todoHandler.GetToDo)
	//Patch Todo
	r.PATCH("/todo/:id", todoHandler.UpdateToDo)
	//Delete Todo
	r.DELETE("/todo/:id", todoHandler.DeleteToDo)
	//Test main.go
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(env.Port)
}
