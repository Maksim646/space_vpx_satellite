package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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
		"cube_sat_projects.id",
		"cube_sat_projects.user_id",
		"cube_sat_projects.name",
		"cube_sat_projects.size",
		"cube_sat_projects.frame_name",
		"cube_sat_projects.solar_panel_side_name",
		"cube_sat_projects.solar_panel_top_name",
		"cube_sat_projects.power_system_name",
		"cube_sat_projects.board_computing_module_name",
		"cube_sat_projects.vhf_antenna_system_name",
		"cube_sat_projects.vhf_transceiver_name",
		"cube_sat_projects.created_at",
		"cube_sat_projects.updated_at",
	).
		From("cube_sat_projects")

	return builder
}

func (r *ProjectRepository) CreatedProject(ctx context.Context, project model.CubeSatProject) (string, error) {
	query, params, err := postgresql.Builder.Insert("cube_sat_projects").
		Columns(
			"name",
			"user_id",
			"frame_name",
			"size",
			"solar_panel_side_name",
			"solar_panel_top_name",
			"power_system_name",
			"board_computing_module_name",
			"vhf_antenna_system_name",
			"vhf_transceiver_name",
		).
		Values(
			project.Name,
			project.UserID,
			project.FrameName,
			project.Size,
			project.SolarPanelSideName,
			project.SolarPanelTopName,
			project.PowerSystemName,
			project.BoardComputingModuleName,
			project.VHFAntennaSystemName,
			project.VhfTransceiverName,
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

func (r *ProjectRepository) GetProjectByID(ctx context.Context, projectID string) (model.CubeSatProject, error) {
	var project model.CubeSatProject
	query, params, err := postgresql.Builder.Select(
		"cube_sat_projects.id",
		"cube_sat_projects.user_id",
		"cube_sat_projects.name",
		"cube_sat_projects.size",
		"cube_sat_projects.frame_name",
		"cube_sat_projects.solar_panel_side_name",
		"cube_sat_projects.solar_panel_top_name",
		"cube_sat_projects.power_system_name",
		"cube_sat_projects.board_computing_module_name",
		"cube_sat_projects.vhf_antenna_system_name",
		"cube_sat_projects.vhf_transceiver_name",
		"cube_sat_projects.updated_at",
		"cube_sat_projects.created_at",
	).
		From("cube_sat_projects").
		Where(sq.Eq{"cube_sat_projects.id": projectID}).ToSql()
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

func (r *ProjectRepository) GetProjectsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatProject, int64, error) {
	cubeSatProjectsCount, err := r.GetCountByFilters(ctx, filters)
	if err != nil {
		return []model.CubeSatProject{}, cubeSatProjectsCount, err
	}
	builder := r.GetProjectBuilder()

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.
		OrderBy(sortParams).
		Offset(uint64(offset)).
		Limit(uint64(limit)).ToSql()

	if err != nil {
		return []model.CubeSatProject{}, cubeSatProjectsCount, err
	}

	var cubeSatProjects []model.CubeSatProject
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &cubeSatProjects, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cubeSatProjects, cubeSatProjectsCount, model.ErrProjectNotFound
		}
	}

	return cubeSatProjects, cubeSatProjectsCount, nil
}

func (r *ProjectRepository) UpdateProjectByID(ctx context.Context, cubeSatProject model.CubeSatProject) error {
	query, params, err := postgresql.Builder.Update("cube_sat_projects").
		Set("name", cubeSatProject.Name).
		Set("frame_name", cubeSatProject.FrameName.String).
		Set("solar_panel_side_name", cubeSatProject.SolarPanelSideName.String).
		Set("solar_panel_top_name", cubeSatProject.SolarPanelTopName.String).
		Set("power_system_name", cubeSatProject.PowerSystemName.String).
		Set("board_computing_module_name", cubeSatProject.BoardComputingModuleName.String).
		Set("vhf_antenna_system_name", cubeSatProject.VHFAntennaSystemName.String).
		Set("vhf_transceiver_name", cubeSatProject.VhfTransceiverName.String).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": cubeSatProject.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *ProjectRepository) DeleteProject(ctx context.Context, projectID string) error {
	query, params, err := postgresql.Builder.Delete("cube_sat_projects").
		Where(sq.Eq{"cube_sat_projects.id": projectID}).
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
		case model.FilterCubeSatProjectsByUser:
			if userID, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"cube_sat_projects.user_id": userID})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *ProjectRepository) GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var count int64
	builder := postgresql.Builder.Select("COUNT(id)").From("cube_sat_projects")
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
