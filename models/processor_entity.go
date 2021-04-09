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

// ProcessorEntity processor entity
// swagger:model ProcessorEntity
type ProcessorEntity struct {

	// The bulletins for this component.
	Bulletins []*BulletinEntity `json:"bulletins"`

	// component
	Component *ProcessorDTO `json:"component,omitempty"`

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The input requirement for this processor.
	InputRequirement string `json:"inputRequirement,omitempty"`

	// The permissions for this component operations.
	OperatePermissions *PermissionsDTO `json:"operatePermissions,omitempty"`

	// The permissions for this component.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// status
	Status *ProcessorStatusDTO `json:"status,omitempty"`

	// The URI for futures requests to the component.
	URI string `json:"uri,omitempty"`
}

// Validate validates this processor entity
func (m *ProcessorEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorEntity) validateBulletins(formats strfmt.Registry) error {

	if swag.IsZero(m.Bulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.Bulletins); i++ {
		if swag.IsZero(m.Bulletins[i]) { // not required
			continue
		}

		if m.Bulletins[i] != nil {
			if err := m.Bulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProcessorEntity) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorEntity) validateOperatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.OperatePermissions) { // not required
		return nil
	}

	if m.OperatePermissions != nil {
		if err := m.OperatePermissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatePermissions")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorEntity) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorEntity) validatePosition(formats strfmt.Registry) error {

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

func (m *ProcessorEntity) validateRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessorEntity) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorEntity) UnmarshalBinary(b []byte) error {
	var res ProcessorEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}