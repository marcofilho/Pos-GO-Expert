package main

import (
	"github.com/marcofilho/Pos-GO-Expert/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, World!"),
	}

	go rabbitmq.Producer(channel, msg.Body, "amq.direct")
}
