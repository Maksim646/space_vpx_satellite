package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrCubeSatFrameNotFound = errors.New("cube sat frame not found")

	CubeSatFrameNotFound  = "cube sat frame not found"
	CubeSatFramesNotFound = "cube sat frames not found"
)

const (
	FilterCubeSatFrameByMinLength = "filter_cube_sat_frame_min_length"
	FilterCubeSatFrameByMaxLength = "filter_cube_sat_frame_max_length"

	DefaultCubeSatFramesSort     = "id desc"
	CubeSatFramesSortByCreatedAt = "created_at desc"
)

type CubeSatFrame struct {
	ID int64 `db:"id"`

	Name                    sql.NullString  `db:"name"`
	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullFloat64 `db:"weight"`
	Size                    int64           `db:"size"`
	Interface               sql.NullString  `db:"interface"`
	OperatingTemperatureMin sql.NullInt64   `db:"operating_temperature_min"`
	OperatingTemperatureMax sql.NullInt64   `db:"operating_temperature_max"`
	MechanicalVibration     sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock         sql.NullInt64   `db:"mechanical_shock"`
	Link                    sql.NullString  `db:"link"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type ICubeSatFrameRepository interface {
	CreateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) (int64, error)

	UpdateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) error

	GetCubeSatFrameByID(ctx context.Context, id int64) (CubeSatFrame, error)
	GetCubeSatFrameByName(ctx context.Context, frameName string) (CubeSatFrame, error)
	GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatFrame, error)

	DeleteCubeSatFrame(ctx context.Context, id int64) error
}

type ICubeSatFrameUsecase interface {
	CreateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) (int64, error)

	UpdateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) error

	GetCubeSatFrameByID(ctx context.Context, id int64) (CubeSatFrame, error)
	GetCubeSatFrameByName(ctx context.Context, frameName string) (CubeSatFrame, error)
	GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatFrame, error)

	DeleteCubeSatFrame(ctx context.Context, id int64) error
}
