// Code generated by go-swagger; DO NOT EDIT.

package input_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new input ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for input ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetInputPort(params *GetInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*GetInputPortOK, error)

	RemoveInputPort(params *RemoveInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveInputPortOK, error)

	UpdateInputPort(params *UpdateInputPortParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateInputPortOK, error)

	UpdateRunStatusInputPorts(params *UpdateRunStatusInputPortsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunStatusInputPortsOK, error)

	SetTransport(transport runtime.ClientTransport)
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
	success, ok := result.(*GetInputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
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
	success, ok := result.(*RemoveInputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeInputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
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
	success, ok := result.(*UpdateInputPortOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateInputPort: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRunStatusInputPorts updates run status of an input port
*/
func (a *Client) UpdateRunStatusInputPorts(params *UpdateRunStatusInputPortsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRunStatusInputPortsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRunStatusInputPortsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRunStatusInputPorts",
		Method:             "PUT",
		PathPattern:        "/input-ports/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateRunStatusInputPortsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRunStatusInputPortsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRunStatusInputPorts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
