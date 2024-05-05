package mapper

import (
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
)

// Interface
type RestrictedMapperContract interface {
	ToRestricted(dto.Restricted) domain.Restricted
	ToRestrictedDto(domain.Restricted) dto.Restricted
	ToRestrictedDtoList([]domain.Restricted) []dto.Restricted
}

// Class
type RestrictedMapper struct {
}

// Constructor
func NewRestrictedMapper() *RestrictedMapper {
	return &RestrictedMapper{}
}

// Implementation

func (m *RestrictedMapper) ToRestricted(restrictedDto dto.Restricted) domain.Restricted {

	return domain.Restricted{
		Id:          restrictedDto.Id,
		Name:        restrictedDto.Name,
		Description: restrictedDto.Description,
	}
}

func (m *RestrictedMapper) ToRestrictedDto(restricted domain.Restricted) dto.Restricted {

	return dto.Restricted{
		Id:          restricted.Id,
		Name:        restricted.Name,
		Description: restricted.Description,
		Area:        NewAreaMapper().ToAreaDtoList(restricted.Area),
		Base:        NewBaseMapper().ToBaseDto(restricted.Base),
	}
}

func (m *RestrictedMapper) ToRestrictedDtoList(restricteds []domain.Restricted) []dto.Restricted {

	restrictedDtos := make([]dto.Restricted, len(restricteds))

	for i, value := range restricteds {
		restrictedDtos[i] = m.ToRestrictedDto(value)
	}

	return restrictedDtos
}
