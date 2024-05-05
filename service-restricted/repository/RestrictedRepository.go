package repository

import (
	"robot-fleet-monitoring/service-restricted/config"
	"robot-fleet-monitoring/service-restricted/domain"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type RestrictedRepositoryContract interface {
	GetAll() []domain.Restricted
	GetById(string) domain.Restricted
	GetByName(string) domain.Restricted

	Create(*domain.Restricted) error
	Update(*domain.Restricted) error
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
	query := r.DB.Model(&restricteds).Unscoped().Where("is_deleted=?", false).Find(&restricteds)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "restricteds")

	return restricteds
}

func (r *RestrictedRepository) GetById(id string) domain.Restricted {

	var restricted domain.Restricted

	// Get Data By Id
	query := r.DB.Model(&restricted).Unscoped().Where("is_deleted=?", false).Where("id=?", id).Find(&restricted)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "restricted_id_"+id)

	return restricted
}

func (r *RestrictedRepository) GetByName(name string) domain.Restricted {

	var restricted domain.Restricted

	// Get Data By Name
	query := r.DB.Model(&restricted).Unscoped().Where("is_deleted=?", false).Where("name=?", name).Find(&restricted)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "restricted_name_"+name)

	return restricted
}

func (r *RestrictedRepository) Create(restricted *domain.Restricted) error {

	// Flush Restricted Cache
	config.FlushData(r.Redis, "restricted*")

	// Create Restricted
	return r.DB.Create(&restricted).Error

}

func (r *RestrictedRepository) Update(restricted *domain.Restricted) error {

	// Flush Restricted Cache
	config.FlushData(r.Redis, "restricted*")

	// Update Restricted
	return r.DB.Model(&restricted).Unscoped().Update(&restricted).Error
}
