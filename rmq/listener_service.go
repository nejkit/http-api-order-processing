package rmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

type ParserFunc[T any] func([]byte) (*T, error)
type GetMessageId[T any] func(*T) string

type Listener[T any] struct {
	rmqChannel *amqp091.Channel
	messages   <-chan amqp091.Delivery
	parser     ParserFunc[T]
	idFunc     GetMessageId[T]
}

func NewListener[T any](factory AmqpFactory, qName string, parser ParserFunc[T], idFunc GetMessageId[T]) Listener[T] {
	channel := factory.GetChannel()
	messages, err := channel.Consume(qName, "", false, false, false, false, nil)
	if err != nil {
		logger.Errorln("FAil created consumer! error: ", err.Error())
	}
	return Listener[T]{
		rmqChannel: channel,
		messages:   messages,
		parser:     parser,
		idFunc:     idFunc,
	}
}

func (l *Listener[T]) ConsumeById(ctx context.Context, id string) *T {
	for msg := range l.messages {
		logger.Infoln("try consume response")
		message, err := l.parser(msg.Body)
		if err != nil {
			logger.Errorln("Fail parse msg. Reazon: ", err.Error())
		}
		if l.idFunc(message) == id {
			msg.Ack(false)
			logger.Println("Message receive")
			return message
		}
	}
	return nil
}
