// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemDiagnosticsEntity system diagnostics entity
// swagger:model SystemDiagnosticsEntity
type SystemDiagnosticsEntity struct {

	// system diagnostics
	SystemDiagnostics *SystemDiagnosticsDTO `json:"systemDiagnostics,omitempty"`
}

// Validate validates this system diagnostics entity
func (m *SystemDiagnosticsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSystemDiagnostics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemDiagnosticsEntity) validateSystemDiagnostics(formats strfmt.Registry) error {

	if swag.IsZero(m.SystemDiagnostics) { // not required
		return nil
	}

	if m.SystemDiagnostics != nil {
		if err := m.SystemDiagnostics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemDiagnostics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemDiagnosticsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDiagnosticsEntity) UnmarshalBinary(b []byte) error {
	var res SystemDiagnosticsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
