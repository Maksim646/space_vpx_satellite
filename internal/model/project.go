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
	FilterProjectsByUser = "filter_projects_by_user"
)

type Project struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`

	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type IProjectRepository interface {
	CreatedProject(ctx context.Context, name string, userID string) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (Project, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]Project, int64, error)

	UpdateProjectByID(ctx context.Context, projectID string, name string) error

	DeleteProject(ctx context.Context, projectID string) error
}

type IProjectUsecase interface {
	CreatedProject(ctx context.Context, name string, userID string) (string, error)

	GetProjectByID(ctx context.Context, projectID string) (Project, error)
	GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]Project, int64, error)

	UpdateProjectByID(ctx context.Context, projectID string, name string) error

	DeleteProject(ctx context.Context, projectID string) error
}
