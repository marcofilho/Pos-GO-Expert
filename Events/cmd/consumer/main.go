package main

import (
	"fmt"

	"github.com/marcofilho/Pos-GO-Expert/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(channel, msgs, "my-queue")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
