package rmq

import (
	"context"
	"fmt"

	proto "github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	googleProto "google.golang.org/protobuf/proto"
)

func InitRabbit() *amqp091.Channel {
	connection, err := amqp091.Dial("amqp://admin:admin@localhost:5672")

	if err != nil {
		panic(err.Error())
	}

	ch, err := connection.Channel()

	if err != nil {
		panic(err.Error())
	}

	return ch
}

func SendEventGetWalletInfo(request *proto.GetWalletInfoRequest) {
	ch := InitRabbit()
	body, _ := googleProto.Marshal(request)

	ch.PublishWithContext(
		context.Background(),
		"e.balances.forward",
		"r.balances.http-api.GetWalletInfoRequest.#",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func PublishMessage(request *proto.EmmitBalanceRequest) {
	ch := InitRabbit()
	body, err := googleProto.Marshal(request)
	if err != nil {
		return
	}

	ch.PublishWithContext(
		context.Background(),
		"e.balances.forward",
		"r.balances.#.EmmitBalanceRequest.#",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func CatchResponseWalletInfo(id string) *proto.WalletInfo {
	ch := InitRabbit()
	listener, _ := ch.Consume(
		"q.balances.response.GetWalletInfoResponse",
		"",
		false,
		false,
		false,
		false,
		nil)

	for msg := range listener {
		var request proto.GetWalletInfoResponse
		err := googleProto.Unmarshal(msg.Body, &request)
		if err != nil {
			fmt.Println(err.Error())
		}
		if request.Id == id {
			err = msg.Ack(false)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(request.String())
			return request.WalletInfo
		} else {
			err = msg.Nack(false, true)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return nil
}
