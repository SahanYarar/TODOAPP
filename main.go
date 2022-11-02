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

/*
func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}*/
