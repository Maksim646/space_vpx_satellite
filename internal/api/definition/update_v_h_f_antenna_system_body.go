// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateVHFAntennaSystemBody update v h f antenna system body
//
// swagger:model UpdateVHFAntennaSystemBody
type UpdateVHFAntennaSystemBody struct {

	// frequency
	Frequency float64 `json:"frequency,omitempty"`

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

	// weight
	Weight float64 `json:"weight,omitempty"`

	// width
	Width float64 `json:"width,omitempty"`
}

// Validate validates this update v h f antenna system body
func (m *UpdateVHFAntennaSystemBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update v h f antenna system body based on context it is used
func (m *UpdateVHFAntennaSystemBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVHFAntennaSystemBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVHFAntennaSystemBody) UnmarshalBinary(b []byte) error {
	var res UpdateVHFAntennaSystemBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
