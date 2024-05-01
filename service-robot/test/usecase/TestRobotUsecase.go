package usecase_test

import (
	"robot-fleet-monitoring/service-robot/comparator"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/mapper"
	"robot-fleet-monitoring/service-robot/repository"
	usecase "robot-fleet-monitoring/service-robot/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the RobotUsecase
type TestSuite struct {
	suite.Suite
	baseMapper *mapper.BaseMapper
	mapper     *mapper.RobotMapper
	repo       *repository.RobotRepository
	comparator *comparator.RobotComparator
}

func (suite *TestSuite) TestRobotUsecase_GetAll(t *testing.T) {

	// Setup
	RobotUsecase := usecase.NewRobotUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// Test
	Robots := RobotUsecase.GetAll()

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotUsecase_GetById(t *testing.T) {

	// Setup
	RobotUsecase := usecase.NewRobotUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// Test
	Robots := RobotUsecase.GetById("aaaa-aaaa-aaaa-aaaa-aaaa-aaaa")

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotUsecase_Create(t *testing.T) {

	// Setup
	RobotUsecase := usecase.NewRobotUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// RobotDto
	RobotDto := dto.Robot{
		Name:      "Lychee",
		Longitude: 1.0000,
		Latitude:  1.0000,
		Base: dto.Base{
			CreatedBy: "System",
		},
	}

	// Test
	err := RobotUsecase.Create(RobotDto)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestRobotUsecase_Update(t *testing.T) {

	// Setup
	RobotUsecase := usecase.NewRobotUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// RobotDto
	RobotDto := dto.Robot{
		Id:        "aaaa-aaaa-aaaa-aaaa-aaaa-aaaa",
		Name:      "Lychee",
		Longitude: 1.0000,
		Latitude:  1.0000,
		Base: dto.Base{
			UpdatedBy: "System",
		},
	}

	// Test
	err := RobotUsecase.Update(RobotDto)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestRobotUsecase_Delete(t *testing.T) {

	// Setup
	RobotUsecase := usecase.NewRobotUsecase(suite.baseMapper, suite.mapper, suite.repo, suite.comparator)

	// RobotDto
	RobotDto := dto.Robot{
		Id: "aaaa-aaaa-aaaa-aaaa-aaaa-aaaa",
		Base: dto.Base{
			DeletedBy: "System",
		},
	}

	// Test
	err := RobotUsecase.Delete(RobotDto)

	// Assert
	assert.NoError(t, err)
}
