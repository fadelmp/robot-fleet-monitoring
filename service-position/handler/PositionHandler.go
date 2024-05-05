package handler

import (
	"robot-fleet-monitoring/service-position/dto"
	"robot-fleet-monitoring/service-position/usecase"

	"github.com/labstack/echo/v4"
)

// Interface
type PositionHandlerContract interface {
	Update(e echo.Context) error
}

// Class
type PositionHandler struct {
	usecase usecase.PositionUsecaseContract
}

// Constructor
func NewPositionHandler(usecase usecase.PositionUsecaseContract) *PositionHandler {
	return &PositionHandler{
		usecase: usecase,
	}
}

// @Summary Update an existing robot
// @Description Update an robot in the robot machine by its ID
// @Success 200 {object} dto.Response
func (h *PositionHandler) Update(e echo.Context) error {

	var positionDto dto.Position

	if e.Bind(&positionDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Update(positionDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, "", "")
}
