package repository

import (
	"robot-fleet-monitoring/service-robot/config"
	"robot-fleet-monitoring/service-robot/domain"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type RobotRepositoryContract interface {
	GetAll() []domain.Robot
	GetById(uint) domain.Robot
	GetByName(string) domain.Robot

	Create(*domain.Robot) error
	Update(*domain.Robot) error
	Delete(*domain.Robot) error
}

// Class
type RobotRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewRobotRepository(DB *gorm.DB, Redis *redis.Client) *RobotRepository {
	return &RobotRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *RobotRepository) GetAll() []domain.Robot {

	var Robots []domain.Robot

	// Get All Data
	query := r.DB.Model(&Robots).
		Unscoped().
		Where("is_deleted=?", false).
		Find(&Robots)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "Robots")

	return Robots
}

func (r *RobotRepository) GetById(id uint) domain.Robot {

	var Robot domain.Robot

	// Get Data By Id
	query := r.DB.Model(&Robot).
		Unscoped().
		Where("is_deleted=?", false).
		Where("id=?", id).
		Find(&Robot)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "Robot_id_"+strconv.FormatUint(uint64(id), 10))

	return Robot
}

func (r *RobotRepository) GetByName(name string) domain.Robot {

	var Robot domain.Robot

	// Get Data By Name
	query := r.DB.Model(&Robot).
		Unscoped().
		Where("is_deleted=?", false).
		Where("name=?", name).
		Find(&Robot)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "Robot_name_"+name)

	return Robot
}

func (r *RobotRepository) Create(Robot *domain.Robot) error {

	// Flush Robot Cache
	config.FlushData(r.Redis, "Robot*")

	// Create Robot
	return r.DB.Create(&Robot).Error

}

func (r *RobotRepository) Update(Robot *domain.Robot) error {

	// Flush Robot Cache
	config.FlushData(r.Redis, "Robot*")

	// Update Robot
	return r.DB.Model(&Robot).Unscoped().Update(&Robot).Error
}

func (r *RobotRepository) Delete(Robot *domain.Robot) error {

	// Flush Robot Cache
	config.FlushData(r.Redis, "Robot*")

	// Delete Robot
	return r.DB.Model(&Robot).Unscoped().Update(&Robot).Error
}
