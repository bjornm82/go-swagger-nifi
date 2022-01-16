// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoteProcessGroupsEntity remote process groups entity
//
// swagger:model RemoteProcessGroupsEntity
type RemoteProcessGroupsEntity struct {

	// remote process groups
	// Unique: true
	RemoteProcessGroups []*RemoteProcessGroupEntity `json:"remoteProcessGroups"`
}

// Validate validates this remote process groups entity
func (m *RemoteProcessGroupsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteProcessGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupsEntity) validateRemoteProcessGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.RemoteProcessGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("remoteProcessGroups", "body", m.RemoteProcessGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.RemoteProcessGroups); i++ {
		if swag.IsZero(m.RemoteProcessGroups[i]) { // not required
			continue
		}

		if m.RemoteProcessGroups[i] != nil {
			if err := m.RemoteProcessGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteProcessGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupsEntity) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
