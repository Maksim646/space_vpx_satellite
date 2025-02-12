package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrSolarPanelTopNotFound = errors.New("solar panel top not found") // Changed error message

	SolarPanelTopNotFound  = "solar panel top not found"  // Changed const name
	SolarPanelsTopNotFound = "solar panels top not found" // Changed const name
)

const (
	FilterSolarPanelTopByMinLength = "filter_solar_panel_top_min_length" // Changed const name
	FilterSolarPanelTopByMaxLength = "filter_solar_panel_top_max_length" // Changed const name

	DefaultSolarPanelTopSort     = "id desc"         // Changed const name
	SolarPanelTopSortByCreatedAt = "created_at desc" // Changed const name
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
	Efficiency              sql.NullFloat64 `db:"efficiency"`      // Changed to NullFloat64
	CoilArea                sql.NullFloat64 `db:"coil_area"`       // Renamed from CoilArea and more descriptive
	CoilResistance          sql.NullFloat64 `db:"coil_resistance"` // Changed to NullFloat64
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MechanicalVibration     sql.NullFloat64 `db:"mechanical_vibration"` // Changed to NullFloat64
	MechanicalShock         sql.NullFloat64 `db:"mechanical_shock"`     // Changed to NullFloat64
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type ISolarPanelTopRepository interface { // Renamed interface
	CreateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) (int64, error) // Changed struct

	GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (SolarPanelTop, error)                                                               // Changed struct
	GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]SolarPanelTop, error) // Changed struct

	UpdateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) error // Changed struct

	DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error
}

type ISolarPanelTopUsecase interface { // Renamed interface
	CreateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) (int64, error) // Changed struct

	GetSolarPanelTopByID(ctx context.Context, solarPanelTopID int64) (SolarPanelTop, error)                                                               // Changed struct
	GetSolarPanelTopByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]SolarPanelTop, error) // Changed struct

	UpdateSolarPanelTop(ctx context.Context, solarPanel SolarPanelTop) error // Changed struct

	DeleteSolarPanelTop(ctx context.Context, solarPanelTopID int64) error
}
