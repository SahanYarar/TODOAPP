package utils

import (
	"os"
	"time"
	"todoapi/entities"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func GenerateJWTToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return "massage", err
	}
	secret_key := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return "massage", err
	}
	return tokenString, err
}
