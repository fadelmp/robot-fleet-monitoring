package mapper_test

import (
	"robot-fleet-monitoring/service-robot/domain"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/mapper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the RobotMapper
type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestRobotMapper_ToRobot(t *testing.T) {

	// Setup
	RobotMap := mapper.NewRobotMapper()

	// RobotDto
	RobotDto := dto.Robot{
		Id:        1,
		Name:      "Lychee",
		Longitude: -1.0000,
		Latitude:  1.0000,
	}

	// Test
	Robot := RobotMap.ToRobot(RobotDto)

	// Assert
	assert.NotNil(t, Robot)
	assert.NotEmpty(t, Robot)
}

func (suite *TestSuite) TestRobotMapper_ToRobotDto(t *testing.T) {

	// Setup
	RobotMap := mapper.NewRobotMapper()

	// RobotDto
	Robot := domain.Robot{
		Id:        1,
		Name:      "Lychee",
		Longitude: -1.0000,
		Latitude:  1.0000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
			DeletedAt: time.Now(),
			DeletedBy: "System",
		},
	}

	// Test
	RobotDto := RobotMap.ToRobotDto(Robot)

	// Assert
	assert.NotNil(t, RobotDto)
	assert.NotEmpty(t, RobotDto)
}

func (suite *TestSuite) TestRobotMapper_ToRobotDtoList(t *testing.T) {

	// Setup
	RobotMap := mapper.NewRobotMapper()

	var Robots []domain.Robot

	// RobotDto
	Robot := domain.Robot{
		Id:        1,
		Name:      "Lychee",
		Longitude: -1.0000,
		Latitude:  1.0000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
			DeletedAt: time.Now(),
			DeletedBy: "System",
		},
	}

	Robots = append(Robots, Robot)
	Robots = append(Robots, Robot)
	Robots = append(Robots, Robot)
	Robots = append(Robots, Robot)
	Robots = append(Robots, Robot)

	// Test
	RobotDtos := RobotMap.ToRobotDtoList(Robots)

	// Assert
	assert.NotNil(t, RobotDtos)
	assert.NotEmpty(t, RobotDtos)
}
