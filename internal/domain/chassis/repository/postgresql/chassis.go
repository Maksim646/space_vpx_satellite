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

type ChassisRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IChassisRepository {
	return &ChassisRepository{sqalxConn: sqalxConn}
}

func (r *ChassisRepository) CreateChassisByID(ctx context.Context, chassis model.Chassis) (int64, error) {
	query, params, err := postgresql.Builder.Insert("chassis").
		Columns(
			"name",
			"slots",
			"size",
			"max_operating_temperature",
			"min_operating_temperature",
			"max_non_operating_temperature",
			"min_non_operating_temperature",
			"overload",
			"max_vibration_sine",
			"min_vibration_sine",
			"max_vibration_random",
			"min_vibration_random",
			"axes",
			"shock_response_spectrum",
			"peak_overload_spectrum_1",
			"peak_overload_spectrum_2",
			"peak_frequency_spectrum_1",
			"peak_frequency_spectrum_2",
			"length",
			"width",
			"height",
			"weight",
			"power_handling_capability_per_board",
			"temperature_per_board",
		).
		Values(
			chassis.Name,
			chassis.Slots,
			chassis.Size,
			chassis.MaxOperatingTemperature,
			chassis.MinOperatingTemperature,
			chassis.MaxNonOperatingTemperature,
			chassis.MinNonOperatingTemperature,
			chassis.Overload,
			chassis.MaxVibrationSine,
			chassis.MinVibrationSine,
			chassis.MaxVibrationRandom,
			chassis.MinVibrationRandom,
			chassis.Axes,
			chassis.ShockResponseSpectrum,
			chassis.PeakOverloadSpectrum1,
			chassis.PeakOverloadSpectrum2,
			chassis.PeakFrequencySpectrum1,
			chassis.PeakFrequencySpectrum2,
			chassis.Length,
			chassis.Width,
			chassis.Height,
			chassis.Weight,
			chassis.PowerHandlingCapabilityPerBoard,
			chassis.TemperaturePerBoard,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newChassisID int64
	err = r.sqalxConn.GetContext(ctx, &newChassisID, query, params...)
	return newChassisID, err
}

func (r *ChassisRepository) UpdateChassisByID(ctx context.Context, chassis model.Chassis) error {
	query, params, err := postgresql.Builder.Update("chassis").
		Set("name", chassis.Name).
		Set("slots", chassis.Slots).
		Set("size", chassis.Size).
		Set("max_operating_temperature", chassis.MaxOperatingTemperature).
		Set("min_operating_temperature", chassis.MinOperatingTemperature).
		Set("max_non_operating_temperature", chassis.MaxNonOperatingTemperature).
		Set("min_non_operating_temperature", chassis.MinNonOperatingTemperature).
		Set("overload", chassis.Overload).
		Set("max_vibration_sine", chassis.MaxVibrationSine).
		Set("min_vibration_sine", chassis.MinVibrationSine).
		Set("max_vibration_random", chassis.MaxVibrationRandom).
		Set("min_vibration_random", chassis.MinVibrationRandom).
		Set("axes", chassis.Axes).
		Set("shock_response_spectrum", chassis.ShockResponseSpectrum).
		Set("peak_overload_spectrum_1", chassis.PeakOverloadSpectrum1).
		Set("peak_overload_spectrum_2", chassis.PeakOverloadSpectrum2).
		Set("peak_frequency_spectrum_1", chassis.PeakFrequencySpectrum1).
		Set("peak_frequency_spectrum_2", chassis.PeakFrequencySpectrum2).
		Set("length", chassis.Length).
		Set("width", chassis.Width).
		Set("height", chassis.Height).
		Set("weight", chassis.Weight).
		Set("power_handling_capability_per_board", chassis.PowerHandlingCapabilityPerBoard).
		Set("temperature_per_board", chassis.TemperaturePerBoard).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": chassis.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *ChassisRepository) DeleteChassisByID(ctx context.Context, id int64) error {
	query, params, err := postgresql.Builder.Delete("chassis").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *ChassisRepository) GetChassisByID(ctx context.Context, id int64) (model.Chassis, error) {
	var chassis model.Chassis
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"slots",
		"size",
		"max_operating_temperature",
		"min_operating_temperature",
		"max_non_operating_temperature",
		"min_non_operating_temperature",
		"overload",
		"max_vibration_sine",
		"min_vibration_sine",
		"max_vibration_random",
		"min_vibration_random",
		"axes",
		"shock_response_spectrum",
		"peak_overload_spectrum_1",
		"peak_overload_spectrum_2",
		"peak_frequency_spectrum_1",
		"peak_frequency_spectrum_2",
		"length",
		"width",
		"height",
		"weight",
		"power_handling_capability_per_board",
		"temperature_per_board",
		"updated_at",
		"created_at",
	).
		From("chassis").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return chassis, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &chassis, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return chassis, model.ErrChassisNotFound
		}
	}

	return chassis, err
}

func (r *ChassisRepository) GetChassisByFilters(ctx context.Context, filters map[string]interface{}) ([]model.Chassis, error) {
	builder := postgresql.Builder.Select(
		"id",
		"name",
		"slots",
		"size",
		"max_operating_temperature",
		"min_operating_temperature",
		"max_non_operating_temperature",
		"min_non_operating_temperature",
		"overload",
		"max_vibration_sine",
		"min_vibration_sine",
		"max_vibration_random",
		"min_vibration_random",
		"axes",
		"shock_response_spectrum",
		"peak_overload_spectrum_1",
		"peak_overload_spectrum_2",
		"peak_frequency_spectrum_1",
		"peak_frequency_spectrum_2",
		"length",
		"width",
		"height",
		"weight",
		"power_handling_capability_per_board",
		"temperature_per_board",
	).From("chassis")

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var chassisList []model.Chassis
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &chassisList, query, params...); err != nil {
		return nil, err
	}

	return chassisList, nil
}

func (r *ChassisRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterChassisByMaxLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"length": length})
			}
		case model.FilterChassisByMinLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"length": length}) // Условие для минимальной длины
			}
		case model.FilterChassisByMaxWidth:
			if width, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"width": width}) // Условие для максимальной ширины
			}
		case model.FilterChassisByMinWidth:
			if width, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"width": width}) // Условие для минимальной ширины
			}
		case model.FilterChassisByMaxHeight:
			if height, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"height": height}) // Условие для максимальной высоты
			}
		case model.FilterChassisByMinHeight:
			if height, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"height": height}) // Условие для минимальной высоты
			}
		case model.FilterChassisByMaxWeight:
			if weight, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"weight": weight}) // Условие для максимального веса
			}
		case model.FilterChassisByMinWeight:
			if weight, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"weight": weight}) // Условие для минимального веса
			}
		case model.FilterChassisByMaxPowerHandlingCapabilityPerBoard:
			if power, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"power_handling_capability_per_board": power}) // Условие для максимальной мощности
			}
		case model.FilterChassisByMinPowerHandlingCapabilityPerBoard:
			if power, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"power_handling_capability_per_board": power}) // Условие для минимальной мощности
			}
		case model.FilterChassisByMaxTemperaturePerBoard:
			if temperature, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"temperature_per_board": temperature}) // Условие для максимальной температуры
			}
		case model.FilterChassisByMinTemperaturePerBoard:
			if temperature, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"temperature_per_board": temperature}) // Условие для минимальной температуры
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}
