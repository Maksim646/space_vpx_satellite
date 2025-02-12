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

type SolarPanelTopRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ISolarPanelTopRepository {
	return &SolarPanelTopRepository{sqalxConn: sqalxConn}
}

func (r *SolarPanelTopRepository) CreateSolarPanelTop(ctx context.Context, solarPanel model.SolarPanelTop) (int64, error) {
	query, params, err := postgresql.Builder.Insert("cube_sat_solar_panel_top").
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
			"min_operating_temperature",
			"max_operating_temperature",
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
			solarPanel.Efficiency.Float64,
			solarPanel.CoilArea.Float64,
			solarPanel.CoilResistance.Float64,
			solarPanel.MinOperatingTemperature.Float64,
			solarPanel.MaxOperatingTemperature.Float64,
			solarPanel.MechanicalVibration.Float64,
			solarPanel.MechanicalShock.Float64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newSolarPanelTopID int64
	err = r.sqalxConn.GetContext(ctx, &newSolarPanelTopID, query, params...)
	return newSolarPanelTopID, err
}

func (r *SolarPanelTopRepository) GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (model.SolarPanelTop, error) {
	var solarPanelTop model.SolarPanelTop
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
		"min_operating_temperature",
		"max_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("cube_sat_solar_panel_top").
		Where(sq.Eq{"id": solarPanelTopID}).
		ToSql()
	if err != nil {
		return solarPanelTop, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &solarPanelTop, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return solarPanelTop, model.ErrSolarPanelTopNotFound
		}
	}

	return solarPanelTop, err
}

func (r *SolarPanelTopRepository) GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.SolarPanelTop, error) {
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
		"min_operating_temperature",
		"max_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("cube_sat_solar_panel_top")

	builder = r.ApplyFilters(builder, filters)

	if sortParams != "" {
		builder = builder.OrderBy(sortParams)
	}

	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var solarPanelsTop []model.SolarPanelTop
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &solarPanelsTop, query, params...); err != nil {
		return nil, err
	}

	return solarPanelsTop, nil
}

// ApplyFilters applies filters to the SQL query builder.
func (r *SolarPanelTopRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterSolarPanelTopByMaxLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"length": length})
			}
		case model.FilterSolarPanelTopByMinLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"length": length})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *SolarPanelTopRepository) UpdateSolarPanelTop(ctx context.Context, solarPanel model.SolarPanelTop) error {
	query, params, err := postgresql.Builder.Update("cube_sat_solar_panel_top").
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
		Set("efficiency", solarPanel.Efficiency.Float64).
		Set("coil_area", solarPanel.CoilArea.Float64).
		Set("coil_resistance", solarPanel.CoilResistance.Float64).
		Set("min_operating_temperature", solarPanel.MinOperatingTemperature.Float64).
		Set("max_operating_temperature", solarPanel.MaxOperatingTemperature.Float64).
		Set("mechanical_vibration", solarPanel.MechanicalVibration.Float64).
		Set("mechanical_shock", solarPanel.MechanicalShock.Float64).
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

func (r *SolarPanelTopRepository) DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error {
	query, params, err := postgresql.Builder.Delete("cube_sat_solar_panel_top").
		Where(sq.Eq{"id": solarPanelTopID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
