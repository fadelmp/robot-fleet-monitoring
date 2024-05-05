// dto/Robot.go

package dto

// Payload Design

// @Summary RobotDto object
// @Description Represents a RobotDto Machine VendingDto
type Robot struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
