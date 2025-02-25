package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrVHFTransceiverNotFound = errors.New("vhf transceiver not found")

	VHFTransceiverNotFound  = "vhf transceiver not found"
	VHFTransceiversNotFound = "vhf transceivers not found"
)

const (
	FilterVHFTransceiverByMinOperatingFrequency = "filter_vhf_transceiver_min_operating_frequency"
	FilterVHFTransceiverByMaxOperatingFrequency = "filter_vhf_transceiver_max_operating_frequency"

	DefaultVHFTransceiversSort     = "id desc"
	VHFTransceiversSortByCreatedAt = "created_at desc"
)

type VHFTransceiver struct {
	ID                      int64           `db:"id"`
	Name                    sql.NullString  `db:"name"`
	Length                  sql.NullFloat64 `db:"length"`
	Width                   sql.NullFloat64 `db:"width"`
	Height                  sql.NullFloat64 `db:"height"`
	Weight                  sql.NullFloat64 `db:"weight"`
	SupplyVoltage           sql.NullFloat64 `db:"supply_voltage"`
	PowerConsumption        sql.NullFloat64 `db:"power_consumption"`
	Interface               sql.NullString  `db:"interface"`
	OperatingFrequency      sql.NullFloat64 `db:"operating_frequency"`
	OutputPower             sql.NullFloat64 `db:"output_power"`
	SensitivityReceiver     sql.NullFloat64 `db:"sensitivity_receiver"`
	MaxOperatingTemperature sql.NullFloat64 `db:"max_operating_temperature"`
	MinOperatingTemperature sql.NullFloat64 `db:"min_operating_temperature"`
	MechanicalVibration     sql.NullInt64   `db:"mechanical_vibration"`
	MechanicalShock         sql.NullInt64   `db:"mechanical_shock"`
	UpdatedAt               sql.NullTime    `db:"updated_at"`
	CreatedAt               time.Time       `db:"created_at"`
}

type IVHFTransceiverRepository interface {
	CreateVHFTransceiver(ctx context.Context, vhfTransceiver VHFTransceiver) (int64, error)

	GetVHFTransceiverByID(ctx context.Context, vhfTransceiverID int64) (VHFTransceiver, error)
	GetVHFTransceiverByName(ctx context.Context, vhfTransceiverName string) (VHFTransceiver, error)
	GetVHFTransceiversByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]VHFTransceiver, error)

	UpdateVHFTransceiver(ctx context.Context, vhfTransceiver VHFTransceiver) error

	DeleteVHFTransceiver(ctx context.Context, vhfTransceiverID int64) error
}

type IVHFTransceiverUsecase interface {
	CreateVHFTransceiver(ctx context.Context, vhfTransceiver VHFTransceiver) (int64, error)

	GetVHFTransceiverByID(ctx context.Context, vhfTransceiverID int64) (VHFTransceiver, error)
	GetVHFTransceiverByName(ctx context.Context, vhfTransceiverName string) (VHFTransceiver, error)
	GetVHFTransceiversByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]VHFTransceiver, error)

	UpdateVHFTransceiver(ctx context.Context, vhfTransceiver VHFTransceiver) error

	DeleteVHFTransceiver(ctx context.Context, vhfTransceiverID int64) error
}
