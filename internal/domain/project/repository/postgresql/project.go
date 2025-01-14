package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	//"strings"
	//"time"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type ProjectRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IProjectRepository {
	return &ProjectRepository{sqalxConn: sqalxConn}
}

func (r *ProjectRepository) GetProjectBuilder() sq.SelectBuilder {
	builder := postgresql.Builder.Select(
		"projects.id",
		"projects.user_id",
		"projects.name",
		"projects.created_at",
		"projects.updated_at",
	).
		From("projects")

	return builder
}

func (r *ProjectRepository) CreatedProject(ctx context.Context, name string, userID string) (string, error) {
	query, params, err := postgresql.Builder.Insert("projects").
		Columns(
			"name",
			"user_id",
		).
		Values(
			name,
			userID,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newProjectID string
	err = r.sqalxConn.GetContext(ctx, &newProjectID, query, params...)
	return newProjectID, err
}

func (r *ProjectRepository) GetProjectByID(ctx context.Context, projectID string) (model.Project, error) {
	var project model.Project
	query, params, err := postgresql.Builder.Select(
		"projects.id",
		"projects.user_id",
		"projects.name",
		"projects.updated_at",
		"projects.created_at",
	).
		From("projects").
		Where(sq.Eq{"projects.id": projectID}).ToSql()
	if err != nil {
		return project, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &project, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project, model.ErrProjectNotFound
		}
	}

	fmt.Println(err)
	return project, err
}

func (r *ProjectRepository) GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.Project, int64, error) {
	projectsCount, err := r.GetCountByFilters(ctx, filters)
	if err != nil {
		return []model.Project{}, projectsCount, err
	}
	builder := r.GetProjectBuilder()

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.
		OrderBy(sortParams).
		Offset(uint64(offset)).
		Limit(uint64(limit)).ToSql()

	if err != nil {
		return []model.Project{}, projectsCount, err
	}

	var projects []model.Project
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &projects, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return projects, projectsCount, model.ErrProjectNotFound
		}
	}

	return projects, projectsCount, nil
}

func (r *ProjectRepository) UpdateProjectByID(ctx context.Context, projectID string, name string) error {
	query, params, err := postgresql.Builder.Update("projects").
		Set("name", name).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": projectID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *ProjectRepository) DeleteProject(ctx context.Context, projectID string) error {
	query, params, err := postgresql.Builder.Delete("projects").
		Where(sq.Eq{"projects.id": projectID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *ProjectRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterProjectsByUser:
			if userID, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"projects.user_id": userID})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *ProjectRepository) GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var count int64
	builder := postgresql.Builder.Select("COUNT(id)").From("projects")
	builder = r.ApplyFilters(builder, filters)
	query, params, err := builder.ToSql()
	if err != nil {
		return count, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &count, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, model.ErrProjectNotFound
		}
	}
	return count, nil
}
