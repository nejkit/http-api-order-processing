package rmq

import (
	proto "github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	googleProto "google.golang.org/protobuf/proto"
)

func InitRabbit() *amqp091.Channel {
	connection, err := amqp091.Dial("amqp://admin:admin@localhost:5670")

	if err != nil {
		panic(err.Error())
	}

	ch, err := connection.Channel()

	if err != nil {
		panic(err.Error())
	}

	return ch
}

func PublishMessage(request proto.EmmitBalanceRequest) {
	ch := InitRabbit()
	body, err := googleProto.Marshal(&request)
	if err != nil {
		return
	}

	ch.Publish(
		"e.balances.forward",
		"r.balances.#.EmmitBalanceRequest.#",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}
