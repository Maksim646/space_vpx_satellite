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

// ChassisesVPX chassises v p x
//
// swagger:model ChassisesVPX
type ChassisesVPX struct {

	// chassises
	Chassises []*ChassisVPX `json:"chassises"`

	// count
	// Required: true
	Count *int64 `json:"count"`
}

// Validate validates this chassises v p x
func (m *ChassisesVPX) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChassises(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisesVPX) validateChassises(formats strfmt.Registry) error {
	if swag.IsZero(m.Chassises) { // not required
		return nil
	}

	for i := 0; i < len(m.Chassises); i++ {
		if swag.IsZero(m.Chassises[i]) { // not required
			continue
		}

		if m.Chassises[i] != nil {
			if err := m.Chassises[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chassises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chassises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChassisesVPX) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this chassises v p x based on the context it is used
func (m *ChassisesVPX) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChassises(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChassisesVPX) contextValidateChassises(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Chassises); i++ {

		if m.Chassises[i] != nil {

			if swag.IsZero(m.Chassises[i]) { // not required
				return nil
			}

			if err := m.Chassises[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chassises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chassises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChassisesVPX) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisesVPX) UnmarshalBinary(b []byte) error {
	var res ChassisesVPX
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
