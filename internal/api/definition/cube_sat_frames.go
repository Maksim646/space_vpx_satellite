// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CubeSatFrames cube sat frames
//
// swagger:model CubeSatFrames
type CubeSatFrames struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// cube sat frames
	CubeSatFrames []*CubeSatFrame `json:"cube_sat_frames"`
}

// Validate validates this cube sat frames
func (m *CubeSatFrames) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCubeSatFrames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CubeSatFrames) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *CubeSatFrames) validateCubeSatFrames(formats strfmt.Registry) error {
	if swag.IsZero(m.CubeSatFrames) { // not required
		return nil
	}

	for i := 0; i < len(m.CubeSatFrames); i++ {
		if swag.IsZero(m.CubeSatFrames[i]) { // not required
			continue
		}

		if m.CubeSatFrames[i] != nil {
			if err := m.CubeSatFrames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cube_sat_frames" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cube_sat_frames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cube sat frames based on the context it is used
func (m *CubeSatFrames) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCubeSatFrames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CubeSatFrames) contextValidateCubeSatFrames(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CubeSatFrames); i++ {

		if m.CubeSatFrames[i] != nil {

			if swag.IsZero(m.CubeSatFrames[i]) { // not required
				return nil
			}

			if err := m.CubeSatFrames[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cube_sat_frames" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cube_sat_frames" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CubeSatFrames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CubeSatFrames) UnmarshalBinary(b []byte) error {
	var res CubeSatFrames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
