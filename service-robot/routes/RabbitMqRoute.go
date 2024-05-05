package routes

import (
	"log"

	"github.com/streadway/amqp"
)

func RabbitMqRoute(ch *amqp.Channel) amqp.Queue {

	q, err := ch.QueueDeclare(
		"robot", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return q
}
