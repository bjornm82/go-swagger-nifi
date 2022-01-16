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

// PortRunStatusEntity port run status entity
//
// swagger:model PortRunStatusEntity
type PortRunStatusEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The run status of the Port.
	// Enum: [RUNNING STOPPED DISABLED]
	State string `json:"state,omitempty"`
}

// Validate validates this port run status entity
func (m *PortRunStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevision(formats); err != nil {
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

func (m *PortRunStatusEntity) validateRevision(formats strfmt.Registry) error {

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

var portRunStatusEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","STOPPED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portRunStatusEntityTypeStatePropEnum = append(portRunStatusEntityTypeStatePropEnum, v)
	}
}

const (

	// PortRunStatusEntityStateRUNNING captures enum value "RUNNING"
	PortRunStatusEntityStateRUNNING string = "RUNNING"

	// PortRunStatusEntityStateSTOPPED captures enum value "STOPPED"
	PortRunStatusEntityStateSTOPPED string = "STOPPED"

	// PortRunStatusEntityStateDISABLED captures enum value "DISABLED"
	PortRunStatusEntityStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *PortRunStatusEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portRunStatusEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortRunStatusEntity) validateState(formats strfmt.Registry) error {

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
func (m *PortRunStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortRunStatusEntity) UnmarshalBinary(b []byte) error {
	var res PortRunStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
