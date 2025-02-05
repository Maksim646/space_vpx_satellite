package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type SolarPanelUsecase struct {
	solarPanelRepository model.ISolarPanelRepository
}

func New(solarPanelRepository model.ISolarPanelRepository) model.ISolarPanelUsecase {
	return &SolarPanelUsecase{
		solarPanelRepository: solarPanelRepository,
	}
}

func (u *SolarPanelUsecase) CreateSolarPanelSide(ctx context.Context, solarPanel model.CubeSatSolarPanelSide) (int64, error) {
	return u.solarPanelRepository.CreateSolarPanelSide(ctx, solarPanel)
}

func (u *SolarPanelUsecase) GetSolarPanelSideByID(ctx context.Context, solarPanelSideID int64) (model.CubeSatSolarPanelSide, error) {
	return u.solarPanelRepository.GetSolarPanelSideByID(ctx, solarPanelSideID)
}

func (u *SolarPanelUsecase) GetSolarPanelSideByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatSolarPanelSide, error) {
	return u.solarPanelRepository.GetSolarPanelSideByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *SolarPanelUsecase) UpdateSolarPanelSide(ctx context.Context, solarPanel model.CubeSatSolarPanelSide) error {
	return u.solarPanelRepository.UpdateSolarPanelSide(ctx, solarPanel)
}

func (u *SolarPanelUsecase) DeleteSolarPanelSide(ctx context.Context, solarPanelSideID int64) error {
	return u.solarPanelRepository.DeleteSolarPanelSide(ctx, solarPanelSideID)
}
