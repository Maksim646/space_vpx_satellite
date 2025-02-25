package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type BoardComputingModuleUsecase struct {
	boardComputingModuleRepository model.IBoardComputingModuleRepository
}

func New(boardComputingModuleRepository model.IBoardComputingModuleRepository) model.IBoardComputingModuleUsecase {
	return &BoardComputingModuleUsecase{
		boardComputingModuleRepository: boardComputingModuleRepository,
	}
}

func (u *BoardComputingModuleUsecase) CreateBoardComputingModule(ctx context.Context, module model.BoardComputingModule) (int64, error) {
	return u.boardComputingModuleRepository.CreateBoardComputingModule(ctx, module)
}

func (u *BoardComputingModuleUsecase) GetBoardComputingModuleByID(ctx context.Context, moduleID int64) (model.BoardComputingModule, error) {
	return u.boardComputingModuleRepository.GetBoardComputingModuleByID(ctx, moduleID)
}

func (u *BoardComputingModuleUsecase) GetBoardComputingModuleByName(ctx context.Context, moduleName string) (model.BoardComputingModule, error) {
	return u.boardComputingModuleRepository.GetBoardComputingModuleByName(ctx, moduleName)
}

func (u *BoardComputingModuleUsecase) GetBoardComputingModulesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.BoardComputingModule, error) {
	return u.boardComputingModuleRepository.GetBoardComputingModulesByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *BoardComputingModuleUsecase) UpdateBoardComputingModule(ctx context.Context, module model.BoardComputingModule) error {
	return u.boardComputingModuleRepository.UpdateBoardComputingModule(ctx, module)
}

func (u *BoardComputingModuleUsecase) DeleteBoardComputingModule(ctx context.Context, moduleID int64) error {
	return u.boardComputingModuleRepository.DeleteBoardComputingModule(ctx, moduleID)
}
