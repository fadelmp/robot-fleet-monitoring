package handler

import (
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/message"

	"github.com/labstack/echo/v4"
)

func CreateResponse(c echo.Context, httpCode int, response dto.Response) error {

	c.Response().WriteHeader(httpCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSONPretty(httpCode, response, "  ")
}

func Success(c echo.Context, messages string, data interface{}) error {

	httpCode := 200
	resp := &dto.Response{Messages: messages, Data: data}

	return CreateResponse(c, httpCode, *resp)
}

func Error(c echo.Context, messages interface{}) error {

	httpCode := 500
	resp := &dto.Response{Messages: messages}

	return CreateResponse(c, httpCode, *resp)
}

func NotFound(c echo.Context, messages interface{}) error {

	httpCode := 404
	resp := &dto.Response{Messages: messages}

	return CreateResponse(c, httpCode, *resp)
}

func BadRequest(c echo.Context) error {

	httpCode := 400
	resp := &dto.Response{Messages: message.BadRequest}

	return CreateResponse(c, httpCode, *resp)
}

func SetUsername(restrictedDto *dto.Restricted, c echo.Context) {

	username := c.Request().Header.Get("username")

	restrictedDto.Base.CreatedBy = username
	restrictedDto.Base.UpdatedBy = username
	restrictedDto.Base.DeletedBy = username
}
