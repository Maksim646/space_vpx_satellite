package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrSolarPanelTopNotFound = errors.New("solar panel top not found")

	SolarPanelTopNotFound  = "solar panel top not found"
	SolarPanelsTopNotFound = "solar panels top not found"
)

const (
	FilterSolarPanelTopByMinLength = "filter_solar_panel_top_min_length"
	FilterSolarPanelTopByMaxLength = "filter_solar_panel_top_max_length"

	DefaultSolarPanelTopSort     = "id desc"
	SolarPanelTopSortByCreatedAt = "created_at desc"
)

type SolarPanelTop struct { // Renamed struct
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
	Efficiency              sql.NullFloat64 `db:"efficiency"`
	CoilArea                sql.NullFloat64 `db:"coil_area"`
	CoilResistance          sql.NullFloat64 `db:"coil_resistance"`
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MechanicalVibration     sql.NullFloat64 `db:"mechanical_vibration"`
	MechanicalShock         sql.NullFloat64 `db:"mechanical_shock"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type ISolarPanelTopRepository interface {
	CreateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) (int64, error)

	GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (SolarPanelTop, error)
	GetSolarPanelTopByName(ctx context.Context, solarPanelTopName string) (SolarPanelTop, error)
	GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]SolarPanelTop, error)

	UpdateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) error

	DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error
}

type ISolarPanelTopUsecase interface {
	CreateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) (int64, error)

	GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (SolarPanelTop, error)
	GetSolarPanelTopByName(ctx context.Context, solarPanelTopName string) (SolarPanelTop, error)
	GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]SolarPanelTop, error)

	UpdateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) error

	DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error
}
