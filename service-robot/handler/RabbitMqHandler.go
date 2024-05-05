package handler

import (
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/usecase"

	"github.com/labstack/echo/v4"
)

// Interface
type RabbitMqHandlerContract interface {
	UpdateFromRabbit(e echo.Context) error
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

func (h *RobotHandler) UpdateFromRabbit(robotDto dto.Robot) {

	SetUsernameFromRabbit(&robotDto)

	h.usecase.Update(robotDto)
}
