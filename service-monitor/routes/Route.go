package routes

import (
	"robot-fleet-monitoring/service-monitor/handler"

	"github.com/streadway/amqp"
)

func Init(channel *amqp.Channel) {

	// Swagger Documentation Route
	//SwaggerRoute(routes)

	// Monitor Route & Injection
	monitor := handler.NewMonitorHandler()
	RabbitMqRoute(channel, monitor)

	WebSocket()
	go handler.StartWebSocketServer()
}
