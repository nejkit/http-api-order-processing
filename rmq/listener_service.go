package rmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type ParserFunc[T any] func([]byte) (*T, error)
type GetMessageId[T any] func(*T) string

type Listener[T any] struct {
	logger     *logrus.Logger
	rmqChannel *amqp091.Channel
	messages   <-chan amqp091.Delivery
	parser     ParserFunc[T]
	idFunc     GetMessageId[T]
}

func NewListener[T any](factory AmqpFactory, qName string, parser ParserFunc[T], idFunc GetMessageId[T]) Listener[T] {
	channel := factory.GetChannel()
	messages, err := channel.Consume(qName, "", false, false, false, false, nil)
	if err != nil {
		factory.logger.Errorln("FAil created consumer! error: ", err.Error())
	}
	return Listener[T]{
		logger:     factory.logger,
		rmqChannel: channel,
		messages:   messages,
		parser:     parser,
		idFunc:     idFunc,
	}
}

func (l *Listener[T]) ConsumeById(ctx context.Context, id string) *T {
	for msg := range l.messages {
		message, err := l.parser(msg.Body)
		if err != nil {
			l.logger.Infoln("Bad message, skipping...")
			msg.Nack(false, false)
			continue
		}
		if l.idFunc(message) == id {
			msg.Ack(false)
			return message
		}
	}
	return nil
}
