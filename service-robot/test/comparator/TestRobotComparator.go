package comparator_test

import (
	comparator "robot-fleet-monitoring/service-robot/comparator"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the RobotComparator
type TestSuite struct {
	suite.Suite
	repository repository.RobotRepository
}

func (suite *TestSuite) TestRobotComparator_CheckId(t *testing.T) {

	// Setup
	RobotComparator := comparator.NewRobotComparator(&suite.repository)

	// Test
	err := RobotComparator.CheckId("aaaa-aaaa-aaaa-aaaa-aaaa-aaaa")

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestRobotComparator_CheckName(t *testing.T) {

	// Setup
	RobotComparator := comparator.NewRobotComparator(&suite.repository)

	// RobotDto
	RobotDto := dto.Robot{
		Id:   "aaaa-aaaa-aaaa-aaaa-aaaa-aaaa",
		Name: "Milo",
	}

	// Test
	err := RobotComparator.CheckName(RobotDto)

	// Assert
	assert.NoError(t, err)
}
