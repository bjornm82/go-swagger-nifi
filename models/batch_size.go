// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BatchSize batch size
//
// swagger:model BatchSize
type BatchSize struct {

	// Preferred number of flow files to include in a transaction.
	Count int32 `json:"count,omitempty"`

	// Preferred amount of time that a transaction should span.
	Duration string `json:"duration,omitempty"`

	// Preferred number of bytes to include in a transaction.
	Size string `json:"size,omitempty"`
}

// Validate validates this batch size
func (m *BatchSize) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BatchSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchSize) UnmarshalBinary(b []byte) error {
	var res BatchSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
