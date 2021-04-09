// Code generated by go-swagger; DO NOT EDIT.

package input_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new input ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for input ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInputPort gets an input port
*/
func (a *Client) GetInputPort(params *GetInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*GetInputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInputPort",
		Method:             "GET",
		PathPattern:        "/input-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInputPortOK), nil

}

/*
InputPortsUpdateRunStatus updates run status of an input port
*/
func (a *Client) InputPortsUpdateRunStatus(params *InputPortsUpdateRunStatusParams, authInfo runtime.ClientAuthInfoWriter) (*InputPortsUpdateRunStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInputPortsUpdateRunStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "inputPortsUpdateRunStatus",
		Method:             "PUT",
		PathPattern:        "/input-ports/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InputPortsUpdateRunStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InputPortsUpdateRunStatusOK), nil

}

/*
RemoveInputPort deletes an input port
*/
func (a *Client) RemoveInputPort(params *RemoveInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveInputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveInputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeInputPort",
		Method:             "DELETE",
		PathPattern:        "/input-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveInputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveInputPortOK), nil

}

/*
UpdateInputPort updates an input port
*/
func (a *Client) UpdateInputPort(params *UpdateInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateInputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateInputPort",
		Method:             "PUT",
		PathPattern:        "/input-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateInputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateInputPortOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}