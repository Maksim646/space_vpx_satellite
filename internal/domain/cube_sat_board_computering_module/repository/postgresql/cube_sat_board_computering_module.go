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

// BoardComputingModuleRepository implements the IBoardComputingModuleRepository interface.
type BoardComputingModuleRepository struct {
	sqalxConn sqalx.Node
}

// New creates a new BoardComputingModuleRepository.
func New(sqalxConn sqalx.Node) model.IBoardComputingModuleRepository {
	return &BoardComputingModuleRepository{sqalxConn: sqalxConn}
}

// CreateBoardComputingModule creates a new board computing module in the database.
func (r *BoardComputingModuleRepository) CreateBoardComputingModule(ctx context.Context, module model.BoardComputingModule) (int64, error) {
	query, params, err := postgresql.Builder.Insert("board_computing_module").
		Columns(
			"name",
			"length",
			"width",
			"height",
			"weight",
			"supply_voltage",
			"power_consumption",
			"interface",
			"max_operating_temperature",
			"min_operating_temperature",
			"mechanical_vibration",
			"mechanical_shock",
		).
		Values(
			module.Name.String,
			module.Length.Float64,
			module.Width.Float64,
			module.Height.Float64,
			module.Weight.Float64,
			module.SupplyVoltage.Float64,
			module.PowerConsumption.Float64,
			module.Interface.String,
			module.MaxOperatingTemperature.Float64,
			module.MinOperatingTemperature.Float64,
			module.MechanicalVibration.Int64,
			module.MechanicalShock.Int64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newModuleID int64
	err = r.sqalxConn.GetContext(ctx, &newModuleID, query, params...)
	return newModuleID, err
}

// GetBoardComputingModuleByID retrieves a board computing module from the database by its ID.
func (r *BoardComputingModuleRepository) GetBoardComputingModuleByID(ctx context.Context, moduleID int64) (model.BoardComputingModule, error) {
	var module model.BoardComputingModule
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"supply_voltage",
		"power_consumption",
		"interface",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("board_computing_module").
		Where(sq.Eq{"id": moduleID}).
		ToSql()
	if err != nil {
		return module, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &module, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return module, model.ErrBoardComputingModuleNotFound
		}
	}

	return module, err
}

// GetBoardComputingModulesByFilters retrieves board computing modules from the database based on filters.
func (r *BoardComputingModuleRepository) GetBoardComputingModulesByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.BoardComputingModule, error) {
	builder := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"supply_voltage",
		"power_consumption",
		"interface",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("board_computing_module")

	builder = r.ApplyFilters(builder, filters)

	if sortParams != "" {
		builder = builder.OrderBy(sortParams)
	}

	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var modules []model.BoardComputingModule
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &modules, query, params...); err != nil {
		return nil, err
	}

	return modules, nil
}

// ApplyFilters applies filters to the SQL query builder.
func (r *BoardComputingModuleRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterBoardComputingModuleByMaxLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"length": length})
			}
		case model.FilterBoardComputingModuleByMinLength:
			if length, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"length": length})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

// UpdateBoardComputingModule updates an existing board computing module in the database.
func (r *BoardComputingModuleRepository) UpdateBoardComputingModule(ctx context.Context, module model.BoardComputingModule) error {
	query, params, err := postgresql.Builder.Update("board_computing_module").
		Set("name", module.Name.String).
		Set("length", module.Length.Float64).
		Set("width", module.Width.Float64).
		Set("height", module.Height.Float64).
		Set("weight", module.Weight.Float64).
		Set("supply_voltage", module.SupplyVoltage.Float64).
		Set("power_consumption", module.PowerConsumption.Float64).
		Set("interface", module.Interface.String).
		Set("max_operating_temperature", module.MaxOperatingTemperature.Float64).
		Set("min_operating_temperature", module.MinOperatingTemperature.Float64).
		Set("mechanical_vibration", module.MechanicalVibration.Int64).
		Set("mechanical_shock", module.MechanicalShock.Int64).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": module.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

// DeleteBoardComputingModule deletes a board computing module from the database by its ID.
func (r *BoardComputingModuleRepository) DeleteBoardComputingModule(ctx context.Context, moduleID int64) error {
	query, params, err := postgresql.Builder.Delete("board_computing_module").
		Where(sq.Eq{"id": moduleID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
