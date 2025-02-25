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

	InvalidSize = "cube sat invalid size, try "

	InvalidInterface = "cube sat invalid interface, try "
)

const (
	FilterCubeSatProjectsByUser = "filter_projects_by_user"
)

type CubeSatProject struct {
	ID                       string         `db:"id"`
	UserID                   string         `db:"user_id"`
	FrameName                sql.NullString `db:"frame_name"`
	Size                     int64          `db:"size"`
	SolarPanelSideName       sql.NullString `db:"solar_panel_side_name"`
	SolarPanelTopName        sql.NullString `db:"solar_panel_top_name"`
	PowerSystemName          sql.NullString `db:"power_system_name"`
	BoardComputingModuleName sql.NullString `db:"board_computing_module_name"`
	VHFAntennaSystemName     sql.NullString `db:"vhf_antenna_system_name"`
	VhfTransceiverName       sql.NullString `db:"vhf_transceiver_name"`
	Name                     string         `db:"name"`
	CreatedAt                time.Time      `db:"created_at"`
	UpdatedAt                sql.NullTime   `db:"updated_at"`
}

type IProjectRepository interface {
	CreatedProject(ctx context.Context, project CubeSatProject) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (CubeSatProject, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatProject, int64, error)

	UpdateProjectByID(ctx context.Context, cubeSatProject CubeSatProject) error

	DeleteProject(ctx context.Context, projectID string) error
}

type IProjectUsecase interface {
	CreatedProject(ctx context.Context, project CubeSatProject) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (CubeSatProject, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatProject, int64, error)

	UpdateProjectByID(ctx context.Context, cubeSatProject CubeSatProject) error

	DeleteProject(ctx context.Context, projectID string) error
}
