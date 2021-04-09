// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BulletinBoardEntity bulletin board entity
// swagger:model BulletinBoardEntity
type BulletinBoardEntity struct {

	// bulletin board
	BulletinBoard *BulletinBoardDTO `json:"bulletinBoard,omitempty"`
}

// Validate validates this bulletin board entity
func (m *BulletinBoardEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletinBoard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulletinBoardEntity) validateBulletinBoard(formats strfmt.Registry) error {

	if swag.IsZero(m.BulletinBoard) { // not required
		return nil
	}

	if m.BulletinBoard != nil {
		if err := m.BulletinBoard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bulletinBoard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulletinBoardEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulletinBoardEntity) UnmarshalBinary(b []byte) error {
	var res BulletinBoardEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}