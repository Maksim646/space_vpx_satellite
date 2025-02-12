package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrSolarPanelNotFound = errors.New("solar panel not found")

	SolarPanelNotFound  = "solar panel not found"
	SolarPanelsNotFound = "solar panels not found"
)

const (
	FilterCubeSatSolarPanelSideByMinLength = "filter_cube_sat_solar_panel_side_min_length"
	FilterCubeSatSolarPanelSideByMaxLength = "filter_cube_sat_solar_panel_side_max_length"

	DefaultCubeSatSolarPanelsSideSort     = "id desc"
	CubeSatSolarPanelsSideSortByCreatedAt = "created_at desc"
)

type CubeSatSolarPanelSide struct {
	ID int64 `db:"id"`

	Name                    sql.NullString  `db:"name"`
	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullFloat64 `db:"weight"`
	Interface               sql.NullString  `db:"interface"`
	Voc                     sql.NullFloat64 `db:"voc"`
	Isc                     sql.NullFloat64 `db:"isc"`
	Vmp                     sql.NullFloat64 `db:"vmp"`
	Imp                     sql.NullFloat64 `db:"imp"`
	Efficiency              sql.NullInt64   `db:"efficiency"`
	CoilArea                sql.NullFloat64 `db:"coil_area"`
	CoilResistance          sql.NullInt64   `db:"coil_resistance"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MechanicalVibration     sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock         sql.NullInt64   `db:"mechanical_shock"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type ISolarPanelRepository interface {
	CreateSolarPanelSide(ctx context.Context, solarPanel CubeSatSolarPanelSide) (int64, error)

	GetSolarPanelSideByID(ctx context.Context, solarPanelSideID int64) (CubeSatSolarPanelSide, error)
	GetSolarPanelSideByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatSolarPanelSide, error)

	UpdateSolarPanelSide(ctx context.Context, solarPanel CubeSatSolarPanelSide) error

	DeleteSolarPanelSide(ctx context.Context, solarPanelSideID int64) error
}

type ISolarPanelUsecase interface {
	CreateSolarPanelSide(ctx context.Context, solarPanel CubeSatSolarPanelSide) (int64, error)

	GetSolarPanelSideByID(ctx context.Context, solarPanelSideID int64) (CubeSatSolarPanelSide, error)
	GetSolarPanelSideByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatSolarPanelSide, error)

	UpdateSolarPanelSide(ctx context.Context, solarPanel CubeSatSolarPanelSide) error

	DeleteSolarPanelSide(ctx context.Context, solarPanelSideID int64) error
}
