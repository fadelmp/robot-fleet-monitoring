package usecase

import (
	"errors"
	"robot-fleet-monitoring/service-restricted/comparator"
	"robot-fleet-monitoring/service-restricted/domain"
	"robot-fleet-monitoring/service-restricted/dto"
	"robot-fleet-monitoring/service-restricted/mapper"
	"robot-fleet-monitoring/service-restricted/message"
	"robot-fleet-monitoring/service-restricted/repository"
)

// Interface
type RestrictedUsecaseContract interface {
	GetAll() []dto.Restricted
	GetById(string) dto.Restricted

	Create(dto.Restricted) error
	Update(dto.Restricted) error
	Delete(dto.Restricted) error
}

// Class
type RestrictedUsecase struct {
	areaUsecase AreaUsecaseContract
	baseMapper  mapper.BaseMapperContract
	mapper      mapper.RestrictedMapperContract
	comparator  comparator.RestrictedComparatorContract
	repo        repository.RestrictedRepositoryContract
}

// Constructor
func NewRestrictedUsecase(
	areaUsecase AreaUsecaseContract,
	baseMapper mapper.BaseMapperContract,
	mapper mapper.RestrictedMapperContract,
	repo repository.RestrictedRepositoryContract,
	comparator comparator.RestrictedComparatorContract) *RestrictedUsecase {
	return &RestrictedUsecase{
		areaUsecase: areaUsecase,
		baseMapper:  baseMapper,
		mapper:      mapper,
		repo:        repo,
		comparator:  comparator,
	}
}

// Implementation

func (u *RestrictedUsecase) GetAll() []dto.Restricted {

	// Get all data
	restricteds := u.repo.GetAll()

	// Map and Return Restricted to Restricted Dto
	return u.mapper.ToRestrictedDtoList(restricteds)
}

func (u *RestrictedUsecase) GetById(id string) dto.Restricted {

	// Get By Id
	restricted := u.repo.GetById(id)

	// Map and Return Restricted to RestrictedDto
	return u.mapper.ToRestrictedDto(restricted)
}

func (u *RestrictedUsecase) Create(dto dto.Restricted) error {

	// Check Name and Return Error if Name Exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Map Restricted Dto to Restricted Domain
	restricted := u.mapper.ToRestricted(dto)

	// Set Created Value
	u.baseMapper.Create(&restricted.Base, dto.Base.CreatedBy)

	// Create Restricted
	if u.repo.Create(&restricted) != nil {
		return errors.New(message.CreateFailed)
	}

	// Create Area
	if u.areaUsecase.Create(dto, restricted) != nil {
		return errors.New(message.CreateFailed)
	}

	return nil
}

func (u *RestrictedUsecase) Update(dto dto.Restricted) error {

	// Check Id whether not found
	if err := u.comparator.CheckId(dto.Id); err != nil {
		return err
	}

	// Check Name whether name exists
	if err := u.comparator.CheckName(dto); err != nil {
		return err
	}

	// Map Restricted dto to Restricted domain
	restricted := u.mapper.ToRestricted(dto)

	// Set Updated Value
	u.baseMapper.Update(&restricted.Base, dto.Base.UpdatedBy)

	// Update Restricted
	if u.repo.Update(&restricted) != nil {
		return errors.New(message.UpdateFailed)
	}

	// Update Area
	if u.areaUsecase.Update(dto, restricted) != nil {
		return errors.New(message.UpdateFailed)
	}

	return nil
}

func (u *RestrictedUsecase) Delete(dto dto.Restricted) error {

	var restricted domain.Restricted

	// Check Id whether not found
	if err := u.comparator.CheckId(dto.Id); err != nil {
		return err
	}

	// Set Deleted Value
	u.baseMapper.Delete(&restricted.Base, dto.Base.UpdatedBy)

	// Set Id
	restricted.Id = dto.Id

	// Delete Restricted
	if u.repo.Update(&restricted) != nil {
		return errors.New(message.DeleteFailed)
	}

	// Delete Area
	if u.areaUsecase.Delete(restricted) != nil {
		return errors.New(message.UpdateFailed)
	}

	return nil
}
