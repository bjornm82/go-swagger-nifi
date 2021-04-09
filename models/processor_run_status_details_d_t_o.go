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

// ProcessorRunStatusDetailsDTO processor run status details d t o
// swagger:model ProcessorRunStatusDetailsDTO
type ProcessorRunStatusDetailsDTO struct {

	// The current number of threads that the processor is currently using
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`

	// The ID of the processor
	ID string `json:"id,omitempty"`

	// The name of the processor
	Name string `json:"name,omitempty"`

	// The run status of the processor
	// Enum: [Running Stopped Invalid Validating Disabled]
	RunStatus string `json:"runStatus,omitempty"`

	// The processor's validation errors
	// Unique: true
	ValidationErrors []string `json:"validationErrors"`
}

// Validate validates this processor run status details d t o
func (m *ProcessorRunStatusDetailsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var processorRunStatusDetailsDTOTypeRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Running","Stopped","Invalid","Validating","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		processorRunStatusDetailsDTOTypeRunStatusPropEnum = append(processorRunStatusDetailsDTOTypeRunStatusPropEnum, v)
	}
}

const (

	// ProcessorRunStatusDetailsDTORunStatusRunning captures enum value "Running"
	ProcessorRunStatusDetailsDTORunStatusRunning string = "Running"

	// ProcessorRunStatusDetailsDTORunStatusStopped captures enum value "Stopped"
	ProcessorRunStatusDetailsDTORunStatusStopped string = "Stopped"

	// ProcessorRunStatusDetailsDTORunStatusInvalid captures enum value "Invalid"
	ProcessorRunStatusDetailsDTORunStatusInvalid string = "Invalid"

	// ProcessorRunStatusDetailsDTORunStatusValidating captures enum value "Validating"
	ProcessorRunStatusDetailsDTORunStatusValidating string = "Validating"

	// ProcessorRunStatusDetailsDTORunStatusDisabled captures enum value "Disabled"
	ProcessorRunStatusDetailsDTORunStatusDisabled string = "Disabled"
)

// prop value enum
func (m *ProcessorRunStatusDetailsDTO) validateRunStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, processorRunStatusDetailsDTOTypeRunStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProcessorRunStatusDetailsDTO) validateRunStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunStatusEnum("runStatus", "body", m.RunStatus); err != nil {
		return err
	}

	return nil
}

func (m *ProcessorRunStatusDetailsDTO) validateValidationErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationErrors) { // not required
		return nil
	}

	if err := validate.UniqueItems("validationErrors", "body", m.ValidationErrors); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorRunStatusDetailsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorRunStatusDetailsDTO) UnmarshalBinary(b []byte) error {
	var res ProcessorRunStatusDetailsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
