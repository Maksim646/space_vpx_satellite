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

// CreateCubeSatProjectBody create cube sat project body
//
// swagger:model CreateCubeSatProjectBody
type CreateCubeSatProjectBody struct {

	// board computing module name
	BoardComputingModuleName string `json:"board_computing_module_name,omitempty"`

	// frame name
	FrameName string `json:"frame_name,omitempty"`

	// power system name
	PowerSystemName string `json:"power_system_name,omitempty"`

	// project name
	// Required: true
	ProjectName *string `json:"project_name"`

	// solar panel side name
	SolarPanelSideName string `json:"solar_panel_side_name,omitempty"`

	// solar panel top name
	SolarPanelTopName string `json:"solar_panel_top_name,omitempty"`

	// vhf antenna system name
	VhfAntennaSystemName string `json:"vhf_antenna_system_name,omitempty"`

	// vhf transceiver name
	VhfTransceiverName string `json:"vhf_transceiver_name,omitempty"`
}

// Validate validates this create cube sat project body
func (m *CreateCubeSatProjectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCubeSatProjectBody) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("project_name", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create cube sat project body based on context it is used
func (m *CreateCubeSatProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateCubeSatProjectBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCubeSatProjectBody) UnmarshalBinary(b []byte) error {
	var res CreateCubeSatProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
