package rmq

import (
	"context"
	"example/mymodule/statics"

	balances "github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func CreateLisWalletInfoResponce(channel *amqp091.Channel, logger *logrus.Logger) <-chan amqp091.Delivery {
	msgs, err := channel.ConsumeWithContext(
		context.Background(),
		statics.QueueGetWalletInfoResponse,
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		logger.Fatalln("Consume fail! error:", err.Error())
	}

	return msgs
}

func ConsumeWalletInfoResponse(id string, messages <-chan amqp091.Delivery, logger *logrus.Logger) *balances.GetWalletInfoResponse {
	var response balances.GetWalletInfoResponse
	for msg := range messages {
		proto.Unmarshal(msg.Body, &response)
		logger.Info("Response from balance-service:", response.String())
		if response.Id == id {
			msg.Ack(false)
			break
		} else {
			msg.Nack(false, true)
		}
	}
	return &response
}
