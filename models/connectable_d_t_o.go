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

// ConnectableDTO connectable d t o
//
// swagger:model ConnectableDTO
type ConnectableDTO struct {

	// The comments for the connectable component.
	Comments string `json:"comments,omitempty"`

	// If the connectable component represents a remote port, indicates if the target exists.
	Exists bool `json:"exists,omitempty"`

	// The id of the group that the connectable component resides in
	// Required: true
	GroupID *string `json:"groupId"`

	// The id of the connectable component.
	// Required: true
	ID *string `json:"id"`

	// The name of the connectable component
	Name string `json:"name,omitempty"`

	// Reflects the current state of the connectable component.
	Running bool `json:"running,omitempty"`

	// If the connectable component represents a remote port, indicates if the target is configured to transmit.
	Transmitting bool `json:"transmitting,omitempty"`

	// The type of component the connectable is.
	// Required: true
	// Enum: [PROCESSOR REMOTE_INPUT_PORT REMOTE_OUTPUT_PORT INPUT_PORT OUTPUT_PORT FUNNEL]
	Type *string `json:"type"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`
}

// Validate validates this connectable d t o
func (m *ConnectableDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectableDTO) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *ConnectableDTO) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var connectableDTOTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROCESSOR","REMOTE_INPUT_PORT","REMOTE_OUTPUT_PORT","INPUT_PORT","OUTPUT_PORT","FUNNEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectableDTOTypeTypePropEnum = append(connectableDTOTypeTypePropEnum, v)
	}
}

const (

	// ConnectableDTOTypePROCESSOR captures enum value "PROCESSOR"
	ConnectableDTOTypePROCESSOR string = "PROCESSOR"

	// ConnectableDTOTypeREMOTEINPUTPORT captures enum value "REMOTE_INPUT_PORT"
	ConnectableDTOTypeREMOTEINPUTPORT string = "REMOTE_INPUT_PORT"

	// ConnectableDTOTypeREMOTEOUTPUTPORT captures enum value "REMOTE_OUTPUT_PORT"
	ConnectableDTOTypeREMOTEOUTPUTPORT string = "REMOTE_OUTPUT_PORT"

	// ConnectableDTOTypeINPUTPORT captures enum value "INPUT_PORT"
	ConnectableDTOTypeINPUTPORT string = "INPUT_PORT"

	// ConnectableDTOTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	ConnectableDTOTypeOUTPUTPORT string = "OUTPUT_PORT"

	// ConnectableDTOTypeFUNNEL captures enum value "FUNNEL"
	ConnectableDTOTypeFUNNEL string = "FUNNEL"
)

// prop value enum
func (m *ConnectableDTO) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectableDTOTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectableDTO) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectableDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectableDTO) UnmarshalBinary(b []byte) error {
	var res ConnectableDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
