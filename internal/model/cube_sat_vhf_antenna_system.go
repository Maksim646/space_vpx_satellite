package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrVHFAntennaSystemNotFound = errors.New("vhf antenna system not found")

	VHFAntennaSystemNotFound  = "vhf antenna system not found"
	VHFAntennaSystemsNotFound = "vhf antenna systems not found"
)

const (
	FilterVHFAntennaSystemByMinFrequency = "filter_vhf_antenna_system_min_frequency"
	FilterVHFAntennaSystemByMaxFrequency = "filter_vhf_antenna_system_max_frequency"

	DefaultVHFAntennaSystemsSort     = "id desc"
	VHFAntennaSystemsSortByCreatedAt = "created_at desc"
)

type VHFAntennaSystem struct {
	ID                      int64           `db:"id"`
	Name                    sql.NullString  `db:"name"`
	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullFloat64 `db:"weight"`
	Interface               sql.NullString  `db:"interface"`
	Frequency               sql.NullFloat64 `db:"frequency"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MechanicalVibration     sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock         sql.NullInt64   `db:"mechanical_shock"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type IVHFAntennaSystemRepository interface {
	CreateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem VHFAntennaSystem) (int64, error)

	GetVHFAntennaSystemByID(ctx context.Context, vhfAntennaSystemID int64) (VHFAntennaSystem, error)
	GetVHFAntennaSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]VHFAntennaSystem, error)

	UpdateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem VHFAntennaSystem) error

	DeleteVHFAntennaSystem(ctx context.Context, vhfAntennaSystemID int64) error
}

type IVHFAntennaSystemUsecase interface {
	CreateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem VHFAntennaSystem) (int64, error)

	GetVHFAntennaSystemByID(ctx context.Context, vhfAntennaSystemID int64) (VHFAntennaSystem, error)
	GetVHFAntennaSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]VHFAntennaSystem, error)

	UpdateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem VHFAntennaSystem) error

	DeleteVHFAntennaSystem(ctx context.Context, vhfAntennaSystemID int64) error
}
