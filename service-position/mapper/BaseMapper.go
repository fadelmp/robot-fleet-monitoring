package mapper

import (
	"robot-fleet-monitoring/service-position/domain"
	"time"
)

// Interface
type BaseMapperContract interface {
	Create(*domain.Base, string)
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
	base.CreatedAt = time.Now()
	base.CreatedBy = name
	base.UpdatedAt = time.Now()
	base.UpdatedBy = name
}
