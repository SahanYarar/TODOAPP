package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todoapi/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	env := common.GetEnvironment()

	secret_key := env.Secret

	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		fmt.Println("request header hatasi")
		fmt.Println(len(splitToken))
		return

	}
	jwtToken := splitToken[1]
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			fmt.Println("Token time hatasi")
			return
		}
		c.Next()

	} else {
		fmt.Print("error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
