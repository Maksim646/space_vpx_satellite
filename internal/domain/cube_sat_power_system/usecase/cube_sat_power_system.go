package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type PowerSystemUsecase struct {
	powerSystemRepository model.IPowerSystemRepository
}

func New(powerSystemRepository model.IPowerSystemRepository) model.IPowerSystemUsecase {
	return &PowerSystemUsecase{
		powerSystemRepository: powerSystemRepository,
	}
}

func (u *PowerSystemUsecase) CreatePowerSystem(ctx context.Context, powerSystem model.CubeSatPowerSystem) (int64, error) {
	return u.powerSystemRepository.CreatePowerSystem(ctx, powerSystem)
}

func (u *PowerSystemUsecase) GetPowerSystemByID(ctx context.Context, powerSystemID int64) (model.CubeSatPowerSystem, error) {
	return u.powerSystemRepository.GetPowerSystemByID(ctx, powerSystemID)
}

func (u *PowerSystemUsecase) GetPowerSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatPowerSystem, error) {
	return u.powerSystemRepository.GetPowerSystemsByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *PowerSystemUsecase) UpdatePowerSystem(ctx context.Context, powerSystem model.CubeSatPowerSystem) error {
	return u.powerSystemRepository.UpdatePowerSystem(ctx, powerSystem)
}

func (u *PowerSystemUsecase) DeletePowerSystem(ctx context.Context, powerSystemID int64) error {
	return u.powerSystemRepository.DeletePowerSystem(ctx, powerSystemID)
}
