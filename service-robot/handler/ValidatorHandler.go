package handler

import (
	"robot-fleet-monitoring/service-robot/dto"

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

	dto.Id = e.Param("id")

	return nil
}
