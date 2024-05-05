package usecase

import (
	"log"

	"github.com/streadway/amqp"
)

// Interface
type ChannelUsecaseContract interface {
	Publish([]byte, string)
}

// Class
type ChannelUsecase struct {
	Channel *amqp.Channel
}

// Constructor
func NewChannelUsecase(Channel *amqp.Channel) *ChannelUsecase {
	return &ChannelUsecase{
		Channel: Channel,
	}
}

// Implementation

func (c *ChannelUsecase) Publish(body []byte, queue string) {

	// Publish DTO data to RabbitMQ
	err := c.Channel.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		log.Fatalf("Failed to publish DTO data to RabbitMQ: %v", err)
	}
}
