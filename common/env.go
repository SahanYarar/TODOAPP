package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Environment struct {
	DatabaseUrl string
	Port        string
}

func GetEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return nil
	}

	databaseUrl := os.Getenv("DNS")
	port := os.Getenv("Port")

	return &Environment{
		DatabaseUrl: databaseUrl,
		Port:        port,
	}
}
