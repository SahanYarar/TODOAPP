package utils

import (
	"time"
	"todoapi/common"
	"todoapi/entities"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func GenerateJWTToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	env := common.GetEnvironment()
	secret_key := env.Secret
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return "massage", err
	}
	return tokenString, err
}
