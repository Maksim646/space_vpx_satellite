package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type SolarPanelRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ISolarPanelSideRepository {
	return &SolarPanelRepository{sqalxConn: sqalxConn}
}

func (r *SolarPanelRepository) CreateSolarPanelSide(ctx context.Context, solarPanel model.CubeSatSolarPanelSide) (int64, error) {
	query, params, err := postgresql.Builder.Insert("cube_sat_solar_panel_side").
		Columns(
			"length",
			"name",
			"width",
			"height",
			"weight",
			"interface",
			"voc",
			"isc",
			"vmp",
			"imp",
			"efficiency",
			"coil_area",
			"coil_resistance",
			"max_operating_temperature",
			"min_operating_temperature",
			"mechanical_vibration",
			"mechanical_shock",
		).
		Values(
			solarPanel.Length.Float64,
			solarPanel.Name.String,
			solarPanel.Width.Float64,
			solarPanel.Height.Float64,
			solarPanel.Weight.Float64,
			solarPanel.Interface.String,
			solarPanel.Voc.Float64,
			solarPanel.Isc.Float64,
			solarPanel.Vmp.Float64,
			solarPanel.Imp.Float64,
			solarPanel.Efficiency.Int64,
			solarPanel.CoilArea.Float64,
			solarPanel.CoilResistance.Int64,
			solarPanel.MaxOperatingTemperature.Float64,
			solarPanel.MinOperatingTemperature.Float64,
			solarPanel.MechanicalVibration.Int64,
			solarPanel.MechanicalShock.Int64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newSolarPanelID int64
	err = r.sqalxConn.GetContext(ctx, &newSolarPanelID, query, params...)
	return newSolarPanelID, err
}

func (r *SolarPanelRepository) GetSolarPanelSideByID(ctx context.Context, solarPanelSideID int64) (model.CubeSatSolarPanelSide, error) {
	var solarPanel model.CubeSatSolarPanelSide
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"interface",
		"voc",
		"isc",
		"vmp",
		"imp",
		"efficiency",
		"coil_area",
		"coil_resistance",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("cube_sat_solar_panel_side").
		Where(sq.Eq{"id": solarPanelSideID}).
		ToSql()
	if err != nil {
		return solarPanel, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &solarPanel, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return solarPanel, model.ErrSolarPanelNotFound
		}
	}

	return solarPanel, err
}

func (r *SolarPanelRepository) GetSolarPanelSideByName(ctx context.Context, solarPanelSideName string) (model.CubeSatSolarPanelSide, error) {
	var solarPanel model.CubeSatSolarPanelSide
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"interface",
		"voc",
		"isc",
		"vmp",
		"imp",
		"efficiency",
		"coil_area",
		"coil_resistance",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("cube_sat_solar_panel_side").
		Where(sq.Eq{"name": solarPanelSideName}).
		ToSql()
	if err != nil {
		return solarPanel, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &solarPanel, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return solarPanel, model.ErrSolarPanelNotFound
		}
	}

	return solarPanel, err
}

func (r *SolarPanelRepository) GetSolarPanelSideByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatSolarPanelSide, error) {
	builder := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"interface",
		"voc",
		"isc",
		"vmp",
		"imp",
		"efficiency",
		"coil_area",
		"coil_resistance",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("cube_sat_solar_panel_side")

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var solarPanels []model.CubeSatSolarPanelSide
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &solarPanels, query, params...); err != nil {
		return nil, err
	}

	return solarPanels, nil
}

func (r *SolarPanelRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterCubeSatSolarPanelSideByMaxLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"length": length})
			}
		case model.FilterCubeSatSolarPanelSideByMinLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"length": length})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *SolarPanelRepository) UpdateSolarPanelSide(ctx context.Context, solarPanel model.CubeSatSolarPanelSide) error {
	query, params, err := postgresql.Builder.Update("cube_sat_solar_panel_side").
		Set("name", solarPanel.Name.String).
		Set("length", solarPanel.Length.Float64).
		Set("width", solarPanel.Width.Float64).
		Set("height", solarPanel.Height.Float64).
		Set("weight", solarPanel.Weight.Float64).
		Set("interface", solarPanel.Interface.String).
		Set("voc", solarPanel.Voc.Float64).
		Set("isc", solarPanel.Isc.Float64).
		Set("vmp", solarPanel.Vmp.Float64).
		Set("imp", solarPanel.Imp.Float64).
		Set("efficiency", solarPanel.Efficiency.Int64).
		Set("coil_area", solarPanel.CoilArea.Float64).
		Set("coil_resistance", solarPanel.CoilResistance.Int64).
		Set("max_operating_temperature", solarPanel.MaxOperatingTemperature.Float64).
		Set("min_operating_temperature", solarPanel.MinOperatingTemperature.Float64).
		Set("mechanical_vibration", solarPanel.MechanicalVibration.Int64).
		Set("mechanical_shock", solarPanel.MechanicalShock.Int64).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": solarPanel.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *SolarPanelRepository) DeleteSolarPanelSide(ctx context.Context, solarPanelSideID int64) error {
	query, params, err := postgresql.Builder.Delete("cube_sat_solar_panel_side").
		Where(sq.Eq{"cube_sat_solar_panel_side.id": solarPanelSideID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
