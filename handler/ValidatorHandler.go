package handler

import (
	"robot-fleet-monitoring/dto"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Validate(dto *dto.Robot) error {

	validate := validator.New()

	if err := validate.Struct(*dto); err != nil {
		return err
	}

	return nil
}

func ParseId(e echo.Context, dto *dto.Robot) error {

	idStr := e.Param("id")

	// Parse the id string to uint
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return err
	}

	dto.Id = uint(id)

	return nil
}
