package repository

import (
	"robot-fleet-monitoring/service-position/domain"

	"github.com/jinzhu/gorm"
)

// Interface
type PositionRepositoryContract interface {
	Create(*domain.Position) error
}

// Class
type PositionRepository struct {
	DB *gorm.DB
}

// Constructor
func NewPositionRepository(DB *gorm.DB) *PositionRepository {
	return &PositionRepository{
		DB: DB,
	}
}

// Implementation

func (r *PositionRepository) Create(position *domain.Position) error {

	// Create Position
	return r.DB.Create(&position).Error
}
