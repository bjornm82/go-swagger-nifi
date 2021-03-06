// Code generated by go-swagger; DO NOT EDIT.

package provenance_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new provenance events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for provenance events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetInputContent(params *GetInputContentParams, authInfo runtime.ClientAuthInfoWriter) (*GetInputContentOK, error)

	GetOutputContent(params *GetOutputContentParams, authInfo runtime.ClientAuthInfoWriter) (*GetOutputContentOK, error)

	GetProvenanceEvent(params *GetProvenanceEventParams, authInfo runtime.ClientAuthInfoWriter) (*GetProvenanceEventOK, error)

	SubmitReplay(params *SubmitReplayParams, authInfo runtime.ClientAuthInfoWriter) (*SubmitReplayCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetInputContent gets the input content for a provenance event
*/
func (a *Client) GetInputContent(params *GetInputContentParams, authInfo runtime.ClientAuthInfoWriter) (*GetInputContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInputContentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInputContent",
		Method:             "GET",
		PathPattern:        "/provenance-events/{id}/content/input",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInputContentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInputContentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInputContent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOutputContent gets the output content for a provenance event
*/
func (a *Client) GetOutputContent(params *GetOutputContentParams, authInfo runtime.ClientAuthInfoWriter) (*GetOutputContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOutputContentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOutputContent",
		Method:             "GET",
		PathPattern:        "/provenance-events/{id}/content/output",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOutputContentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOutputContentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOutputContent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProvenanceEvent gets a provenance event
*/
func (a *Client) GetProvenanceEvent(params *GetProvenanceEventParams, authInfo runtime.ClientAuthInfoWriter) (*GetProvenanceEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvenanceEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProvenanceEvent",
		Method:             "GET",
		PathPattern:        "/provenance-events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProvenanceEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProvenanceEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProvenanceEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubmitReplay replays content from a provenance event
*/
func (a *Client) SubmitReplay(params *SubmitReplayParams, authInfo runtime.ClientAuthInfoWriter) (*SubmitReplayCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitReplayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "submitReplay",
		Method:             "POST",
		PathPattern:        "/provenance-events/replays",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SubmitReplayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubmitReplayCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitReplay: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
