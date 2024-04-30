package routes

import (
	"robot-fleet-monitoring/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// Swagger Documentation Route
	//SwaggerRoute(routes)

	// Robot Route & Injection
	robot := injection.RobotInjection(db, redis)
	RobotRoute(routes, robot)

	return routes
}
