package routes

import (
	"robot-fleet-monitoring/service-position/handler"

	"github.com/labstack/echo/v4"
)

func PositionRoute(routes *echo.Echo, handler *handler.PositionHandler) {

	position := routes.Group("/positions")
	{
		position.POST("", handler.Update)
	}
}
