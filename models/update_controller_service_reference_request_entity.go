// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateControllerServiceReferenceRequestEntity update controller service reference request entity
//
// swagger:model UpdateControllerServiceReferenceRequestEntity
type UpdateControllerServiceReferenceRequestEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The identifier of the Controller Service.
	ID string `json:"id,omitempty"`

	// The revisions for all referencing components.
	ReferencingComponentRevisions map[string]RevisionDTO `json:"referencingComponentRevisions,omitempty"`

	// The new state of the references for the controller service.
	// Enum: [ENABLED DISABLED RUNNING STOPPED]
	State string `json:"state,omitempty"`
}

// Validate validates this update controller service reference request entity
func (m *UpdateControllerServiceReferenceRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferencingComponentRevisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateControllerServiceReferenceRequestEntity) validateReferencingComponentRevisions(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferencingComponentRevisions) { // not required
		return nil
	}

	for k := range m.ReferencingComponentRevisions {

		if err := validate.Required("referencingComponentRevisions"+"."+k, "body", m.ReferencingComponentRevisions[k]); err != nil {
			return err
		}
		if val, ok := m.ReferencingComponentRevisions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var updateControllerServiceReferenceRequestEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","RUNNING","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateControllerServiceReferenceRequestEntityTypeStatePropEnum = append(updateControllerServiceReferenceRequestEntityTypeStatePropEnum, v)
	}
}

const (

	// UpdateControllerServiceReferenceRequestEntityStateENABLED captures enum value "ENABLED"
	UpdateControllerServiceReferenceRequestEntityStateENABLED string = "ENABLED"

	// UpdateControllerServiceReferenceRequestEntityStateDISABLED captures enum value "DISABLED"
	UpdateControllerServiceReferenceRequestEntityStateDISABLED string = "DISABLED"

	// UpdateControllerServiceReferenceRequestEntityStateRUNNING captures enum value "RUNNING"
	UpdateControllerServiceReferenceRequestEntityStateRUNNING string = "RUNNING"

	// UpdateControllerServiceReferenceRequestEntityStateSTOPPED captures enum value "STOPPED"
	UpdateControllerServiceReferenceRequestEntityStateSTOPPED string = "STOPPED"
)

// prop value enum
func (m *UpdateControllerServiceReferenceRequestEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateControllerServiceReferenceRequestEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateControllerServiceReferenceRequestEntity) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateControllerServiceReferenceRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateControllerServiceReferenceRequestEntity) UnmarshalBinary(b []byte) error {
	var res UpdateControllerServiceReferenceRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
