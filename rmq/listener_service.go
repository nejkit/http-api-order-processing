package rmq

import (
	"context"
	"example/mymodule/statics"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
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
