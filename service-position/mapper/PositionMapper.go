package mapper

import (
	"robot-fleet-monitoring/service-position/domain"
	"robot-fleet-monitoring/service-position/dto"
)

// Interface
type PositionMapperContract interface {
	ToRobotDto(dto.Position) dto.Robot
	ToPosition(dto.Position, string, string) domain.Position
}

// Class
type PositionMapper struct {
}

// Constructor
func NewPositionMapper() *PositionMapper {
	return &PositionMapper{}
}

// Implementation

func (m *PositionMapper) ToRobotDto(positionDto dto.Position) dto.Robot {

	return dto.Robot{
		Id:        positionDto.RobotId,
		Name:      positionDto.RobotName,
		Longitude: positionDto.Longitude,
		Latitude:  positionDto.Latitude,
	}
}

func (m *PositionMapper) ToPosition(
	positionDto dto.Position,
	restrictedId string,
	restrictedName string,
) domain.Position {

	return domain.Position{
		Id:             positionDto.Id,
		RobotId:        positionDto.RobotId,
		RobotName:      positionDto.RobotName,
		RestrictedId:   restrictedId,
		RestrictedName: restrictedName,
	}
}
