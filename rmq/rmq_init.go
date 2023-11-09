package rmq

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func GetChannel(connectionString string, logger *logrus.Logger) *amqp091.Channel {
	con, err := amqp091.Dial(connectionString)
	if err != nil {
		logger.Fatalln("Connection is failed")
	}
	channel, _ := con.Channel()
	return channel
}
