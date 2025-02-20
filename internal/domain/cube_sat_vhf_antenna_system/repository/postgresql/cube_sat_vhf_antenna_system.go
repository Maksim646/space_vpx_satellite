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

// VHFAntennaSystemRepository implements the IVHFAntennaSystemRepository interface.
type VHFAntennaSystemRepository struct {
	sqalxConn sqalx.Node
}

// New creates a new VHFAntennaSystemRepository.
func New(sqalxConn sqalx.Node) model.IVHFAntennaSystemRepository {
	return &VHFAntennaSystemRepository{sqalxConn: sqalxConn}
}

// CreateVHFAntennaSystem creates a new VHF antenna system in the database.
func (r *VHFAntennaSystemRepository) CreateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem model.VHFAntennaSystem) (int64, error) {
	query, params, err := postgresql.Builder.Insert("vhf_antenna_system").
		Columns(
			"name",
			"length",
			"width",
			"height",
			"weight",
			"interface",
			"frequency",
			"max_operating_temperature",
			"min_operating_temperature",
			"mechanical_vibration",
			"mechanical_shock",
		).
		Values(
			vhfAntennaSystem.Name.String,
			vhfAntennaSystem.Length.Float64,
			vhfAntennaSystem.Width.Float64,
			vhfAntennaSystem.Height.Float64,
			vhfAntennaSystem.Weight.Float64,
			vhfAntennaSystem.Interface.String,
			vhfAntennaSystem.Frequency.Float64,
			vhfAntennaSystem.MaxOperatingTemperature.Float64,
			vhfAntennaSystem.MinOperatingTemperature.Float64,
			vhfAntennaSystem.MechanicalVibration.Int64,
			vhfAntennaSystem.MechanicalShock.Int64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newAntennaSystemID int64
	err = r.sqalxConn.GetContext(ctx, &newAntennaSystemID, query, params...)
	return newAntennaSystemID, err
}

// GetVHFAntennaSystemByID retrieves a VHF antenna system from the database by its ID.
func (r *VHFAntennaSystemRepository) GetVHFAntennaSystemByID(ctx context.Context, vhfAntennaSystemID int64) (model.VHFAntennaSystem, error) {
	var antennaSystem model.VHFAntennaSystem
	query, params, err := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"interface",
		"frequency",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("vhf_antenna_system").
		Where(sq.Eq{"id": vhfAntennaSystemID}).
		ToSql()
	if err != nil {
		return antennaSystem, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &antennaSystem, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return antennaSystem, model.ErrVHFAntennaSystemNotFound
		}
	}

	return antennaSystem, err
}

// GetVHFAntennaSystemsByFilters retrieves VHF antenna systems from the database based on filters.
func (r *VHFAntennaSystemRepository) GetVHFAntennaSystemsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.VHFAntennaSystem, error) {
	builder := postgresql.Builder.Select(
		"id",
		"name",
		"length",
		"width",
		"height",
		"weight",
		"interface",
		"frequency",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("vhf_antenna_system")

	builder = r.ApplyFilters(builder, filters)

	if sortParams != "" {
		builder = builder.OrderBy(sortParams)
	}

	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var antennaSystems []model.VHFAntennaSystem
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &antennaSystems, query, params...); err != nil {
		return nil, err
	}

	return antennaSystems, nil
}

// ApplyFilters applies filters to the SQL query builder.
func (r *VHFAntennaSystemRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterVHFAntennaSystemByMinFrequency:
			if frequency, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"frequency": frequency})
			}
		case model.FilterVHFAntennaSystemByMaxFrequency:
			if frequency, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"frequency": frequency})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

// UpdateVHFAntennaSystem updates an existing VHF antenna system in the database.
func (r *VHFAntennaSystemRepository) UpdateVHFAntennaSystem(ctx context.Context, vhfAntennaSystem model.VHFAntennaSystem) error {
	query, params, err := postgresql.Builder.Update("vhf_antenna_system").
		Set("name", vhfAntennaSystem.Name.String).
		Set("length", vhfAntennaSystem.Length.Float64).
		Set("width", vhfAntennaSystem.Width.Float64).
		Set("height", vhfAntennaSystem.Height.Float64).
		Set("weight", vhfAntennaSystem.Weight.Float64).
		Set("interface", vhfAntennaSystem.Interface.String).
		Set("frequency", vhfAntennaSystem.Frequency.Float64).
		Set("max_operating_temperature", vhfAntennaSystem.MaxOperatingTemperature.Float64).
		Set("min_operating_temperature", vhfAntennaSystem.MinOperatingTemperature.Float64).
		Set("mechanical_vibration", vhfAntennaSystem.MechanicalVibration.Int64).
		Set("mechanical_shock", vhfAntennaSystem.MechanicalShock.Int64).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": vhfAntennaSystem.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

// DeleteVHFAntennaSystem deletes a VHF antenna system from the database by its ID.
func (r *VHFAntennaSystemRepository) DeleteVHFAntennaSystem(ctx context.Context, vhfAntennaSystemID int64) error {
	query, params, err := postgresql.Builder.Delete("vhf_antenna_system").
		Where(sq.Eq{"id": vhfAntennaSystemID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
