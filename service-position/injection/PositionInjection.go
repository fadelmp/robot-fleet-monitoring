package injection

import (
	"robot-fleet-monitoring/service-position/comparator"
	"robot-fleet-monitoring/service-position/handler"
	"robot-fleet-monitoring/service-position/mapper"
	"robot-fleet-monitoring/service-position/repository"
	"robot-fleet-monitoring/service-position/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func PositionInjection(db *gorm.DB, redis *redis.Client, channel *amqp.Channel) *handler.PositionHandler {

	baseMapper := mapper.NewBaseMapper()
	positionMapper := mapper.NewPositionMapper()

	positionRepo := repository.NewPositionRepository(db)
	restrictedRepo := repository.NewRestrictedRepository(db, redis)

	channelUsecase := usecase.NewChannelUsecase(channel)
	positionComparator := comparator.NewPositionComparator(restrictedRepo)

	positionUsecase := usecase.NewPositionUsecase(
		channelUsecase,
		baseMapper,
		positionMapper,
		positionComparator,
		positionRepo)

	handler := handler.NewPositionHandler(positionUsecase)

	return handler
}
