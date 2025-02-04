package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Maksim646/space_vpx_satellite/internal/database/postgresql"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type SolarPanelRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ISolarPanelRepository {
	return &SolarPanelRepository{sqalxConn: sqalxConn}
}

func (r *SolarPanelRepository) CreateSolarPanel(ctx context.Context, solarPanel model.SolarPanel) (int64, error) {
	query, params, err := postgresql.Builder.Insert("solar_panel").
		Columns(
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
		).
		Values(
			solarPanel.Length.Float64,
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
			solarPanel.MechanicalVibration.Int64,
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

func (r *SolarPanelRepository) GetSolarPanelByID(ctx context.Context, id int64) (model.SolarPanel, error) {
	var solarPanel model.SolarPanel
	query, params, err := postgresql.Builder.Select(
		"id",
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
		From("solar_panel").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return solarPanel, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &solarPanel, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return solarPanel, model.ErrChassisNotFound
		}
	}

	return solarPanel, err
}
