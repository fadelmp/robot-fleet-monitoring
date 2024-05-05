package repository

import (
	"robot-fleet-monitoring/service-restricted/domain"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type AreaRepositoryContract interface {
	Create([]domain.Area) error
	Delete(domain.Area) error
}

// Class
type AreaRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewAreaRepository(DB *gorm.DB, Redis *redis.Client) *AreaRepository {
	return &AreaRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *AreaRepository) Create(areas []domain.Area) error {

	for _, area := range areas {
		if err := r.DB.Create(&area).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *AreaRepository) Delete(area domain.Area) error {

	// Delete Area
	return r.DB.Delete(&area).Error
}
