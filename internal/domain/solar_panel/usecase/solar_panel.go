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

func (u *SolarPanelUsecase) CreateSolarPanel(ctx context.Context, solarPanel model.SolarPanel) (int64, error) {
	return u.solarPanelRepository.CreateSolarPanel(ctx, solarPanel)
}

func (u *SolarPanelUsecase) GetSolarPanelByID(ctx context.Context, id int64) (model.SolarPanel, error) {
	return u.solarPanelRepository.GetSolarPanelByID(ctx, id)
}
