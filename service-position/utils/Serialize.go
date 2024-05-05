package utils

import (
	"encoding/json"
	"log"
	"robot-fleet-monitoring/service-position/dto"
)

// GenerateUUID generates a new UUID (Universally Unique Identifier)
func SerializePosition(dto dto.Position) []byte {

	// Serialize DTO object to JSON
	body, err := json.Marshal(dto)

	if err != nil {
		log.Fatalf("Failed to serialize DTO to JSON: %v", err)
	}

	return body
}

func SerializeRobot(dto dto.Robot) []byte {

	// Serialize DTO object to JSON
	body, err := json.Marshal(dto)

	if err != nil {
		log.Fatalf("Failed to serialize DTO to JSON: %v", err)
	}

	return body
}
