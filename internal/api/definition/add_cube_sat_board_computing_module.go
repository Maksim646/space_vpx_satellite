// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddCubeSatBoardComputingModule add cube sat board computing module
//
// swagger:model AddCubeSatBoardComputingModule
type AddCubeSatBoardComputingModule struct {

	// board computing module name
	// Required: true
	BoardComputingModuleName *string `json:"board_computing_module_name"`
}

// Validate validates this add cube sat board computing module
func (m *AddCubeSatBoardComputingModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoardComputingModuleName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCubeSatBoardComputingModule) validateBoardComputingModuleName(formats strfmt.Registry) error {

	if err := validate.Required("board_computing_module_name", "body", m.BoardComputingModuleName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add cube sat board computing module based on context it is used
func (m *AddCubeSatBoardComputingModule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddCubeSatBoardComputingModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCubeSatBoardComputingModule) UnmarshalBinary(b []byte) error {
	var res AddCubeSatBoardComputingModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
