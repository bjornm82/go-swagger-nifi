// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GarbageCollectionDTO garbage collection d t o
// swagger:model GarbageCollectionDTO
type GarbageCollectionDTO struct {

	// The number of times garbage collection has run.
	CollectionCount int64 `json:"collectionCount,omitempty"`

	// The total number of milliseconds spent garbage collecting.
	CollectionMillis int64 `json:"collectionMillis,omitempty"`

	// The total amount of time spent garbage collecting.
	CollectionTime string `json:"collectionTime,omitempty"`

	// The name of the garbage collector.
	Name string `json:"name,omitempty"`
}

// Validate validates this garbage collection d t o
func (m *GarbageCollectionDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GarbageCollectionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GarbageCollectionDTO) UnmarshalBinary(b []byte) error {
	var res GarbageCollectionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
