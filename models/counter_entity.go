// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CounterEntity counter entity
// swagger:model CounterEntity
type CounterEntity struct {

	// counter
	Counter *CounterDTO `json:"counter,omitempty"`
}

// Validate validates this counter entity
func (m *CounterEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterEntity) validateCounter(formats strfmt.Registry) error {

	if swag.IsZero(m.Counter) { // not required
		return nil
	}

	if m.Counter != nil {
		if err := m.Counter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CounterEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterEntity) UnmarshalBinary(b []byte) error {
	var res CounterEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}