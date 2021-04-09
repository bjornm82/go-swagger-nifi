// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProcessorsRunStatusDetailsEntity processors run status details entity
// swagger:model ProcessorsRunStatusDetailsEntity
type ProcessorsRunStatusDetailsEntity struct {

	// run status details
	RunStatusDetails []*ProcessorRunStatusDetailsEntity `json:"runStatusDetails"`
}

// Validate validates this processors run status details entity
func (m *ProcessorsRunStatusDetailsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunStatusDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorsRunStatusDetailsEntity) validateRunStatusDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RunStatusDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.RunStatusDetails); i++ {
		if swag.IsZero(m.RunStatusDetails[i]) { // not required
			continue
		}

		if m.RunStatusDetails[i] != nil {
			if err := m.RunStatusDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runStatusDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorsRunStatusDetailsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorsRunStatusDetailsEntity) UnmarshalBinary(b []byte) error {
	var res ProcessorsRunStatusDetailsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}