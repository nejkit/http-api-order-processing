package rmq

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func GetChannel(con *amqp091.Connection, logger *logrus.Logger) *amqp091.Channel {
	channel, _ := con.Channel()
	return channel
}

func GetConnection(connectionString string, logger *logrus.Logger) (*amqp091.Connection, error) {
	for attempt := 0; attempt < 5; attempt++ {
		con, err := amqp091.Dial(connectionString)
		if err != nil {
			if attempt == 4 {
				return nil, err
			} else {
				logger.Warnln("Connect failed. Message: ", err.Error())
				continue
			}
		}
		logger.Infoln("Connection succesful!")
		return con, nil
	}
	return nil, nil
}
