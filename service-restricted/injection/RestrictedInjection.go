package injection

import (
	"robot-fleet-monitoring/service-restricted/comparator"
	"robot-fleet-monitoring/service-restricted/handler"
	"robot-fleet-monitoring/service-restricted/mapper"
	"robot-fleet-monitoring/service-restricted/repository"
	"robot-fleet-monitoring/service-restricted/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func RestrictedInjection(db *gorm.DB, redis *redis.Client) *handler.RestrictedHandler {

	baseMapper := mapper.NewBaseMapper()
	restrictedMapper := mapper.NewRestrictedMapper()
	repo := repository.NewRestrictedRepository(db, redis)

	comparator := comparator.NewRestrictedComparator(repo)
	usecase := usecase.NewRestrictedUsecase(baseMapper, restrictedMapper, repo, comparator)
	handler := handler.NewRestrictedHandler(usecase)

	return handler
}
