package comparator

import (
	"errors"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/message"
	"robot-fleet-monitoring/service-robot/repository"
)

// Interface
type RobotComparatorContract interface {
	CheckId(uint) error
	CheckName(dto.Robot) error
}

// Class
type RobotComparator struct {
	repo repository.RobotRepositoryContract
}

// Constructor
func NewRobotComparator(repo repository.RobotRepositoryContract) *RobotComparator {
	return &RobotComparator{
		repo: repo,
	}
}

// Implementation

func (c *RobotComparator) CheckId(id uint) error {

	// Get Data By Id
	robot := c.repo.GetById(id)

	// Return Error If Data Not Found
	if robot.Id == 0 {

		return errors.New(message.NotFound)
	}

	return nil
}

func (c *RobotComparator) CheckName(dto dto.Robot) error {

	// Get Data By Name
	robot := c.repo.GetByName(dto.Name)

	// Return Error If Data Exists
	if robot.Id != 0 && robot.Id != dto.Id {

		return errors.New(message.NameExists)
	}

	return nil
}
