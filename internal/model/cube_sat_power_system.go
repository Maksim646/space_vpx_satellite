package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrPowerSystemNotFound = errors.New("power system not found")

	PowerSystemNotFound  = "power system not found"
	PowerSystemsNotFound = "power systems not found"
)

const (
	FilterCubeSatPowerSystemByMinLength = "filter_cube_sat_power_system_min_length"
	FilterCubeSatPowerSystemByMaxLength = "filter_cube_sat_power_system_max_length"

	DefaultCubeSatPowerSystemsSort     = "id desc"
	CubeSatPowerSystemsSortByCreatedAt = "created_at desc"
)

type CubeSatPowerSystem struct {
	ID                             int64           `db:"id"`
	Name                           sql.NullString  `db:"name"`
	Length                         sql.NullFloat64 `db:"length"`
	Width                          sql.NullFloat64 `db:"width"`
	Height                         sql.NullFloat64 `db:"height"`
	Weight                         sql.NullFloat64 `db:"weight"`
	SolarPanelChannels             sql.NullInt64   `db:"solar_panel_channels"`
	SolarPanelsType                sql.NullString  `db:"solar_panels_type"`
	SolarPanelVoltageMin           sql.NullFloat64 `db:"solar_panel_voltage_min"`
	SolarPanelVoltageMax           sql.NullFloat64 `db:"solar_panel_voltage_max"`
	SolarPanelCurrentPerChannelMax sql.NullFloat64 `db:"solar_panel_current_per_channel_max"`
	TotalCurrentOfSolarPanelsMax   sql.NullFloat64 `db:"total_current_of_solar_panels_max"`
	OutputChannels                 sql.NullInt64   `db:"output_channels"`
	SystemBusVoltageSolarPanels    sql.NullFloat64 `db:"system_bus_voltage_solar_panels"`
	SystemBusVoltageOutputChannels sql.NullFloat64 `db:"system_bus_voltage_output_channels"`
	CurrentOutputChannelsMax       sql.NullFloat64 `db:"current_output_channels_max"`
	TotalOutputCurrent             sql.NullFloat64 `db:"total_output_current"`
	DataInterface                  sql.NullString  `db:"data_interface"`
	MaxOperatingTemperature        sql.NullFloat64 `db:"max_operating_temperature"`
	MinOperatingTemperature        sql.NullFloat64 `db:"min_operating_temperature"`
	MechanicalVibration            sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock                sql.NullInt64   `db:"mechanical_shock"`
	UpdatedAt                      sql.NullTime    `db:"updated_at"`
	CreatedAt                      time.Time       `db:"created_at"`
}

type IPowerSystemRepository interface {
	CreatePowerSystem(ctx context.Context, powerSystem CubeSatPowerSystem) (int64, error)

	GetPowerSystemByID(ctx context.Context, powerSystemID int64) (CubeSatPowerSystem, error)
	GetPowerSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatPowerSystem, error)

	UpdatePowerSystem(ctx context.Context, powerSystem CubeSatPowerSystem) error

	DeletePowerSystem(ctx context.Context, powerSystemID int64) error
}

type IPowerSystemUsecase interface {
	CreatePowerSystem(ctx context.Context, powerSystem CubeSatPowerSystem) (int64, error)

	GetPowerSystemByID(ctx context.Context, powerSystemID int64) (CubeSatPowerSystem, error)
	GetPowerSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatPowerSystem, error)

	UpdatePowerSystem(ctx context.Context, powerSystem CubeSatPowerSystem) error

	DeletePowerSystem(ctx context.Context, powerSystemID int64) error
}
