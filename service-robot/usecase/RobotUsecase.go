package usecase

import (
	"errors"
	"robot-fleet-monitoring/service-robot/comparator"
	"robot-fleet-monitoring/service-robot/domain"
	"robot-fleet-monitoring/service-robot/dto"
	"robot-fleet-monitoring/service-robot/mapper"
	"robot-fleet-monitoring/service-robot/message"
	"robot-fleet-monitoring/service-robot/repository"
	"robot-fleet-monitoring/service-robot/utils"
)

// Interface
type RobotUsecaseContract interface {
	GetAll() []dto.Robot
	GetById(string) dto.Robot

	Create(dto.Robot) error
	Update(dto.Robot) error
	Delete(dto.Robot) error
}

// Class
type RobotUsecase struct {
	baseMapper mapper.BaseMapperContract
	mapper     mapper.RobotMapperContract
	comparator comparator.RobotComparatorContract
	repo       repository.RobotRepositoryContract
}

// Constructor
func NewRobotUsecase(
	baseMapper mapper.BaseMapperContract,
	mapper mapper.RobotMapperContract,
	repo repository.RobotRepositoryContract,
	comparator comparator.RobotComparatorContract) *RobotUsecase {
	return &RobotUsecase{
		baseMapper: baseMapper,
		mapper:     mapper,
		repo:       repo,
		comparator: comparator,
	}
}

// Implementation

func (u *RobotUsecase) GetAll() []dto.Robot {

	// Get all data
	robots := u.repo.GetAll()

	// Map and Return Robot to Robot Dto
	return u.mapper.ToRobotDtoList(robots)
}

func (u *RobotUsecase) GetById(id string) dto.Robot {

	// Get By Id
	robot := u.repo.GetById(id)

	// Map and Return Robot to RobotDto
	return u.mapper.ToRobotDto(robot)
}

func (u *RobotUsecase) Create(dto dto.Robot) error {

	// Check Name and Return Error if Name Exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Generate UUID
	dto.Id, _ = utils.GenerateUUID()

	// Map Robot Dto to Robot Domain
	robot := u.mapper.ToRobot(dto)

	// Set Created Value
	u.baseMapper.Create(&robot.Base, dto.Base.CreatedBy)

	// Create Robot
	if u.repo.Create(&robot) != nil {
		return errors.New(message.CreateFailed)
	}

	return nil
}

func (u *RobotUsecase) Update(dto dto.Robot) error {

	// Check Id whether not found
	if err := u.comparator.CheckId(dto.Id); err != nil {
		return err
	}

	// Check Name whether name exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Map Robot dto to Robot domain
	robot := u.mapper.ToRobot(dto)

	// Set Updated Value
	u.baseMapper.Update(&robot.Base, dto.Base.UpdatedBy)

	// Update Robot and return
	if u.repo.Update(&robot) != nil {
		return errors.New(message.UpdateFailed)
	}

	return nil
}

func (u *RobotUsecase) Delete(dto dto.Robot) error {

	var robot domain.Robot

	// Check Id whether not found
	if err := u.comparator.CheckId(dto.Id); err != nil {
		return err
	}

	// Set Deleted Value
	u.baseMapper.Delete(&robot, dto.Id, dto.Base.UpdatedBy)

	// Delete Robot and return
	if u.repo.Update(&robot) != nil {
		return errors.New(message.DeleteFailed)
	}

	return nil
}
