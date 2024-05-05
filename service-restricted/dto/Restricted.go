// dto/Robot.go

package dto

// Payload Design

// @Summary RobotDto object
// @Description Represents a RobotDto Machine VendingDto
type Restricted struct {
	Id          string `json:"id"`
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Description string `json:"description"`
	Area        []Area `json:"area"`
	Base
}
