package rmq

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

type AmqpFactory struct {
	conn *amqp091.Connection
}

func NewFactory(connString string) AmqpFactory {
	var con *amqp091.Connection
	var err error
	for {
		con, err = amqp091.Dial(connString)
		if err != nil {
			logger.Errorln("Failed to connect rabbit! message: ", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return AmqpFactory{conn: con}
}

func (f *AmqpFactory) GetChannel() *amqp091.Channel {
	channel, err := f.conn.Channel()
	if err != nil {
		logger.Errorln("Fail create a channel! ", err.Error())
		return nil
	}
	return channel
}

func (f *AmqpFactory) NewSender(ex string, rk string) AmqpSender {
	channel := f.GetChannel()
	return AmqpSender{
		channel: channel,
		rk:      rk,
		ex:      ex,
	}
}
