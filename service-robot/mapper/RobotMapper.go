package mapper

import (
	"robot-fleet-monitoring/service-robot/domain"
	"robot-fleet-monitoring/service-robot/dto"
)

// Interface
type RobotMapperContract interface {
	ToRobot(dto.Robot) domain.Robot
	ToRobotDto(domain.Robot) dto.Robot
	ToRobotDtoList([]domain.Robot) []dto.Robot
}

// Class
type RobotMapper struct {
}

// Constructor
func NewRobotMapper() *RobotMapper {
	return &RobotMapper{}
}

// Implementation

func (m *RobotMapper) ToRobot(robotDto dto.Robot) domain.Robot {

	return domain.Robot{
		Id:        robotDto.Id,
		Name:      robotDto.Name,
		Longitude: robotDto.Longitude,
		Latitude:  robotDto.Latitude,
	}
}

func (m *RobotMapper) ToRobotDto(robot domain.Robot) dto.Robot {

	return dto.Robot{
		Id:        robot.Id,
		Name:      robot.Name,
		Longitude: robot.Longitude,
		Latitude:  robot.Latitude,
		Base:      NewBaseMapper().ToBaseDto(robot.Base),
	}
}

func (m *RobotMapper) ToRobotDtoList(robots []domain.Robot) []dto.Robot {

	robotDtos := make([]dto.Robot, len(robots))

	for i, value := range robots {
		robotDtos[i] = m.ToRobotDto(value)
	}

	return robotDtos
}
