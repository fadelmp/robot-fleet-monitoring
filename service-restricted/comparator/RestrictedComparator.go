package comparator

import (
	"errors"
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/message"
	"robot-fleet-monitoring/service-restricted/repository"
)

// Interface
type RestrictedComparatorContract interface {
	CheckId(string) error
	CheckName(dto.Restricted) error
}

// Class
type RestrictedComparator struct {
	repo repository.RestrictedRepositoryContract
}

// Constructor
func NewRestrictedComparator(repo repository.RestrictedRepositoryContract) *RestrictedComparator {
	return &RestrictedComparator{
		repo: repo,
	}
}

// Implementation

func (c *RestrictedComparator) CheckId(id string) error {

	// Get Data By Id
	restricted := c.repo.GetById(id)

	// Return Error If Data Not Found
	if restricted.Id == "" {

		return errors.New(message.NotFound)
	}

	return nil
}

func (c *RestrictedComparator) CheckName(dto dto.Restricted) error {

	// Get Data By Name
	restricted := c.repo.GetByName(dto.Name)

	// Return Error If Data Exists
	if restricted.Id != "" && restricted.Id != dto.Id {

		return errors.New(message.NameExists)
	}

	return nil
}
