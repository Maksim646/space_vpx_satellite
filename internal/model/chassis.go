package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrChassisNotFound = errors.New("chassis not found")

	ChassisNotFound = "chassis not found"
)

const (
	FilterChassisByLength                          = "filter_chassis_by_length"
	FilterChassisByWidth                           = "filter_chassis_by_width"
	FilterChassisByHeight                          = "filter_chassis_by_height"
	FilterChassisByWeight                          = "filter_chassis_by_weight"
	FilterChassisByPowerHandlingCapabilityPerBoard = "filter_chassis_by_power_handling_capability_per_board"
	FilterChassisByTemperaturePerBoard             = "filter_chassis_by_temperature_per_board"
)

type Chassis struct {
	ID                              int64        `db:"id"`
	Name                            string       `db:"name"`
	Slots                           int64        `db:"slots"`
	Size                            string       `db:"size"`
	MaxOperatingTemperature         float64      `db:"max_operating_temperature"`
	MinOperatingTemperature         float64      `db:"min_operating_temperature"`
	MaxNonOperatingTemperature      float64      `db:"max_non_operating_temperature"`
	MinNonOperatingTemperature      float64      `db:"min_non_operating_temperature"`
	Overload                        float64      `db:"overload"`
	MaxVibrationSine                float64      `db:"max_vibration_sine"`
	MinVibrationSine                float64      `db:"min_vibration_sine"`
	MaxVibrationRandom              float64      `db:"max_vibration_random"`
	MinVibrationRandom              float64      `db:"min_vibration_random"`
	Axes                            int64        `db:"axes"`
	ShockResponseSpectrum           float64      `db:"shock_response_spectrum"`
	PeakOverloadSpectrum1           float64      `db:"peak_overload_spectrum_1"`
	PeakOverloadSpectrum2           float64      `db:"peak_overload_spectrum_2"`
	PeakFrequencySpectrum1          float64      `db:"peak_frequency_spectrum_1"`
	PeakFrequencySpectrum2          float64      `db:"peak_frequency_spectrum_2"`
	Length                          float64      `db:"length"`
	Width                           float64      `db:"width"`
	Height                          float64      `db:"height"`
	Weight                          float64      `db:"weight"`
	PowerHandlingCapabilityPerBoard float64      `db:"power_handling_capability_per_board"`
	TemperaturePerBoard             float64      `db:"temperature_per_board"`
	UpdatedAt                       sql.NullTime `db:"updated_at"`
	CreatedAt                       time.Time    `db:"created_at"`
}

type IChassisRepository interface {
	CreateChassisByID(ctx context.Context, chassis Chassis) (int64, error)

	UpdateChassisByID(ctx context.Context, chassis Chassis) error

	DeleteChassisByID(ctx context.Context, id int64) error

	GetChassisByID(ctx context.Context, id int64) (Chassis, error)
	GetChassisByFilters(ctx context.Context, filters map[string]interface{}) ([]Chassis, error)
}

type IChassisUsecase interface {
	CreateChassisByID(ctx context.Context, chassis Chassis) (int64, error)

	UpdateChassisByID(ctx context.Context, chassis Chassis) error

	DeleteChassisByID(ctx context.Context, id int64) error

	GetChassisByID(ctx context.Context, id int64) (Chassis, error)
	GetChassisByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]Chassis, error)
}
