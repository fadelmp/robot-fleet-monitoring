package routes

import (
	"log"

	"robot-fleet-monitoring/service-robot/handler"

	"github.com/streadway/amqp"
)

func RabbitMqRoute(ch *amqp.Channel, rabbit *handler.RabbitMqHandler) {

	// Consume messages
	msgs, err := ch.Consume("robot", "", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process incoming messages
	for msg := range msgs {
		// Handle each message in a separate goroutine
		go rabbit.Update(msg)
	}
}
