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
	KafkaURL    string
	KafkaBroker string
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
	kafkaUrl := os.Getenv("KafkaURL")
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	return &Environment{
		DatabaseUrl: databaseUrl,
		Port:        port,
		AuthURL:     authUrl,
		KafkaURL:    kafkaUrl,
		KafkaBroker: kafkaBroker,
	}
}
