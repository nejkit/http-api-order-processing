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
	q, err := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil)

	body := "hello world"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain]",
			Body:        []byte(body),
		})

}
