package kafka

import (
	"context"
	"encoding/json"
	"strconv"
	"todoapi/common"
	"todoapi/models"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func Produce(topic string, user *models.UserResponse) {
	env := common.GetEnvironment()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{env.KafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	userData, err := json.Marshal(user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		return
	}
	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(strconv.Itoa(int(user.ID))),
			Value: userData,
		})

	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}
}
