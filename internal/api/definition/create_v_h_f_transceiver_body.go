// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateVHFTransceiverBody create v h f transceiver body
//
// swagger:model CreateVHFTransceiverBody
type CreateVHFTransceiverBody struct {

	// height
	Height float64 `json:"height,omitempty"`

	// interface
	Interface string `json:"interface,omitempty"`

	// length
	Length float64 `json:"length,omitempty"`

	// max operating temperature
	MaxOperatingTemperature float64 `json:"max_operating_temperature,omitempty"`

	// mechanical shock
	MechanicalShock int64 `json:"mechanical_shock,omitempty"`

	// mechanical vibration
	MechanicalVibration int64 `json:"mechanical_vibration,omitempty"`

	// min operating temperature
	MinOperatingTemperature float64 `json:"min_operating_temperature,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operating frequency
	OperatingFrequency float64 `json:"operating_frequency,omitempty"`

	// output power
	OutputPower float64 `json:"output_power,omitempty"`

	// power consumption
	PowerConsumption float64 `json:"power_consumption,omitempty"`

	// sensitivity receiver
	SensitivityReceiver float64 `json:"sensitivity_receiver,omitempty"`

	// supply voltage
	SupplyVoltage float64 `json:"supply_voltage,omitempty"`

	// weight
	Weight float64 `json:"weight,omitempty"`

	// width
	Width float64 `json:"width,omitempty"`
}

// Validate validates this create v h f transceiver body
func (m *CreateVHFTransceiverBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create v h f transceiver body based on context it is used
func (m *CreateVHFTransceiverBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateVHFTransceiverBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVHFTransceiverBody) UnmarshalBinary(b []byte) error {
	var res CreateVHFTransceiverBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
