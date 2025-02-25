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

type VHFTransceiverRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IVHFTransceiverRepository {
	return &VHFTransceiverRepository{sqalxConn: sqalxConn}
}

func (r *VHFTransceiverRepository) CreateVHFTransceiver(ctx context.Context, vhfTransceiver model.VHFTransceiver) (int64, error) {
	query, params, err := postgresql.Builder.Insert("vhf_transceiver").
		Columns(
			"name",
			"length",
			"width",
			"height",
			"weight",
			"supply_voltage",
			"power_consumption",
			"interface",
			"operating_frequency",
			"output_power",
			"sensitivity_receiver",
			"max_operating_temperature",
			"min_operating_temperature",
			"mechanical_vibration",
			"mechanical_shock",
		).
		Values(
			vhfTransceiver.Name.String,
			vhfTransceiver.Length.Float64,
			vhfTransceiver.Width.Float64,
			vhfTransceiver.Height.Float64,
			vhfTransceiver.Weight.Float64,
			vhfTransceiver.SupplyVoltage.Float64,
			vhfTransceiver.PowerConsumption.Float64,
			vhfTransceiver.Interface.String,
			vhfTransceiver.OperatingFrequency.Float64,
			vhfTransceiver.OutputPower.Float64,
			vhfTransceiver.SensitivityReceiver.Float64,
			vhfTransceiver.MaxOperatingTemperature.Float64,
			vhfTransceiver.MinOperatingTemperature.Float64,
			vhfTransceiver.MechanicalVibration.Int64,
			vhfTransceiver.MechanicalShock.Int64,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newTransceiverID int64
	err = r.sqalxConn.GetContext(ctx, &newTransceiverID, query, params...)
	return newTransceiverID, err
}

func (r *VHFTransceiverRepository) GetVHFTransceiverByID(ctx context.Context, vhfTransceiverID int64) (model.VHFTransceiver, error) {
	var transceiver model.VHFTransceiver
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
		"operating_frequency",
		"output_power",
		"sensitivity_receiver",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("vhf_transceiver").
		Where(sq.Eq{"id": vhfTransceiverID}).
		ToSql()
	if err != nil {
		return transceiver, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &transceiver, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return transceiver, model.ErrVHFTransceiverNotFound
		}
	}

	return transceiver, err
}

func (r *VHFTransceiverRepository) GetVHFTransceiverByName(ctx context.Context, vhfTransceiverName string) (model.VHFTransceiver, error) {
	var transceiver model.VHFTransceiver
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
		"operating_frequency",
		"output_power",
		"sensitivity_receiver",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).
		From("vhf_transceiver").
		Where(sq.Eq{"name": vhfTransceiverName}).
		ToSql()
	if err != nil {
		return transceiver, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &transceiver, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return transceiver, model.ErrVHFTransceiverNotFound
		}
	}

	return transceiver, err
}

func (r *VHFTransceiverRepository) GetVHFTransceiversByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.VHFTransceiver, error) {
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
		"operating_frequency",
		"output_power",
		"sensitivity_receiver",
		"max_operating_temperature",
		"min_operating_temperature",
		"mechanical_vibration",
		"mechanical_shock",
		"updated_at",
		"created_at",
	).From("vhf_transceiver")

	builder = r.ApplyFilters(builder, filters)

	if sortParams != "" {
		builder = builder.OrderBy(sortParams)
	}

	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var transceivers []model.VHFTransceiver
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &transceivers, query, params...); err != nil {
		return nil, err
	}

	return transceivers, nil
}

func (r *VHFTransceiverRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterVHFTransceiverByMinOperatingFrequency:
			if frequency, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.GtOrEq{"operating_frequency": frequency})
			}
		case model.FilterVHFTransceiverByMaxOperatingFrequency:
			if frequency, ok := filterValue.(float64); ok {
				builder = builder.Where(sq.LtOrEq{"operating_frequency": frequency})
			}
		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
		}
	}

	return builder
}

func (r *VHFTransceiverRepository) UpdateVHFTransceiver(ctx context.Context, vhfTransceiver model.VHFTransceiver) error {
	query, params, err := postgresql.Builder.Update("vhf_transceiver").
		Set("name", vhfTransceiver.Name.String).
		Set("length", vhfTransceiver.Length.Float64).
		Set("width", vhfTransceiver.Width.Float64).
		Set("height", vhfTransceiver.Height.Float64).
		Set("weight", vhfTransceiver.Weight.Float64).
		Set("supply_voltage", vhfTransceiver.SupplyVoltage.Float64).
		Set("power_consumption", vhfTransceiver.PowerConsumption.Float64).
		Set("interface", vhfTransceiver.Interface.String).
		Set("operating_frequency", vhfTransceiver.OperatingFrequency.Float64).
		Set("output_power", vhfTransceiver.OutputPower.Float64).
		Set("sensitivity_receiver", vhfTransceiver.SensitivityReceiver.Float64).
		Set("max_operating_temperature", vhfTransceiver.MaxOperatingTemperature.Float64).
		Set("min_operating_temperature", vhfTransceiver.MinOperatingTemperature.Float64).
		Set("mechanical_vibration", vhfTransceiver.MechanicalVibration.Int64).
		Set("mechanical_shock", vhfTransceiver.MechanicalShock.Int64).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": vhfTransceiver.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *VHFTransceiverRepository) DeleteVHFTransceiver(ctx context.Context, vhfTransceiverID int64) error {
	query, params, err := postgresql.Builder.Delete("vhf_transceiver").
		Where(sq.Eq{"id": vhfTransceiverID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
