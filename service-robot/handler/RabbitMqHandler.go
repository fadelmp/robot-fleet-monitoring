package handler

import (
	"encoding/json"
	"log"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/usecase"

	"github.com/streadway/amqp"
)

// Interface
type RabbitMqHandlerContract interface {
	Update(amqp.Delivery) error
}

// Class
type RabbitMqHandler struct {
	usecase usecase.RobotUsecaseContract
}

// Constructor
func NewRabbitMqHandler(usecase usecase.RobotUsecaseContract) *RabbitMqHandler {
	return &RabbitMqHandler{
		usecase: usecase,
	}
}

func (h *RabbitMqHandler) Update(msg amqp.Delivery) {

	var robotDto dto.Robot

	if err := json.Unmarshal(msg.Body, &robotDto); err != nil {

		log.Printf("Failed to unmarshal message body: %v", err)

		return
	}

	SetUsernameFromRabbit(&robotDto)

	h.usecase.Update(robotDto)
}
