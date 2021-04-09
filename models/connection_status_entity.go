// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConnectionStatusEntity connection status entity
// swagger:model ConnectionStatusEntity
type ConnectionStatusEntity struct {

	// Indicates whether the user can read a given resource.
	// Read Only: true
	CanRead *bool `json:"canRead,omitempty"`

	// connection status
	ConnectionStatus *ConnectionStatusDTO `json:"connectionStatus,omitempty"`
}

// Validate validates this connection status entity
func (m *ConnectionStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionStatusEntity) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionStatusEntity) UnmarshalBinary(b []byte) error {
	var res ConnectionStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
