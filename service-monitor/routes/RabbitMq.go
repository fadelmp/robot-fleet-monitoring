package routes

import (
	"log"

	"robot-fleet-monitoring/service-monitor/handler"

	"github.com/streadway/amqp"
)

func RabbitMqRoute(ch *amqp.Channel, rabbit *handler.MonitorHandler) {

	// Consume messages
	msgs, err := ch.Consume("monitor", "", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process incoming messages
	for msg := range msgs {
		// Handle each message in a separate goroutine
		go rabbit.Consume(msg)
	}
}
