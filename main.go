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
	//Sign user
	r.POST("/sign_up", userHandler.SignUser)
	//Get user
	r.GET("/user/:id", userHandler.GetUser)
	//Delete user
	r.DELETE("/user/delete/:id", userHandler.DeleteUser)
	//Change password
	r.PATCH("/user/update/:id", middleware.AuthMiddleware(), userHandler.UpdateUserPassword)

	//Create Todo
	r.POST("/todo/create", todoHandler.CreateToDo)
	//Get all Todos
	r.GET("/todos/", todoHandler.GetAllToDos)
	//Get Todo by id
	r.GET("/todo/:id", todoHandler.GetToDo)
	//Patch Todo
	r.PATCH("user/:userid/todo/update/:todoid", middleware.AuthMiddleware(), todoHandler.UpdateToDo)
	//Delete Todo
	r.DELETE("user/:userid/todo/delete/:todoid", middleware.AuthMiddleware(), todoHandler.DeleteToDo)
	//Test main.go
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(env.Port)
}
