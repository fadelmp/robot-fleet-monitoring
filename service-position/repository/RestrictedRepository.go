package repository

import (
	"robot-fleet-monitoring/service-position/config"
	"robot-fleet-monitoring/service-position/domain"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type RestrictedRepositoryContract interface {
	GetAll() []domain.Restricted
}

// Class
type RestrictedRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewRestrictedRepository(DB *gorm.DB, Redis *redis.Client) *RestrictedRepository {
	return &RestrictedRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *RestrictedRepository) GetAll() []domain.Restricted {

	var restricteds []domain.Restricted

	// Get All Data
	query := r.DB.Model(&restricteds).Unscoped().Where("is_deleted=?", false).Preload("area").Find(&restricteds)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "restricteds")

	return restricteds
}
