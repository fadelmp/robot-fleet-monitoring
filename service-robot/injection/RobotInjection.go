package injection

import (
	"robot-fleet-monitoring/service-robot/comparator"
	"robot-fleet-monitoring/service-robot/handler"
	"robot-fleet-monitoring/service-robot/mapper"
	"robot-fleet-monitoring/service-robot/repository"
	"robot-fleet-monitoring/service-robot/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func RobotInjection(db *gorm.DB, redis *redis.Client) *handler.RobotHandler {

	usecase := robotUsecase(db, redis)
	handler := handler.NewRobotHandler(usecase)

	return handler
}

func RabbitInjection(db *gorm.DB, redis *redis.Client) *handler.RabbitMqHandler {

	usecase := robotUsecase(db, redis)
	handler := handler.NewRabbitMqHandler(usecase)

	return handler
}

func robotUsecase(db *gorm.DB, redis *redis.Client) *usecase.RobotUsecase {

	baseMapper := mapper.NewBaseMapper()
	robotMapper := mapper.NewRobotMapper()
	repo := repository.NewRobotRepository(db, redis)

	comparator := comparator.NewRobotComparator(repo)
	usecase := usecase.NewRobotUsecase(baseMapper, robotMapper, repo, comparator)

	return usecase
}
