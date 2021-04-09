// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StartVersionControlRequestEntity start version control request entity
// swagger:model StartVersionControlRequestEntity
type StartVersionControlRequestEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The Revision of the Process Group under Version Control
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The versioned flow
	VersionedFlow *VersionedFlowDTO `json:"versionedFlow,omitempty"`
}

// Validate validates this start version control request entity
func (m *StartVersionControlRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionedFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StartVersionControlRequestEntity) validateProcessGroupRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessGroupRevision) { // not required
		return nil
	}

	if m.ProcessGroupRevision != nil {
		if err := m.ProcessGroupRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *StartVersionControlRequestEntity) validateVersionedFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionedFlow) { // not required
		return nil
	}

	if m.VersionedFlow != nil {
		if err := m.VersionedFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StartVersionControlRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartVersionControlRequestEntity) UnmarshalBinary(b []byte) error {
	var res StartVersionControlRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
