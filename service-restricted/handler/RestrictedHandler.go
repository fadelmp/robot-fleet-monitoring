package handler

import (
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/message"
	"robot-fleet-monitoring/service-restricted/usecase"

	"github.com/labstack/echo/v4"
)

// Interface
type RestrictedHandlerContract interface {
	GetAll(e echo.Context) error
	GetById(e echo.Context) error

	Create(e echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
}

// Class
type RestrictedHandler struct {
	usecase usecase.RestrictedUsecaseContract
}

// Constructor
func NewRestrictedHandler(usecase usecase.RestrictedUsecaseContract) *RestrictedHandler {
	return &RestrictedHandler{
		usecase: usecase,
	}
}

// @Summary Get All Restricted
// @Description Retrieve a list of all Restricted
// @Success 200 {object} dto.Response
func (h *RestrictedHandler) GetAll(e echo.Context) error {

	restrictedDtos := h.usecase.GetAll()

	// Check Value
	if len(restrictedDtos) == 0 {
		return NotFound(e, message.NotFound)
	}

	// Return Success
	return Success(e, message.GetSuccess, restrictedDtos)
}

// @Summary Get Restricted by ID
// @Description Retrieve an Restricted from the Restricted machine by its ID
// @Success 200 {object} dto.Response
func (h *RestrictedHandler) GetById(e echo.Context) error {

	var restricted dto.Restricted

	// Bind Value
	if e.Bind(&restricted) != nil {
		return BadRequest(e)
	}

	// Parse Id
	if ParseId(e, &restricted) != nil {
		return BadRequest(e)
	}

	// Get By Id
	RestrictedDto := h.usecase.GetById(restricted.Id)

	// Check Value
	if RestrictedDto.Id == "" {
		return NotFound(e, message.NotFound)
	}

	// Return Success
	return Success(e, message.GetSuccess, RestrictedDto)
}

// @Summary Create a new Restricted
// @Description Add a new Restricted to the Restricted machine
// @Success 200 {object} dto.Response
func (h *RestrictedHandler) Create(e echo.Context) error {

	var restrictedDto dto.Restricted
	SetUsername(&restrictedDto, e)

	if e.Bind(&restrictedDto) != nil {
		return BadRequest(e)
	}

	if Validate(&restrictedDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Create(restrictedDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.CreateSuccess, "")
}

// @Summary Update an existing Restricted
// @Description Update an Restricted in the Restricted machine by its ID
// @Success 200 {object} dto.Response
func (h *RestrictedHandler) Update(e echo.Context) error {

	var restrictedDto dto.Restricted
	SetUsername(&restrictedDto, e)

	if e.Bind(&restrictedDto) != nil {
		return BadRequest(e)
	}

	if Validate(&restrictedDto) != nil {
		return BadRequest(e)
	}

	if ParseId(e, &restrictedDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Update(restrictedDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.UpdateSuccess, "")
}

// @Summary Delete an dto.Restricted
// @Description Remove an dto.Restricted from the Restricted machine by its ID
// @Success 200 {object} dto.Response
func (h *RestrictedHandler) Delete(e echo.Context) error {

	var restrictedDto dto.Restricted
	SetUsername(&restrictedDto, e)

	if e.Bind(&restrictedDto) != nil {
		return BadRequest(e)
	}

	if ParseId(e, &restrictedDto) != nil {
		return BadRequest(e)
	}

	if err := h.usecase.Delete(restrictedDto); err != nil {
		return Error(e, err.Error())
	}

	return Success(e, message.DeleteSuccess, "")
}
