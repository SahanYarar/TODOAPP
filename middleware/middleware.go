package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func RequireAuth(c *gin.Context) {

	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}
	secret_key := os.Getenv("SECRET")

	//get token off req
	jwtToken, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("Patlayan jwt tokeni çekme")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	/*bearerToken := c.Request.Header.Get("Authorization") Anlamadım sor
	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		fmt.Println("  request header hatasi")
		return
	}
	jwtToken := splitToken[1]*/

	//Decode/validate it
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check exp
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
