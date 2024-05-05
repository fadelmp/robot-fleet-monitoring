package usecase

import (
	"robot-fleet-monitoring/service-position/comparator"
	"robot-fleet-monitoring/service-position/dto"
	"robot-fleet-monitoring/service-position/mapper"
	"robot-fleet-monitoring/service-position/repository"
	"robot-fleet-monitoring/service-position/utils"
)

// Interface
type PositionUsecaseContract interface {
	Update(dto.Position) error
}

// Class
type PositionUsecase struct {
	channel    ChannelUsecase
	baseMapper mapper.BaseMapperContract
	mapper     mapper.PositionMapperContract
	comparator comparator.PositionComparatorContract
	repo       repository.PositionRepositoryContract
}

// Constructor
func NewPositionUsecase(
	channel ChannelUsecaseContract,
	baseMapper mapper.BaseMapperContract,
	mapper mapper.PositionMapperContract,
	comparator comparator.PositionComparatorContract,
	repo repository.PositionRepositoryContract,
) *PositionUsecase {
	return &PositionUsecase{
		baseMapper: baseMapper,
		mapper:     mapper,
		comparator: comparator,
		repo:       repo,
	}
}

// Implementation

func (u *PositionUsecase) Update(dto dto.Position) error {

	// Check Restricted Area
	alert, restricted_id, restricted_name := u.comparator.CheckArea(dto)
	dto.Message = alert

	// Send Alert to Service Monitor
	positionBody := utils.SerializePosition(dto)

	u.channel.Publish(positionBody, "monitor")

	// Map and Save Position Dto to DB
	position := u.mapper.ToPosition(dto, restricted_id, restricted_name)

	u.baseMapper.Create(&position.Base, dto.RobotName)

	u.repo.Create(&position)

	// Send Robot Data to Service Robot
	robotDto := u.mapper.ToRobotDto(dto)

	robotBody := utils.SerializeRobot(robotDto)

	u.channel.Publish(robotBody, "robot")

	// Return
	return nil
}
