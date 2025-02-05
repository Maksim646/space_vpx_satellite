package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrProjectsNotFound = errors.New("projects not found")

	ProjectNotFound  = "project not found"
	ProjectsNotFound = "projects not found"
)

const (
	FilterCubeSatProjectsByUser = "filter_projects_by_user"
)

type CubeSatProject struct {
	ID                    string         `db:"id"`
	UserID                string         `db:"user_id"`
	CubeSatFrameName      sql.NullString `db:"cube_sat_frame_id"`
	CubeSatSolarPanelName sql.NullString `db:"solar_panael_id"`
	Name                  string         `db:"name"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             sql.NullTime   `db:"updated_at"`
}

type IProjectRepository interface {
	CreatedProject(ctx context.Context, name string, userID string) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (CubeSatProject, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatProject, int64, error)

	UpdateProjectByID(ctx context.Context, projectID string, name string) error

	DeleteProject(ctx context.Context, projectID string) error
}

type IProjectUsecase interface {
	CreatedProject(ctx context.Context, name string, userID string) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (CubeSatProject, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatProject, int64, error)

	UpdateProjectByID(ctx context.Context, projectID string, name string) error

	DeleteProject(ctx context.Context, projectID string) error
}
