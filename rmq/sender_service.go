package rmq

import (
	"context"
	"example/mymodule/requests"
	"example/mymodule/statics"

	"github.com/google/uuid"
	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func SendEmmitBalanceRequest(request requests.EmitBalanceRequest, channel *amqp091.Channel, logger *logrus.Logger) {
	id, err := uuid.NewRandom()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infoln("Generated id to balance-service: ", id.String())
	event := &balances.EmmitBalanceRequest{
		Id:       id.String(),
		Address:  request.Address,
		Currency: request.Currency,
		Amount:   request.Amount,
	}

	logger.Infoln("Generated event to balance-service: ", event.String())

	bytes, _ := proto.Marshal(event)

	err = channel.PublishWithContext(
		context.Background(),
		statics.BalanceExchangeName,
		statics.RkEmmitBalanceRequest,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})

	if err != nil {
		logger.Errorln("Fail to publish message! error: ", err.Error())
	}
}

func SendGetWalletInfoRequest(request *requests.GetBalance, channel *amqp091.Channel, logger *logrus.Logger) string {
	id, _ := uuid.NewRandom()

	event := &balances.GetWalletInfoRequest{
		Id:      id.String(),
		Address: request.Address,
	}

	logger.Infoln("Generated event to balance-service: ", event.String())

	bytes, _ := proto.Marshal(event)

	err := channel.PublishWithContext(
		context.Background(),
		statics.BalanceExchangeName,
		statics.RkGetWalletInfoRequest,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})

	if err != nil {
		logger.Errorln("Fail to publish message! error: ", err.Error())
	}

	return id.String()
}
