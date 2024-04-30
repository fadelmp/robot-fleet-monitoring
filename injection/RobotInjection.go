package injection

import (
	"robot-fleet-monitoring/comparator"
	"robot-fleet-monitoring/handler"
	"robot-fleet-monitoring/mapper"
	"robot-fleet-monitoring/repository"
	"robot-fleet-monitoring/usecase"

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
