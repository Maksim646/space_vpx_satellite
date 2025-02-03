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
