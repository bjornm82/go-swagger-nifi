// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterSearchResultsEntity cluster search results entity
//
// swagger:model ClusterSearchResultsEntity
type ClusterSearchResultsEntity struct {

	// node results
	NodeResults []*NodeSearchResultDTO `json:"nodeResults"`
}

// Validate validates this cluster search results entity
func (m *ClusterSearchResultsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSearchResultsEntity) validateNodeResults(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeResults) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeResults); i++ {
		if swag.IsZero(m.NodeResults[i]) { // not required
			continue
		}

		if m.NodeResults[i] != nil {
			if err := m.NodeResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSearchResultsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSearchResultsEntity) UnmarshalBinary(b []byte) error {
	var res ClusterSearchResultsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
