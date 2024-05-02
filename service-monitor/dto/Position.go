// dto/Robot.go

package dto

// Payload Design

// @Summary RobotDto object
// @Description Represents a RobotDto Machine VendingDto
type Position struct {
	RobotId   string  `json:"robot_id"`
	RobotName string  `json:"robot_name"`
	Timestamp string  `json:"timestamp"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Message   string  `json:"message"`
}
