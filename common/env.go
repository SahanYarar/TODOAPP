package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Environment struct {
	DatabaseUrl string
	Port        string
	AuthURL     string
}

func GetEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return nil
	}

	databaseUrl := os.Getenv("DNS")
	port := os.Getenv("Port")
	authUrl := os.Getenv("AuthURL")

	return &Environment{
		DatabaseUrl: databaseUrl,
		Port:        port,
		AuthURL:     authUrl,
	}
}
