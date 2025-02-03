package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type Usecase struct {
	cubeSatFrameRepository model.ICubeSatFrameRepository
}

func New(cubeSatFrameRepository model.ICubeSatFrameRepository) model.ICubeSatFrameUsecase {
	return &Usecase{
		cubeSatFrameRepository: cubeSatFrameRepository,
	}
}

func (u *Usecase) CreateCubeSatFrame(ctx context.Context, cubeSatFrame model.CubeSatFrame) (int64, error) {
	return u.cubeSatFrameRepository.CreateCubeSatFrame(ctx, cubeSatFrame)
}

func (u *Usecase) UpdateCubeSatFrame(ctx context.Context, cubeSatFrame model.CubeSatFrame) error {
	return u.cubeSatFrameRepository.UpdateCubeSatFrame(ctx, cubeSatFrame)
}

func (u *Usecase) GetCubeSatFrameByID(ctx context.Context, id int64) (model.CubeSatFrame, error) {
	return u.cubeSatFrameRepository.GetCubeSatFrameByID(ctx, id)
}

// func (u *Usecase) GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatFrame, error) {
// 	return u.GetCubeSatFramesByFilters(ctx, offset, limit, sortParams, filters)
// }

// func (u *Usecase) DeleteCubeSatFrame(ctx context.Context, id int64) error {
// 	return u.DeleteCubeSatFrame(ctx, id)
// }
