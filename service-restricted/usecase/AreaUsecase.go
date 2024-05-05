package usecase

import (
	"errors"
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/mapper"
	"robot-fleet-monitoring/service-restricted/message"
	"robot-fleet-monitoring/service-restricted/repository"
)

// Interface
type AreaUsecaseContract interface {
	Create(dto.Restricted, domain.Restricted) error
	Update(dto.Restricted, domain.Restricted) error
	Delete(domain.Restricted) error
}

// Class
type AreaUsecase struct {
	baseMapper mapper.BaseMapperContract
	mapper     mapper.AreaMapperContract
	repo       repository.AreaRepositoryContract
}

// Constructor
func NewAreaUsecase(
	baseMapper mapper.BaseMapperContract,
	mapper mapper.AreaMapperContract,
	repo repository.AreaRepositoryContract) *AreaUsecase {
	return &AreaUsecase{
		baseMapper: baseMapper,
		mapper:     mapper,
		repo:       repo,
	}
}

// Implementation

func (u *AreaUsecase) Create(dto dto.Restricted, restricted domain.Restricted) error {

	// Map Restricted Dto to Restricted Domain
	areas := u.mapper.ToAreaList(dto.Area, restricted)

	// Set Created Value
	u.baseMapper.Create(&restricted.Base, dto.Base.CreatedBy)

	// Create Restricted
	if u.repo.Create(areas) != nil {
		return errors.New(message.CreateFailed)
	}

	return nil
}

func (u *AreaUsecase) Update(dto dto.Restricted, restricted domain.Restricted) error {

	if u.Delete(restricted) != nil {
		return errors.New(message.UpdateFailed)
	}

	if u.Create(dto, restricted) != nil {
		return errors.New(message.UpdateFailed)
	}

	return nil
}

func (u *AreaUsecase) Delete(restricted domain.Restricted) error {

	var area domain.Area
	area.Restricted.Id = restricted.Id

	// Delete Restricted and return
	if u.repo.Delete(area) != nil {
		return errors.New(message.DeleteFailed)
	}

	return nil
}
