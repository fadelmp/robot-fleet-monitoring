package repository_test

import (
	"robot-fleet-monitoring/service-robot/domain"
	"robot-fleet-monitoring/service-robot/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite represents the test suite for the RobotRepository
type TestSuite struct {
	suite.Suite
	repo *repository.RobotRepository
}

func (suite *TestSuite) TestRobotRepository_GetAll(t *testing.T) {

	// Setup
	repo := repository.NewRobotRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	Robots := repo.GetAll()

	// Assert
	assert.NotNil(t, Robots)
	assert.NotEmpty(t, Robots)
}

func (suite *TestSuite) TestRobotRepository_GetById(t *testing.T) {

	// Setup
	repo := repository.NewRobotRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	Robot := repo.GetById("aaaa-aaaa-aaaa-aaaa-aaaa-aaaa")

	// Assert
	assert.NotNil(t, Robot)
	assert.NotEmpty(t, Robot)
}

func (suite *TestSuite) TestRobotRepository_GetByName(t *testing.T) {

	// Setup
	repo := repository.NewRobotRepository(suite.repo.DB, suite.repo.Redis)

	// Test
	Robot := repo.GetByName("Milo")

	// Assert
	assert.NotNil(t, Robot)
	assert.NotEmpty(t, Robot)
}

func (suite *TestSuite) TestRobotRepository_Create(t *testing.T) {

	// Setup
	repo := repository.NewRobotRepository(suite.repo.DB, suite.repo.Redis)

	// Robot
	Robot := &domain.Robot{
		Name:      "Lychee",
		Longitude: 1.0000,
		Latitude:  1.0000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			CreatedAt: time.Now(),
			CreatedBy: "System",
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
		},
	}

	// Test
	err := repo.Create(Robot)

	// Assert
	assert.NoError(t, err)
}

func (suite *TestSuite) TestRobotRepository_Update(t *testing.T) {

	// Setup
	repo := repository.NewRobotRepository(suite.repo.DB, suite.repo.Redis)

	// Robot
	Robot := &domain.Robot{
		Id:        "aaaa-aaaa-aaaa-aaaa-aaaa-aaaa",
		Name:      "Lychee",
		Longitude: 1.0000,
		Latitude:  1.0000,
		Base: domain.Base{
			IsActived: true,
			IsDeleted: false,
			UpdatedAt: time.Now(),
			UpdatedBy: "System",
		},
	}

	// Test
	err := repo.Update(Robot)

	// Assert
	assert.NoError(t, err)
}
