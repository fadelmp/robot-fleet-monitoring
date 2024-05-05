package handler

import (
	"robot-fleet-monitoring/service-restricted/dto"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Validate(dto *dto.Restricted) error {

	validate := validator.New()

	if err := validate.Struct(*dto); err != nil {
		return err
	}

	return nil
}

func ParseId(e echo.Context, dto *dto.Restricted) error {

	dto.Id = e.Param("id")

	return nil
}
