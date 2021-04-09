// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ExportTemplate exports a template
*/
func (a *Client) ExportTemplate(params *ExportTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*ExportTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportTemplate",
		Method:             "GET",
		PathPattern:        "/templates/{id}/download",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExportTemplateOK), nil

}

/*
RemoveTemplate deletes a template
*/
func (a *Client) RemoveTemplate(params *RemoveTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeTemplate",
		Method:             "DELETE",
		PathPattern:        "/templates/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveTemplateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
