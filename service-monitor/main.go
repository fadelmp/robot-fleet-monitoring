// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"net/http"
	config "robot-fleet-monitoring/service-robot/config"
)

func main() {

	config.ReadEnv()

	rabbitConfig := config.InitRabbit()

	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
