// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HistoryEntity history entity
// swagger:model HistoryEntity
type HistoryEntity struct {

	// history
	History *HistoryDTO `json:"history,omitempty"`
}

// Validate validates this history entity
func (m *HistoryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryEntity) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.History) { // not required
		return nil
	}

	if m.History != nil {
		if err := m.History.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryEntity) UnmarshalBinary(b []byte) error {
	var res HistoryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
