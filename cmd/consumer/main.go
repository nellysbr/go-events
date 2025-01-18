package main

import (
	"fmt"

	"github.com/nellysbr/go-events/pkg/rabbitmq"
	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan ampq.Delivery)
	go rabbitmq.Consume(ch, msgs, "test")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) //mensagem confirmada
	}

}
