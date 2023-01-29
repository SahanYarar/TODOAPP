package middleware

import (
	"fmt"
	"net/http"
	"todoapi/common"

	"github.com/gin-gonic/gin"
)

type MiddleWareHandler struct {
}

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
			fmt.Println("Burasi calisti")
			return
		}
		defer resp.Body.Close()
		c.Next()
	}

}

/*
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...

*/
