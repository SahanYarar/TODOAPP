package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Environment struct {
	DatabaseUrl   string
	Port          string
	Secret        string
	RedisAddr     string
	RedisPassword string
}

func GetEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return nil
	}

	databaseUrl := os.Getenv("DNS")
	port := os.Getenv("Port")
	secret_key := os.Getenv("SECRET")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	return &Environment{
		DatabaseUrl:   databaseUrl,
		Port:          port,
		Secret:        secret_key,
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
	}
}
