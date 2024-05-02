package config

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func InitRabbit() *amqp.Channel {

	rabbitUser := os.Getenv("RABBITMQ_USERNAME")
	rabbitPass := os.Getenv("RABBITMQ_PASSWORD")
	rabbitHost := os.Getenv("RABBITMQ_HOST")
	rabbitPort := os.Getenv("RABBITMQ_PORT")

	conn, err := amqp.Dial("amqp://" + rabbitUser + ":" + rabbitPass + "@" + rabbitHost + ":" + rabbitPort + "/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ:", err)
	}

	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error creating channel:", err)
	}

	defer ch.Close()

	return ch
}
