package main

import (
	"todoapi/common"
	"todoapi/handler"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	env := common.Get_Environment()
	db := common.Connect_Database(env.DatabaseUrl)
	User_Repository := repository.New_Repo_User(db)

	User_Handler := handler.New_Handler(User_Repository)

	r := gin.Default()
	r.POST("/user/create", User_Handler.Create_User)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	r.Run(":9920")
}
