package routes

import (
	"robot-fleet-monitoring/service-robot/handler"

	"github.com/labstack/echo/v4"
)

func RobotRoute(routes *echo.Echo, handler *handler.RobotHandler) {

	robot := routes.Group("/robots")
	{
		robot.GET("", handler.GetAll)
		robot.GET("/:id", handler.GetById)

		robot.POST("", handler.Create)
		robot.PUT("/:id", handler.Update)
		robot.DELETE("/:id", handler.Delete)
	}
}
