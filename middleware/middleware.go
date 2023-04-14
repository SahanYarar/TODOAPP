package middleware

import (
	"net/http"
	"todoapi/common"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		env := common.GetEnvironment()
		url := env.AuthURL
		client := &http.Client{}
		reqToken := c.Request.Header.Get("Authorization")
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Authorization", reqToken)
		resp, _ := client.Do(req)

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "Not authorized"})
			return
		}
		defer resp.Body.Close()
		c.Next()
	}

}
