// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"log"
	"net/http"
	config "robot-fleet-monitoring/service-monitor/config"

	routes2 "robot-fleet-monitoring/service-monitor/routes"
)

func main() {

	config.ReadEnv()

	rabbitConfig := config.InitRabbit()
	routes2.Init(rabbitConfig)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
