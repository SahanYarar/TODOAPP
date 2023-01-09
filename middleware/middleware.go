package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todoapi/common"
	"todoapi/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MiddlewareHandler struct {
	RedisRepository repository.RedisRepositoryInterface
}

func CreateMiddlewareHandler(redisRepository repository.RedisRepositoryInterface) *MiddlewareHandler {
	return &MiddlewareHandler{
		RedisRepository: redisRepository,
	}
}

func (handler *MiddlewareHandler) RequireAuth(c *gin.Context) {

	env := common.GetEnvironment()

	secret_key := env.Secret

	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		fmt.Println("request header hatasi")
		fmt.Println(len(splitToken))
		c.AbortWithStatus(http.StatusUnauthorized)
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
		convertedId := claims["id"].(float64)
		stringUserId := strconv.FormatFloat(convertedId, 'f', 0, 64)
		isExists := handler.RedisRepository.Exists(c, stringUserId)

		if isExists == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		converted_time := claims["exp"].(float64)
		if converted_time < float64(time.Now().Unix()) {

			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()

	} else {
		fmt.Print("error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
