// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComponentHistoryDTO component history d t o
//
// swagger:model ComponentHistoryDTO
type ComponentHistoryDTO struct {

	// The component id.
	ComponentID string `json:"componentId,omitempty"`

	// The history for the properties of the component.
	PropertyHistory map[string]PropertyHistoryDTO `json:"propertyHistory,omitempty"`
}

// Validate validates this component history d t o
func (m *ComponentHistoryDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropertyHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentHistoryDTO) validatePropertyHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyHistory) { // not required
		return nil
	}

	for k := range m.PropertyHistory {

		if err := validate.Required("propertyHistory"+"."+k, "body", m.PropertyHistory[k]); err != nil {
			return err
		}
		if val, ok := m.PropertyHistory[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentHistoryDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentHistoryDTO) UnmarshalBinary(b []byte) error {
	var res ComponentHistoryDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
