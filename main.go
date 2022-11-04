package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db := con_database()
	User_Repo := New_Repo_User(db)

	User_Hand := New_Hand(User_Repo)

	r := gin.Default()
	r.POST("/user/crate", User_Hand.New_User)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	r.Run(":8085")
}
