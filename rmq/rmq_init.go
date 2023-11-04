package rmq

import (
	"github.com/rabbitmq/amqp091-go"
)

func InitRabbit() {
	connection, err := amqp091.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err.Error())
	}
	ch, err := connection.Channel()
	if err != nil {
		panic(err.Error())
	}

	err = ch.ExchangeDeclare(
		"e.balances.forward",
		"topic",
		true,
		false,
		true,
		false,
		amqp091.NewConnectionProperties())

	queueEmmitBalanceRequest, err := ch.QueueDeclare(
		"q.balances-service.EmmitBalanceRequest",
		true,
		false,
		false,
		false,
		amqp091.NewConnectionProperties())

	queueEmmitBalanceResponse, err := ch.QueueDeclare(
		"q.balances-service.EmmitBalanceResponse",
		true,
		false,
		false,
		false,
		amqp091.NewConnectionProperties())

	err = ch.QueueBind(
		queueEmmitBalanceRequest.Name,
		"r.balances-service.EmmitBalanceRequest",
		"e.balances.forward",
		false,
		amqp091.NewConnectionProperties())

	err = ch.QueueBind(
		queueEmmitBalanceResponse.Name,
		"r.balances-service.EmmitBalanceResponse",
		"e.balances.forward",
		false,
		amqp091.NewConnectionProperties())

}
