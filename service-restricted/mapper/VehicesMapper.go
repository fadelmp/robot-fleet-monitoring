package mapper

import (
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/utils"
)

// Interface
type VehicesMapperContract interface {
	ToVehicesList([]dto.Vehices) domain.Vehices
	ToVehicesDtoList([]domain.Vehices) []dto.Vehices
}

// Class
type VehicesMapper struct {
}

// Constructor
func NewVehicesMapper() *VehicesMapper {
	return &VehicesMapper{}
}

// Implementation

func (m *VehicesMapper) ToVehicesList(vehiceDtos []dto.Vehices) []domain.Vehices {

	vehicess := make([]domain.Vehices, len(vehiceDtos))

	for i, value := range vehiceDtos {
		vehicess[i] = m.toVehices(value)
	}

	return vehicess
}

func (m *VehicesMapper) ToVehicesDtoList(vehicess []domain.Vehices) []dto.Vehices {

	vehicesDtos := make([]dto.Vehices, len(vehicess))

	for i, value := range vehicess {
		vehicesDtos[i] = m.toVehicesDto(value)
	}

	return vehicesDtos
}

func (m *VehicesMapper) toVehices(vehicesDto dto.Vehices) domain.Vehices {

	id, _ := utils.GenerateUUID()

	return domain.Vehices{
		Id:        id,
		Longitude: vehicesDto.Longitude,
		Latitude:  vehicesDto.Latitude,
	}
}

func (m *VehicesMapper) toVehicesDto(vehices domain.Vehices) dto.Vehices {

	return dto.Vehices{
		Longitude: vehices.Longitude,
		Latitude:  vehices.Latitude,
	}
}
