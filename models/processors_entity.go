// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProcessorsEntity processors entity
// swagger:model ProcessorsEntity
type ProcessorsEntity struct {

	// processors
	// Unique: true
	Processors []*ProcessorEntity `json:"processors"`
}

// Validate validates this processors entity
func (m *ProcessorsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorsEntity) validateProcessors(formats strfmt.Registry) error {

	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	if err := validate.UniqueItems("processors", "body", m.Processors); err != nil {
		return err
	}

	for i := 0; i < len(m.Processors); i++ {
		if swag.IsZero(m.Processors[i]) { // not required
			continue
		}

		if m.Processors[i] != nil {
			if err := m.Processors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorsEntity) UnmarshalBinary(b []byte) error {
	var res ProcessorsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
