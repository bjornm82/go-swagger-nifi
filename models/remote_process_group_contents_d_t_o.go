// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoteProcessGroupContentsDTO remote process group contents d t o
// swagger:model RemoteProcessGroupContentsDTO
type RemoteProcessGroupContentsDTO struct {

	// The input ports to which data can be sent.
	// Unique: true
	InputPorts []*RemoteProcessGroupPortDTO `json:"inputPorts"`

	// The output ports from which data can be retrieved.
	// Unique: true
	OutputPorts []*RemoteProcessGroupPortDTO `json:"outputPorts"`
}

// Validate validates this remote process group contents d t o
func (m *RemoteProcessGroupContentsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputPorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupContentsDTO) validateInputPorts(formats strfmt.Registry) error {

	if swag.IsZero(m.InputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("inputPorts", "body", m.InputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.InputPorts); i++ {
		if swag.IsZero(m.InputPorts[i]) { // not required
			continue
		}

		if m.InputPorts[i] != nil {
			if err := m.InputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RemoteProcessGroupContentsDTO) validateOutputPorts(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("outputPorts", "body", m.OutputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.OutputPorts); i++ {
		if swag.IsZero(m.OutputPorts[i]) { // not required
			continue
		}

		if m.OutputPorts[i] != nil {
			if err := m.OutputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupContentsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupContentsDTO) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupContentsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
