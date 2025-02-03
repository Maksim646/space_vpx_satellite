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

type CubeSatFrameRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ICubeSatFrameRepository {
	return &CubeSatFrameRepository{sqalxConn: sqalxConn}
}

func (r *CubeSatFrameRepository) CreateCubeSatFrame(ctx context.Context, cubeSatFrame model.CubeSatFrame) (int64, error) {
	query, params, err := postgresql.Builder.Insert("cube_sat_frame").
		Columns(
			"dimension",
			"weight",
			"operating_temperature_min",
			"operating_temperature_max",
			"mechanical_vibration",
			"mechanical_shock",
			"link",
		).
		Values(
			cubeSatFrame.Dimension,
			cubeSatFrame.Weight,
			cubeSatFrame.OperatingTemperatureMin,
			cubeSatFrame.OperatingTemperatureMax,
			cubeSatFrame.MechanicalVibration,
			cubeSatFrame.MechanicalVibration,
			cubeSatFrame.Link,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var cubeSatFrameID int64
	err = r.sqalxConn.GetContext(ctx, &cubeSatFrameID, query, params...)
	return cubeSatFrameID, err
}

func (r *CubeSatFrameRepository) UpdateCubeSatFrame(ctx context.Context, cubeSatFrame model.CubeSatFrame) error {
	query, params, err := postgresql.Builder.Update("cube_sat_frame").
		Set("dimension", cubeSatFrame.Dimension).
		Set("weight", cubeSatFrame.Weight).
		Set("operating_temperature_min", cubeSatFrame.OperatingTemperatureMin).
		Set("operating_temperature_max", cubeSatFrame.OperatingTemperatureMax).
		Set("mechanical_vibration", cubeSatFrame.MechanicalVibration).
		Set("mechanical_shock", cubeSatFrame.MechanicalShock).
		Set("link", cubeSatFrame.Link).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": cubeSatFrame.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *CubeSatFrameRepository) GetCubeSatFrameByID(ctx context.Context, id int64) (model.CubeSatFrame, error) {
	var cubeSatFrame model.CubeSatFrame
	query, params, err := postgresql.Builder.Select(
		"id",
		"dimension",
		"weight",
		"operating_temperature_min",
		"operating_temperature_max",
		"mechanical_vibration",
		"mechanical_shock",
		"link",
	).
		From("cube_sat_frame").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return cubeSatFrame, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &cubeSatFrame, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cubeSatFrame, model.ErrChassisNotFound
		}
	}

	return cubeSatFrame, err
}

func (r *CubeSatFrameRepository) GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatFrame, error) {
	builder := postgresql.Builder.Select(
		"id",
		"dimension",
		"weight",
		"operating_temperature_min",
		"operating_temperature_max",
		"mechanical_vibration",
		"mechanical_shock",
		"link",
	).From("cube_sat_frame")

	// builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var cubeSatFrames []model.CubeSatFrame
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &cubeSatFrames, query, params...); err != nil {
		return nil, err
	}

	return cubeSatFrames, nil
}

// func (r *CubeSatFrameRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
// 	for filterKey, filterValue := range filters {
// 		switch filterKey {
// 		case model.FilterCubeSatFrameByMaxLength:
// 			if length, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"length": length})
// 			}
// 		case model.FilterCubeSatFrameByMinLength:
// 			if length, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"length": length})
// 			}
// 		case model.FilterChassisByMaxWidth:
// 			if width, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"width": width}) // Условие для максимальной ширины
// 			}
// 		case model.FilterChassisByMinWidth:
// 			if width, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"width": width}) // Условие для минимальной ширины
// 			}
// 		case model.FilterChassisByMaxHeight:
// 			if height, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"height": height}) // Условие для максимальной высоты
// 			}
// 		case model.FilterChassisByMinHeight:
// 			if height, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"height": height}) // Условие для минимальной высоты
// 			}
// 		case model.FilterChassisByMaxWeight:
// 			if weight, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"weight": weight}) // Условие для максимального веса
// 			}
// 		case model.FilterChassisByMinWeight:
// 			if weight, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"weight": weight}) // Условие для минимального веса
// 			}
// 		case model.FilterChassisByMaxPowerHandlingCapabilityPerBoard:
// 			if power, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"power_handling_capability_per_board": power}) // Условие для максимальной мощности
// 			}
// 		case model.FilterChassisByMinPowerHandlingCapabilityPerBoard:
// 			if power, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"power_handling_capability_per_board": power}) // Условие для минимальной мощности
// 			}
// 		case model.FilterChassisByMaxTemperaturePerBoard:
// 			if temperature, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.LtOrEq{"temperature_per_board": temperature}) // Условие для максимальной температуры
// 			}
// 		case model.FilterChassisByMinTemperaturePerBoard:
// 			if temperature, ok := filterValue.(float64); ok {
// 				builder = builder.Where(sq.GtOrEq{"temperature_per_board": temperature}) // Условие для минимальной температуры
// 			}
// 		default:
// 			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
// 		}
// 	}

// 	return builder
// }

// func (r *CubeSatFrameRepository) DeleteCubeSatFrame(ctx context.Context, id int64) error {

// }
