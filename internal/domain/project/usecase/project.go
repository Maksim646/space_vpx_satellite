package usecase

import (
	"context"

	"github.com/Maksim646/space_vpx_satellite/internal/model"
)

type Usecase struct {
	projectRepository model.IProjectRepository
}

func New(projectRepository model.IProjectRepository) model.IProjectUsecase {
	return &Usecase{
		projectRepository: projectRepository,
	}
}

func (u *Usecase) CreatedProject(ctx context.Context, project model.CubeSatProject) (string, error) {
	return u.projectRepository.CreatedProject(ctx, project)
}

func (u *Usecase) GetProjectByID(ctx context.Context, projectID string) (model.CubeSatProject, error) {
	return u.projectRepository.GetProjectByID(ctx, projectID)
}

func (u *Usecase) GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatProject, int64, error) {
	return u.projectRepository.GetProjectsByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *Usecase) UpdateProjectByID(ctx context.Context, cubeSatProject model.CubeSatProject) error {
	return u.projectRepository.UpdateProjectByID(ctx, cubeSatProject)
}

func (u *Usecase) DeleteProject(ctx context.Context, projectID string) error {
	return u.projectRepository.DeleteProject(ctx, projectID)
}
