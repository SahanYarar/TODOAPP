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
	rbd := common.CreateRedisClient()
	redisRepository := repository.CreateRepositoryRedis(rbd)
	userRepository := repository.CreateRepositoryUser(db)
	todoRepository := repository.CreateRepositoryToDo(db)

	middlewareHandler := middleware.CreateMiddlewareHandler(redisRepository)
	userHandler := handler.CreateUserHandler(userRepository, redisRepository)
	todoHandler := handler.CreateToDoHandler(todoRepository)

	r := gin.Default()
	//Get all users
	r.GET("/users", userHandler.GetAllUsers)
	//Sign user
	r.POST("/sign_up", userHandler.SignUser)
	//Login user
	r.POST("/login", userHandler.Login)
	//Logout user
	r.POST("/logout", middlewareHandler.RequireAuth, userHandler.Logout)
	//Testing middleware
	r.GET("/validate", middlewareHandler.RequireAuth, userHandler.ValidateToken)

	//Get user
	r.GET("/user/:id", userHandler.GetUser)
	//Delete user
	r.DELETE("/user/delete/:id", userHandler.DeleteUser)
	//Change password
	r.PATCH("/user/update/:id", middlewareHandler.RequireAuth, userHandler.UpdateUserPassword)

	//Create Todo
	r.POST("/todo/create", todoHandler.CreateToDo)
	//Get all Todos
	r.GET("/todos/", todoHandler.GetAllToDos)
	//Get Todo by id
	r.GET("/todo/:id", todoHandler.GetToDo)
	//Patch Todo
	r.PATCH("/todo/update/:id", todoHandler.UpdateToDo)
	//Delete Todo
	r.DELETE("/todo/delete/:id", todoHandler.DeleteToDo)
	//Test main.go
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})

	})

	r.Run(env.Port)
}
