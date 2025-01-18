package rabbitmq

import ampq "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*ampq.Channel, error) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// abrindo canal com rabbitmq

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil

}

// fazemos apenas auto-ack quando podemos perder a msg
func Consume(ch *ampq.Channel, out chan<- ampq.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,         // queue
		"go-consumer", // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg // jogando a msg no canal
	}
	return nil
}

func Publish(ch *ampq.Channel, msg string, exchange string) error {
	err := ch.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return err
	}
	return nil
}
