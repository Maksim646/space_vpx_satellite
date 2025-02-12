package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type SolarPanelTopUsecase struct {
	solarPanelTopRepository model.ISolarPanelTopRepository
}

func New(solarPanelTopRepository model.ISolarPanelTopRepository) model.ISolarPanelTopUsecase {
	return &SolarPanelTopUsecase{
		solarPanelTopRepository: solarPanelTopRepository,
	}
}

func (u *SolarPanelTopUsecase) CreateSolarPanelTop(ctx context.Context, solarPanel model.SolarPanelTop) (int64, error) {
	return u.solarPanelTopRepository.CreateSolarPanelTop(ctx, solarPanel)
}

func (u *SolarPanelTopUsecase) GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (model.SolarPanelTop, error) {
	return u.solarPanelTopRepository.GetSolarPanelTopByID(ctx, solarPanelTopID)
}

func (u *SolarPanelTopUsecase) GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.SolarPanelTop, error) {
	return u.solarPanelTopRepository.GetSolarPanelTopByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *SolarPanelTopUsecase) UpdateSolarPanelTop(ctx context.Context, solarPanel model.SolarPanelTop) error {
	return u.solarPanelTopRepository.UpdateSolarPanelTop(ctx, solarPanel)
}

func (u *SolarPanelTopUsecase) DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error {
	return u.solarPanelTopRepository.DeleteSolarPanelTop(ctx, solarPanelTopID)
}
