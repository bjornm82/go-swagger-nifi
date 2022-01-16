// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueueSizeDTO queue size d t o
//
// swagger:model QueueSizeDTO
type QueueSizeDTO struct {

	// The size of objects in a queue.
	ByteCount int64 `json:"byteCount,omitempty"`

	// The count of objects in a queue.
	ObjectCount int32 `json:"objectCount,omitempty"`
}

// Validate validates this queue size d t o
func (m *QueueSizeDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueueSizeDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueueSizeDTO) UnmarshalBinary(b []byte) error {
	var res QueueSizeDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
