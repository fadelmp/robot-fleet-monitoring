package comparator

import (
	"robot-fleet-monitoring/service-position/dto"
	"robot-fleet-monitoring/service-position/message"
	"robot-fleet-monitoring/service-position/repository"
	"strings"
)

// Interface
type PositionComparatorContract interface {
	CheckArea(dto.Position) (string, string, string)
}

// Class
type PositionComparator struct {
	repo repository.RestrictedRepositoryContract
}

// Constructor
func NewPositionComparator(repo repository.RestrictedRepositoryContract) *PositionComparator {
	return &PositionComparator{
		repo: repo,
	}
}

// Implementation

func (c *PositionComparator) CheckArea(positionDto dto.Position) (string, string, string) {

	alert := ""
	restrictedId, restrictedName := c.isInside(positionDto.Longitude, positionDto.Latitude)

	if restrictedId == "" {
		alert = message.Alert
		alert = strings.Replace(alert, "robot_name", positionDto.RobotId, 1)
		alert = strings.Replace(alert, "restricted_name", restrictedName, 1)
	}

	return alert, restrictedId, restrictedName
}

func (c *PositionComparator) isInside(longitude float64, latitude float64) (string, string) {

	// Get Data By Id
	restricteds := c.repo.GetAll()

	for _, value := range restricteds {

		n := len(value.Area)
		intersections := 0

		for i := 0; i < n; i++ {
			p1 := value.Area[i]
			p2 := value.Area[(i+1)%n]

			if latitude >= p1.Latitude && latitude < p2.Latitude ||
				latitude >= p2.Latitude && latitude < p1.Latitude {

				if (latitude-p1.Latitude)*(p2.Longitude-p1.Longitude)/
					(p2.Latitude-p1.Latitude)+p1.Longitude < longitude {
					intersections++
				}
			}
		}

		if intersections%2 != 0 {
			return value.Id, value.Name
		}
	}

	return "", ""
}
