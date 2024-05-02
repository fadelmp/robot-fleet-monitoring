package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketHandlerContract interface {
	Broadcast(http.ResponseWriter, *http.Request)
}

// Class
type WebSocketHandler struct {
}

// Constructor
func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{}
}

func Broadcast(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}

	defer conn.Close()

	for {
		// Read message from client
		msg, _ := <-messageQueue

		// Echo message back to client
		err = conn.WriteMessage(websocket.TextMessage, msg.Body)
	}
}

var messageQueue = make(chan amqp.Delivery)
