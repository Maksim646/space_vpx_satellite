package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model" // Измените путь на актуальный для вашего проекта
)

type ChassisUsecase struct {
	chassisRepository model.IChassisRepository
}

func New(chassisRepository model.IChassisRepository) model.IChassisUsecase {
	return &ChassisUsecase{
		chassisRepository: chassisRepository,
	}
}

func (u *ChassisUsecase) CreateChassisByID(ctx context.Context, chassis model.Chassis) (int64, error) {
	return u.chassisRepository.CreateChassisByID(ctx, chassis)
}

func (u *ChassisUsecase) UpdateChassisByID(ctx context.Context, chassis model.Chassis) error {
	return u.chassisRepository.UpdateChassisByID(ctx, chassis)
}

func (u *ChassisUsecase) DeleteChassisByID(ctx context.Context, id int64) error {
	return u.chassisRepository.DeleteChassisByID(ctx, id)
}

func (u *ChassisUsecase) GetChassisByID(ctx context.Context, id int64) (model.Chassis, error) {
	return u.chassisRepository.GetChassisByID(ctx, id)
}

func (u *ChassisUsecase) GetChassisByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.Chassis, error) {
	return u.chassisRepository.GetChassisByFilters(ctx, filters)
}
