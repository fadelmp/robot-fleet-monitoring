package mapper

import (
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/utils"
)

// Interface
type AreaMapperContract interface {
	ToAreaList([]dto.Area, domain.Restricted) []domain.Area
	ToAreaDtoList([]domain.Area) []dto.Area
}

// Class
type AreaMapper struct {
}

// Constructor
func NewAreaMapper() *AreaMapper {
	return &AreaMapper{}
}

// Implementation

func (m *AreaMapper) ToAreaList(areaDtos []dto.Area, restricted domain.Restricted) []domain.Area {

	areas := make([]domain.Area, len(areaDtos))

	for i, value := range areaDtos {
		areas[i] = m.toArea(value, restricted)
	}

	return areas
}

func (m *AreaMapper) ToAreaDtoList(areas []domain.Area) []dto.Area {

	areaDtos := make([]dto.Area, len(areas))

	for i, value := range areas {
		areaDtos[i] = m.toAreaDto(value)
	}

	return areaDtos
}

func (m *AreaMapper) toArea(areaDto dto.Area, restricted domain.Restricted) domain.Area {

	id, _ := utils.GenerateUUID()

	return domain.Area{
		Id:         id,
		Restricted: restricted,
		Longitude:  areaDto.Longitude,
		Latitude:   areaDto.Latitude,
	}
}

func (m *AreaMapper) toAreaDto(Area domain.Area) dto.Area {

	return dto.Area{
		Longitude: Area.Longitude,
		Latitude:  Area.Latitude,
	}
}
