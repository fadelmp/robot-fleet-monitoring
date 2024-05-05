package routes

import (
	"robot-fleet-monitoring/service-robot/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client, channel *amqp.Channel) *echo.Echo {

	// Swagger Documentation Route
	//SwaggerRoute(routes)

	// Robot Route & Injection
	robot := injection.RobotInjection(db, redis)
	RobotRoute(routes, robot)

	// Rabbit Route
	rabbit := injection.RabbitInjection(db, redis)
	RabbitMqRoute(channel, rabbit)

	return routes
}
