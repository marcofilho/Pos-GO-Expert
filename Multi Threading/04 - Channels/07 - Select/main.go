package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id      int64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1) //avoid concurrency
			msg := Message{Id: i, Message: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1) //avoid concurrency
			msg := Message{Id: i, Message: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ: ID: %d - Message: %s\n", msg.Id, msg.Message)
		case msg := <-c2:
			fmt.Printf("Received from Kafka: ID: %d - Message: %s\n", msg.Id, msg.Message)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
