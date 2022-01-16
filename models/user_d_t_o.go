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

// UserDTO user d t o
//
// swagger:model UserDTO
type UserDTO struct {

	// The access policies this user belongs to.
	// Unique: true
	AccessPolicies []*AccessPolicySummaryEntity `json:"accessPolicies"`

	// Whether this tenant is configurable.
	Configurable bool `json:"configurable,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The identity of the tenant.
	Identity string `json:"identity,omitempty"`

	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The groups to which the user belongs. This field is read only and it provided for convenience.
	// Unique: true
	UserGroups []*TenantEntity `json:"userGroups"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`
}

// Validate validates this user d t o
func (m *UserDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDTO) validateAccessPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessPolicies) { // not required
		return nil
	}

	if err := validate.UniqueItems("accessPolicies", "body", m.AccessPolicies); err != nil {
		return err
	}

	for i := 0; i < len(m.AccessPolicies); i++ {
		if swag.IsZero(m.AccessPolicies[i]) { // not required
			continue
		}

		if m.AccessPolicies[i] != nil {
			if err := m.AccessPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserDTO) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *UserDTO) validateUserGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.UserGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("userGroups", "body", m.UserGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.UserGroups); i++ {
		if swag.IsZero(m.UserGroups[i]) { // not required
			continue
		}

		if m.UserGroups[i] != nil {
			if err := m.UserGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDTO) UnmarshalBinary(b []byte) error {
	var res UserDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
