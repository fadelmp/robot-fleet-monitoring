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

	baseMapper := mapper.NewBaseMapper()
	robotMapper := mapper.NewRobotMapper()
	repo := repository.NewRobotRepository(db, redis)

	comparator := comparator.NewRobotComparator(repo)
	usecase := usecase.NewRobotUsecase(baseMapper, robotMapper, repo, comparator)
	handler := handler.NewRobotHandler(usecase)

	return handler
}
