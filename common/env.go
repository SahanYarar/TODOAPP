package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Environment struct {
	DatabaseUrl string
}

func Get_Environment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return nil
	}

	database_url := os.Getenv("DNS")

	return &Environment{
		DatabaseUrl: database_url,
	}
}
