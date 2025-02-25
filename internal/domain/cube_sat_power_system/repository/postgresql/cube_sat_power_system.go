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

type PowerSystemRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IPowerSystemRepository {
	return &PowerSystemRepository{sqalxConn: sqalxConn}
}

func (r *PowerSystemRepository) CreatePowerSystem(ctx context.Context, powerSystem model.CubeSatPowerSystem) (int64, error) {
	query, params, err := postgresql.Builder.Insert("cube_sat_power_system").
		Columns(
			"name",
			"length",
			"width",
			"height",
			"weight",
			"solar_panel_channels",
			"solar_panels_type",
			"solar_panel_voltage_min",
			"solar_panel_voltage_max",
			"solar_panel_current_per_channel_max",
			"total_current_of_solar_panels_max",
			"output_channels",
			"system_bus_voltage_solar_panels",
			"system_bus_voltage_output_channels",
			"current_output_channels_max",
			"total_output_current",
			"data_interface",
			"max_operating_temperature",
			"min_operating_temperature",
			"mechanical_vibration",
			"mechanical_shock",
		).
		Values(
			powerSystem.Name.String,
			powerSystem.Length.Float64,
			powerSystem.Width.Float64,
			powerSystem.Height.Float64,
			powerSystem.Weight.Float64,
			powerSystem.SolarPanelChannels.Int64,
			powerSystem.SolarPanelsType.String,
			powerSystem.SolarPanelVoltageMin.Float64,
			powerSystem.SolarPanelVoltageMax.Float64,
			powerSystem.SolarPanelCurrentPerChannelMax.Float64,
			powerSystem.TotalCurrentOfSolarPanelsMax.Float64,
			powerSystem.OutputChannels.Int64,
			powerSystem.SystemBusVoltageSolarPanels.Float64,
			powerSystem.SystemBusVoltageOutputChannels.Float64,
			powerSystem.CurrentOutputChannelsMax.Float64,
			powerSystem.TotalOutputCurrent.Float64,
			powerSystem.DataInterface.String,
			powerSystem.MaxOperatingTemperature.Float64,
			powerSystem.MinOperatingTemperature.Float64,
			powerSystem.MechanicalVibration.Int64,
			powerSystem.MechanicalShock.Int64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newPowerSystemID int64
	err = r.sqalxConn.GetContext(ctx, &newPowerSystemID, query, params...)
	return newPowerSystemID, err
}

func (r *PowerSystemRepository) GetPowerSystemByID(ctx context.Context, powerSystemID int64) (model.CubeSatPowerSystem, error) {
	var powerSystem model.CubeSatPowerSystem
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"solar_panel_channels",
		"solar_panels_type",
		"solar_panel_voltage_min",
		"solar_panel_voltage_max",
		"solar_panel_current_per_channel_max",
		"total_current_of_solar_panels_max",
		"output_channels",
		"system_bus_voltage_solar_panels",
		"system_bus_voltage_output_channels",
		"current_output_channels_max",
		"total_output_current",
		"data_interface",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("cube_sat_power_system").
		Where(sq.Eq{"id": powerSystemID}).
		ToSql()
	if err != nil {
		return powerSystem, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &powerSystem, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return powerSystem, model.ErrPowerSystemNotFound
		}
	}

	return powerSystem, err
}

func (r *PowerSystemRepository) GetPowerSystemByName(ctx context.Context, powerSystemName string) (model.CubeSatPowerSystem, error) {
	var powerSystem model.CubeSatPowerSystem
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"solar_panel_channels",
		"solar_panels_type",
		"solar_panel_voltage_min",
		"solar_panel_voltage_max",
		"solar_panel_current_per_channel_max",
		"total_current_of_solar_panels_max",
		"output_channels",
		"system_bus_voltage_solar_panels",
		"system_bus_voltage_output_channels",
		"current_output_channels_max",
		"total_output_current",
		"data_interface",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("cube_sat_power_system").
		Where(sq.Eq{"name": powerSystemName}).
		ToSql()
	if err != nil {
		return powerSystem, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &powerSystem, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return powerSystem, model.ErrPowerSystemNotFound
		}
	}

	return powerSystem, err
}

func (r *PowerSystemRepository) GetPowerSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.CubeSatPowerSystem, error) {
	builder := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"solar_panel_channels",
		"solar_panels_type",
		"solar_panel_voltage_min",
		"solar_panel_voltage_max",
		"solar_panel_current_per_channel_max",
		"total_current_of_solar_panels_max",
		"output_channels",
		"system_bus_voltage_solar_panels",
		"system_bus_voltage_output_channels",
		"current_output_channels_max",
		"total_output_current",
		"data_interface",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("cube_sat_power_system")

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var powerSystems []model.CubeSatPowerSystem
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &powerSystems, query, params...); err != nil {
		return nil, err
	}

	return powerSystems, nil
}

func (r *PowerSystemRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterCubeSatPowerSystemByMaxLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"length": length})
			}
		case model.FilterCubeSatPowerSystemByMinLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"length": length})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *PowerSystemRepository) UpdatePowerSystem(ctx context.Context, powerSystem model.CubeSatPowerSystem) error {
	query, params, err := postgresql.Builder.Update("cube_sat_power_system").
		Set("name", powerSystem.Name.String).
		Set("length", powerSystem.Length.Float64).
		Set("width", powerSystem.Width.Float64).
		Set("height", powerSystem.Height.Float64).
		Set("weight", powerSystem.Weight.Float64).
		Set("solar_panel_channels", powerSystem.SolarPanelChannels.Int64).
		Set("solar_panels_type", powerSystem.SolarPanelsType.String).
		Set("solar_panel_voltage_min", powerSystem.SolarPanelVoltageMin.Float64).
		Set("solar_panel_voltage_max", powerSystem.SolarPanelVoltageMax.Float64).
		Set("solar_panel_current_per_channel_max", powerSystem.SolarPanelCurrentPerChannelMax.Float64).
		Set("total_current_of_solar_panels_max", powerSystem.TotalCurrentOfSolarPanelsMax.Float64).
		Set("output_channels", powerSystem.OutputChannels.Int64).
		Set("system_bus_voltage_solar_panels", powerSystem.SystemBusVoltageSolarPanels.Float64).
		Set("system_bus_voltage_output_channels", powerSystem.SystemBusVoltageOutputChannels.Float64).
		Set("current_output_channels_max", powerSystem.CurrentOutputChannelsMax.Float64).
		Set("total_output_current", powerSystem.TotalOutputCurrent.Float64).
		Set("data_interface", powerSystem.DataInterface.String).
		Set("max_operating_temperature", powerSystem.MaxOperatingTemperature.Float64).
		Set("min_operating_temperature", powerSystem.MinOperatingTemperature.Float64).
		Set("mechanical_vibration", powerSystem.MechanicalVibration.Int64).
		Set("mechanical_shock", powerSystem.MechanicalShock.Int64).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": powerSystem.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *PowerSystemRepository) DeletePowerSystem(ctx context.Context, powerSystemID int64) error {
	query, params, err := postgresql.Builder.Delete("cube_sat_power_system").
		Where(sq.Eq{"cube_sat_power_system.id": powerSystemID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
