package handler

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

// Interface
type MonitorHandlerContract interface {
	StartWebSocketServer()
	Consume(amqp.Delivery) error
	WebSocket(http.ResponseWriter, *http.Request)
}

// Class
type MonitorHandler struct {
}

// Constructor
func NewMonitorHandler() *MonitorHandler {
	return &MonitorHandler{}
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Define a channel for broadcasting messages to WebSocket clients
	broadcast = make(chan []byte)

	// Define a mutex to protect concurrent access to the clients map
	clientsMu sync.Mutex
	clients   = make(map[*websocket.Conn]bool) // Map to store connected WebSocket clients
)

// WebSocket handler function
func WebSocket(w http.ResponseWriter, r *http.Request) {

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}

	defer conn.Close()

	// Register the client's WebSocket connection
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	// Handle WebSocket events
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()
			break
		}
	}
}

// Start the WebSocket server
func StartWebSocketServer() {

	// Broadcast messages to connected WebSocket clients
	for {

		msg := <-broadcast
		for client := range clients {

			err := client.WriteMessage(websocket.TextMessage, msg)

			if err != nil {
				log.Println("Error broadcasting message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (h *MonitorHandler) Consume(msg amqp.Delivery) {

	go func() {
		h.broadcastMsg(msg.Body)
	}()
}

func (h *MonitorHandler) broadcastMsg(msg []byte) {

	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {

		err := client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
