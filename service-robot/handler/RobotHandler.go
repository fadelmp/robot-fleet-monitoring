package handler

import (
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/message"
	"robot-fleet-monitoring/service-robot/usecase"

	"github.com/labstack/echo/v4"
)

// Interface
type RobotHandlerContract interface {
	GetAll(e echo.Context) error
	GetById(e echo.Context) error

	Create(e echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
}

// Class
type RobotHandler struct {
	usecase usecase.RobotUsecaseContract
}

// Constructor
func NewRobotHandler(usecase usecase.RobotUsecaseContract) *RobotHandler {
	return &RobotHandler{
		usecase: usecase,
	}
}

// @Summary Get All robot
// @Description Retrieve a list of all robot
// @Success 200 {object} dto.Response
func (h *RobotHandler) GetAll(e echo.Context) error {

	robotDtos := h.usecase.GetAll()

	// Check Value
	if len(robotDtos) == 0 {
		return NotFound(e, message.NotFound)
	}

	// Return Success
	return Success(e, message.GetSuccess, robotDtos)
}

// @Summary Get robot by ID
// @Description Retrieve an robot from the robot machine by its ID
// @Success 200 {object} dto.Response
func (h *RobotHandler) GetById(e echo.Context) error {

	var robot dto.Robot

	// Bind Value
	if e.Bind(&robot) != nil {
		return BadRequest(e)
	}

	// Parse Id
	if ParseId(e, &robot) != nil {
		return BadRequest(e)
	}

	// Get By Id
	robotDto := h.usecase.GetById(robot.Id)

	// Check Value
	if robotDto.Id == "" {
		return NotFound(e, message.NotFound)
	}

	// Return Success
	return Success(e, message.GetSuccess, robotDto)
}

// @Summary Create a new robot
// @Description Add a new robot to the robot machine
// @Success 200 {object} dto.Response
func (h *RobotHandler) Create(e echo.Context) error {

	var robotDto dto.Robot
	SetUsername(&robotDto, e)

	if e.Bind(&robotDto) != nil {
		return BadRequest(e)
	}

	if Validate(&robotDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Create(robotDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.CreateSuccess, "")
}

// @Summary Update an existing robot
// @Description Update an robot in the robot machine by its ID
// @Success 200 {object} dto.Response
func (h *RobotHandler) Update(e echo.Context) error {

	var robotDto dto.Robot
	SetUsername(&robotDto, e)

	if e.Bind(&robotDto) != nil {
		return BadRequest(e)
	}

	if Validate(&robotDto) != nil {
		return BadRequest(e)
	}

	if ParseId(e, &robotDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Update(robotDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.UpdateSuccess, "")
}

// @Summary Delete an dto.robot
// @Description Remove an dto.robot from the robot machine by its ID
// @Success 200 {object} dto.Response
func (h *RobotHandler) Delete(e echo.Context) error {

	var robotDto dto.Robot
	SetUsername(&robotDto, e)

	if e.Bind(&robotDto) != nil {
		return BadRequest(e)
	}

	if ParseId(e, &robotDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Delete(robotDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.DeleteSuccess, "")
}
