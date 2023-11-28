package rmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AmqpSender struct {
	channel *amqp091.Channel
	rk      string
	ex      string
}

func (s *AmqpSender) SendMessage(ctx context.Context, message protoreflect.ProtoMessage) {
	body, _ := proto.Marshal(message)
	err := s.channel.PublishWithContext(ctx, s.ex, s.rk, false, false, amqp091.Publishing{ContentType: "text/plain", Body: body})
	if err != nil {
		logger.Errorln("Fail publish a message! error: ", err.Error())
		return
	}
	logger.Infoln("Succesfully send event to rk: %s, ex: %s", s.rk, s.ex)
}
