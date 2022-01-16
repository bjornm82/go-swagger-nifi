// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchResultsEntity search results entity
//
// swagger:model SearchResultsEntity
type SearchResultsEntity struct {

	// search results d t o
	SearchResultsDTO *SearchResultsDTO `json:"searchResultsDTO,omitempty"`
}

// Validate validates this search results entity
func (m *SearchResultsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchResultsDTO(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResultsEntity) validateSearchResultsDTO(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchResultsDTO) { // not required
		return nil
	}

	if m.SearchResultsDTO != nil {
		if err := m.SearchResultsDTO.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("searchResultsDTO")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultsEntity) UnmarshalBinary(b []byte) error {
	var res SearchResultsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
