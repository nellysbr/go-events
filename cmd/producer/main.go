package main

import "github.com/nellysbr/go-events/pkg/rabbitmq"

func main() {

	// abre um canal com o rabbitmq
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, World!", "amq.direct")
}
