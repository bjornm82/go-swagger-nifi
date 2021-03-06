// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BulletinEntity bulletin entity
//
// swagger:model BulletinEntity
type BulletinEntity struct {

	// bulletin
	Bulletin *BulletinDTO `json:"bulletin,omitempty"`

	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// node address
	NodeAddress string `json:"nodeAddress,omitempty"`

	// source Id
	SourceID string `json:"sourceId,omitempty"`

	// When this bulletin was generated.
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this bulletin entity
func (m *BulletinEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulletinEntity) validateBulletin(formats strfmt.Registry) error {

	if swag.IsZero(m.Bulletin) { // not required
		return nil
	}

	if m.Bulletin != nil {
		if err := m.Bulletin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bulletin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulletinEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulletinEntity) UnmarshalBinary(b []byte) error {
	var res BulletinEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
