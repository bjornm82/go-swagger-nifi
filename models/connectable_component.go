// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectableComponent connectable component
// swagger:model ConnectableComponent
type ConnectableComponent struct {

	// The comments for the connectable component.
	Comments string `json:"comments,omitempty"`

	// The id of the group that the connectable component resides in
	// Required: true
	GroupID *string `json:"groupId"`

	// The id of the connectable component.
	// Required: true
	ID *string `json:"id"`

	// The name of the connectable component
	Name string `json:"name,omitempty"`

	// The type of component the connectable is.
	// Required: true
	// Enum: [PROCESSOR REMOTE_INPUT_PORT REMOTE_OUTPUT_PORT INPUT_PORT OUTPUT_PORT FUNNEL]
	Type *string `json:"type"`
}

// Validate validates this connectable component
func (m *ConnectableComponent) Validate(formats strfmt.Registry) error {
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

func (m *ConnectableComponent) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *ConnectableComponent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var connectableComponentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROCESSOR","REMOTE_INPUT_PORT","REMOTE_OUTPUT_PORT","INPUT_PORT","OUTPUT_PORT","FUNNEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectableComponentTypeTypePropEnum = append(connectableComponentTypeTypePropEnum, v)
	}
}

const (

	// ConnectableComponentTypePROCESSOR captures enum value "PROCESSOR"
	ConnectableComponentTypePROCESSOR string = "PROCESSOR"

	// ConnectableComponentTypeREMOTEINPUTPORT captures enum value "REMOTE_INPUT_PORT"
	ConnectableComponentTypeREMOTEINPUTPORT string = "REMOTE_INPUT_PORT"

	// ConnectableComponentTypeREMOTEOUTPUTPORT captures enum value "REMOTE_OUTPUT_PORT"
	ConnectableComponentTypeREMOTEOUTPUTPORT string = "REMOTE_OUTPUT_PORT"

	// ConnectableComponentTypeINPUTPORT captures enum value "INPUT_PORT"
	ConnectableComponentTypeINPUTPORT string = "INPUT_PORT"

	// ConnectableComponentTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	ConnectableComponentTypeOUTPUTPORT string = "OUTPUT_PORT"

	// ConnectableComponentTypeFUNNEL captures enum value "FUNNEL"
	ConnectableComponentTypeFUNNEL string = "FUNNEL"
)

// prop value enum
func (m *ConnectableComponent) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, connectableComponentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConnectableComponent) validateType(formats strfmt.Registry) error {

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
func (m *ConnectableComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectableComponent) UnmarshalBinary(b []byte) error {
	var res ConnectableComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
