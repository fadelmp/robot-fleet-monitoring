package routes

import (
	"robot-fleet-monitoring/service-restricted/handler"

	"github.com/labstack/echo/v4"
)

func RestrictedRoute(routes *echo.Echo, handler *handler.RestrictedHandler) {

	restricted := routes.Group("/restricteds")
	{
		restricted.GET("", handler.GetAll)
		restricted.GET("/:id", handler.GetById)

		restricted.POST("", handler.Create)
		restricted.PUT("/:id", handler.Update)
		restricted.DELETE("/:id", handler.Delete)
	}
}
