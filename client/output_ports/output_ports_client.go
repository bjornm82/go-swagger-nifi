// Code generated by go-swagger; DO NOT EDIT.

package output_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new output ports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for output ports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetOutputPort gets an output port
*/
func (a *Client) GetOutputPort(params *GetOutputPortParams, authInfo runtime.ClientAuthInfoWriter) (*GetOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOutputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOutputPort",
		Method:             "GET",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOutputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOutputPortOK), nil

}

/*
OutputPortsUpdateRunStatus updates run status of an output port
*/
func (a *Client) OutputPortsUpdateRunStatus(params *OutputPortsUpdateRunStatusParams, authInfo runtime.ClientAuthInfoWriter) (*OutputPortsUpdateRunStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOutputPortsUpdateRunStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "outputPortsUpdateRunStatus",
		Method:             "PUT",
		PathPattern:        "/output-ports/{id}/run-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OutputPortsUpdateRunStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OutputPortsUpdateRunStatusOK), nil

}

/*
RemoveOutputPort deletes an output port
*/
func (a *Client) RemoveOutputPort(params *RemoveOutputPortParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOutputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeOutputPort",
		Method:             "DELETE",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveOutputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveOutputPortOK), nil

}

/*
UpdateOutputPort updates an output port
*/
func (a *Client) UpdateOutputPort(params *UpdateOutputPortParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOutputPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOutputPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOutputPort",
		Method:             "PUT",
		PathPattern:        "/output-ports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOutputPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOutputPortOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}