package routes

import (
	"net/http"
	"robot-fleet-monitoring/service-monitor/handler"
)

func WebSocket() {

	http.HandleFunc("/ws", handler.Broadcast)
}
