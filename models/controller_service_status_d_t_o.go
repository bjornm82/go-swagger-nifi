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

// ControllerServiceStatusDTO controller service status d t o
//
// swagger:model ControllerServiceStatusDTO
type ControllerServiceStatusDTO struct {

	// The number of active threads for the component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`

	// The run status of this ControllerService
	// Enum: [ENABLED ENABLING DISABLED DISABLING]
	RunStatus string `json:"runStatus,omitempty"`

	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	// Enum: [VALID INVALID VALIDATING]
	ValidationStatus string `json:"validationStatus,omitempty"`
}

// Validate validates this controller service status d t o
func (m *ControllerServiceStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var controllerServiceStatusDTOTypeRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","ENABLING","DISABLED","DISABLING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		controllerServiceStatusDTOTypeRunStatusPropEnum = append(controllerServiceStatusDTOTypeRunStatusPropEnum, v)
	}
}

const (

	// ControllerServiceStatusDTORunStatusENABLED captures enum value "ENABLED"
	ControllerServiceStatusDTORunStatusENABLED string = "ENABLED"

	// ControllerServiceStatusDTORunStatusENABLING captures enum value "ENABLING"
	ControllerServiceStatusDTORunStatusENABLING string = "ENABLING"

	// ControllerServiceStatusDTORunStatusDISABLED captures enum value "DISABLED"
	ControllerServiceStatusDTORunStatusDISABLED string = "DISABLED"

	// ControllerServiceStatusDTORunStatusDISABLING captures enum value "DISABLING"
	ControllerServiceStatusDTORunStatusDISABLING string = "DISABLING"
)

// prop value enum
func (m *ControllerServiceStatusDTO) validateRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, controllerServiceStatusDTOTypeRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ControllerServiceStatusDTO) validateRunStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunStatusEnum("runStatus", "body", m.RunStatus); err != nil {
		return err
	}

	return nil
}

var controllerServiceStatusDTOTypeValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VALID","INVALID","VALIDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		controllerServiceStatusDTOTypeValidationStatusPropEnum = append(controllerServiceStatusDTOTypeValidationStatusPropEnum, v)
	}
}

const (

	// ControllerServiceStatusDTOValidationStatusVALID captures enum value "VALID"
	ControllerServiceStatusDTOValidationStatusVALID string = "VALID"

	// ControllerServiceStatusDTOValidationStatusINVALID captures enum value "INVALID"
	ControllerServiceStatusDTOValidationStatusINVALID string = "INVALID"

	// ControllerServiceStatusDTOValidationStatusVALIDATING captures enum value "VALIDATING"
	ControllerServiceStatusDTOValidationStatusVALIDATING string = "VALIDATING"
)

// prop value enum
func (m *ControllerServiceStatusDTO) validateValidationStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, controllerServiceStatusDTOTypeValidationStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ControllerServiceStatusDTO) validateValidationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationStatusEnum("validationStatus", "body", m.ValidationStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerServiceStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerServiceStatusDTO) UnmarshalBinary(b []byte) error {
	var res ControllerServiceStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
