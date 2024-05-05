package mapper

import (
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
	"time"
)

// Interface
type BaseMapperContract interface {
	Create(*domain.Base, string)
	Update(*domain.Base, string)
	Delete(*domain.Restricted, string, string)

	ToBaseDto(base domain.Base) dto.Base
}

// Class
type BaseMapper struct {
}

// Constructor
func NewBaseMapper() *BaseMapper {
	return &BaseMapper{}
}

// Implementation

func (m *BaseMapper) Create(base *domain.Base, name string) {

	// Set Created Value
	base.IsActived = true
	base.IsDeleted = false
	base.CreatedBy = name
	base.UpdatedBy = name
}

func (m *BaseMapper) Update(Base *domain.Base, name string) {

	// Set Updated Value
	Base.IsActived = true
	Base.IsDeleted = false
	Base.UpdatedBy = name
}

func (m *BaseMapper) Delete(domain *domain.Restricted, id string, name string) {

	// Set Deleted Value
	domain.Id = id
	domain.Base.IsActived = false
	domain.Base.IsDeleted = true
	domain.Base.UpdatedBy = name
	domain.Base.DeletedBy = name
	domain.Base.DeletedAt = time.Now()
}

func (m *BaseMapper) ToBaseDto(base domain.Base) dto.Base {

	// Map Base to Base Dto
	return dto.Base{
		IsActived: base.IsActived,
		IsDeleted: base.IsDeleted,
		CreatedAt: base.CreatedAt,
		CreatedBy: base.CreatedBy,
		UpdatedAt: base.UpdatedAt,
		UpdatedBy: base.UpdatedBy,
		DeletedAt: base.DeletedAt,
		DeletedBy: base.DeletedBy,
	}
}
