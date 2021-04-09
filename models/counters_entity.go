// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CountersEntity counters entity
// swagger:model CountersEntity
type CountersEntity struct {

	// counters
	Counters *CountersDTO `json:"counters,omitempty"`
}

// Validate validates this counters entity
func (m *CountersEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CountersEntity) validateCounters(formats strfmt.Registry) error {

	if swag.IsZero(m.Counters) { // not required
		return nil
	}

	if m.Counters != nil {
		if err := m.Counters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CountersEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountersEntity) UnmarshalBinary(b []byte) error {
	var res CountersEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
