package repository

import (
	"robot-fleet-monitoring/service-robot/config"
	"robot-fleet-monitoring/service-robot/domain"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type RobotRepositoryContract interface {
	GetAll() []domain.Robot
	GetById(string) domain.Robot
	GetByName(string) domain.Robot

	Create(*domain.Robot) error
	Update(*domain.Robot) error
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

	var robots []domain.Robot

	// Get All Data
	query := r.DB.Model(&robots).
		Unscoped().
		Where("is_deleted=?", false).
		Find(&robots)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "robots")

	return robots
}

func (r *RobotRepository) GetById(id string) domain.Robot {

	var robot domain.Robot

	// Get Data By Id
	query := r.DB.Model(&robot).
		Unscoped().
		Where("is_deleted=?", false).
		Where("id=?", id).
		Find(&robot)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "robot_id_"+id)

	return robot
}

func (r *RobotRepository) GetByName(name string) domain.Robot {

	var robot domain.Robot

	// Get Data By Name
	query := r.DB.Model(&robot).
		Unscoped().
		Where("is_deleted=?", false).
		Where("name=?", name).
		Find(&robot)

	// Get Data From Redis
	config.QueryData(r.Redis, query, "robot_name_"+name)

	return robot
}

func (r *RobotRepository) Create(robot *domain.Robot) error {

	// Flush Robot Cache
	config.FlushData(r.Redis, "robot*")

	// Create Robot
	return r.DB.Create(&robot).Error

}

func (r *RobotRepository) Update(robot *domain.Robot) error {

	// Flush Robot Cache
	config.FlushData(r.Redis, "robot*")

	// Update Robot
	return r.DB.Model(&robot).Unscoped().Update(&robot).Error
}
