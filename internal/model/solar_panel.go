package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrSolarPanelNotFound = errors.New("solar panel not found")

	SolarPanelNotFound = "solar panel not found"
)

type SolarPanel struct {
	ID                      int64           `db:"id"`
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
	CreateSolarPanel(ctx context.Context, solarPanel SolarPanel) (int64, error)
	GetSolarPanelByID(ctx context.Context, id int64) (SolarPanel, error)
}

type ISolarPanelUsecase interface {
	CreateSolarPanel(ctx context.Context, solarPanel SolarPanel) (int64, error)
	GetSolarPanelByID(ctx context.Context, id int64) (SolarPanel, error)
}
