package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrBoardComputingModuleNotFound = errors.New("board computing module not found")

	BoardComputingModuleNotFound  = "board computing module not found"
	BoardComputingModulesNotFound = "board computing modules not found"
)

const (
	FilterBoardComputingModuleByMinLength = "filter_board_computing_module_min_length"
	FilterBoardComputingModuleByMaxLength = "filter_board_computing_module_max_length"

	DefaultBoardComputingModulesSort     = "id desc"
	BoardComputingModulesSortByCreatedAt = "created_at desc"
)

type BoardComputingModule struct {
	ID                      int64           `db:"id"`
	Name                    sql.NullString  `db:"name"`
	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullFloat64 `db:"weight"`
	SupplyVoltage           sql.NullFloat64 `db:"supply_voltage"`
	PowerConsumption        sql.NullFloat64 `db:"power_consumption"`
	Interface               sql.NullString  `db:"interface"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MechanicalVibration     sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock         sql.NullInt64   `db:"mechanical_shock"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type IBoardComputingModuleRepository interface {
	CreateBoardComputingModule(ctx context.Context, module BoardComputingModule) (int64, error)

	GetBoardComputingModuleByID(ctx context.Context, moduleID int64) (BoardComputingModule, error)
	GetBoardComputingModulesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]BoardComputingModule, error)

	UpdateBoardComputingModule(ctx context.Context, module BoardComputingModule) error

	DeleteBoardComputingModule(ctx context.Context, moduleID int64) error
}

type IBoardComputingModuleUsecase interface {
	CreateBoardComputingModule(ctx context.Context, module BoardComputingModule) (int64, error)

	GetBoardComputingModuleByID(ctx context.Context, moduleID int64) (BoardComputingModule, error)
	GetBoardComputingModulesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]BoardComputingModule, error)

	UpdateBoardComputingModule(ctx context.Context, module BoardComputingModule) error

	DeleteBoardComputingModule(ctx context.Context, moduleID int64) error
}
