// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccessPolicy(params *CreateAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAccessPolicyOK, error)

	GetAccessPolicy(params *GetAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccessPolicyOK, error)

	GetAccessPolicyForResource(params *GetAccessPolicyForResourceParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccessPolicyForResourceOK, error)

	RemoveAccessPolicy(params *RemoveAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveAccessPolicyOK, error)

	UpdateAccessPolicy(params *UpdateAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAccessPolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAccessPolicy creates an access policy
*/
func (a *Client) CreateAccessPolicy(params *CreateAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAccessPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccessPolicy",
		Method:             "POST",
		PathPattern:        "/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAccessPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccessPolicy gets an access policy
*/
func (a *Client) GetAccessPolicy(params *GetAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccessPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessPolicy",
		Method:             "GET",
		PathPattern:        "/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccessPolicyForResource gets an access policy for the specified action and resource

  Will return the effective policy if no component specific policy exists for the specified action and resource. Must have Read permissions to the policy with the desired action and resource. Permissions for the policy that is returned will be indicated in the response. This means the client could be authorized to get the policy for a given component but the effective policy may be inherited from an ancestor Process Group. If the client does not have permissions to that policy, the response will not include the policy and the permissions in the response will be marked accordingly. If the client does not have permissions to the policy of the desired action and resource a 403 response will be returned.
*/
func (a *Client) GetAccessPolicyForResource(params *GetAccessPolicyForResourceParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccessPolicyForResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessPolicyForResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessPolicyForResource",
		Method:             "GET",
		PathPattern:        "/policies/{action}/{resource}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessPolicyForResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessPolicyForResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessPolicyForResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveAccessPolicy deletes an access policy
*/
func (a *Client) RemoveAccessPolicy(params *RemoveAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveAccessPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveAccessPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeAccessPolicy",
		Method:             "DELETE",
		PathPattern:        "/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveAccessPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveAccessPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeAccessPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAccessPolicy updates a access policy
*/
func (a *Client) UpdateAccessPolicy(params *UpdateAccessPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAccessPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccessPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAccessPolicy",
		Method:             "PUT",
		PathPattern:        "/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAccessPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAccessPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAccessPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
