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

// VersionedFlowsEntity versioned flows entity
//
// swagger:model VersionedFlowsEntity
type VersionedFlowsEntity struct {

	// versioned flows
	// Unique: true
	VersionedFlows []*VersionedFlowEntity `json:"versionedFlows"`
}

// Validate validates this versioned flows entity
func (m *VersionedFlowsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionedFlows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowsEntity) validateVersionedFlows(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionedFlows) { // not required
		return nil
	}

	if err := validate.UniqueItems("versionedFlows", "body", m.VersionedFlows); err != nil {
		return err
	}

	for i := 0; i < len(m.VersionedFlows); i++ {
		if swag.IsZero(m.VersionedFlows[i]) { // not required
			continue
		}

		if m.VersionedFlows[i] != nil {
			if err := m.VersionedFlows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versionedFlows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowsEntity) UnmarshalBinary(b []byte) error {
	var res VersionedFlowsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
