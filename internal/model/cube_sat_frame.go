package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrCubeSatFrameNotFound = errors.New("cube sat frame not found")

	CubeSatFrameNotFound = "cube sat frame not found"
)

type CubeSatFrame struct {
	ID int64 `db:"id"`

	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullInt64   `db:"weight"`
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
	// GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatFrame, error)

	// DeleteCubeSatFrame(ctx context.Context, id int64) error
}

type ICubeSatFrameUsecase interface {
	CreateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) (int64, error)

	UpdateCubeSatFrame(ctx context.Context, cubeSatFrame CubeSatFrame) error

	GetCubeSatFrameByID(ctx context.Context, id int64) (CubeSatFrame, error)
	// 	GetCubeSatFramesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]CubeSatFrame, error)

	// 	DeleteCubeSatFrame(ctx context.Context, id int64) error
}
