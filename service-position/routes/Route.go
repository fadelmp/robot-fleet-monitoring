package routes

import (
	"robot-fleet-monitoring/service-position/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client, channel *amqp.Channel) *echo.Echo {

	// Swagger Documentation Route
	//SwaggerRoute(routes)

	// Robot Route & Injection
	position := injection.PositionInjection(db, redis, channel)
	PositionRoute(routes, position)

	return routes
}
