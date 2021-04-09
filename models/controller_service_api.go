// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ControllerServiceAPI controller service API
// swagger:model ControllerServiceAPI
type ControllerServiceAPI struct {

	// The details of the artifact that bundled this service interface.
	Bundle *Bundle `json:"bundle,omitempty"`

	// The fully qualified name of the service interface.
	Type string `json:"type,omitempty"`
}

// Validate validates this controller service API
func (m *ControllerServiceAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBundle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServiceAPI) validateBundle(formats strfmt.Registry) error {

	if swag.IsZero(m.Bundle) { // not required
		return nil
	}

	if m.Bundle != nil {
		if err := m.Bundle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundle")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerServiceAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerServiceAPI) UnmarshalBinary(b []byte) error {
	var res ControllerServiceAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}