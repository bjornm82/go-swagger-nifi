// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserCreated, error)

	CreateUserGroup(params *CreateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserGroupCreated, error)

	GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error)

	GetUserGroup(params *GetUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGroupOK, error)

	GetUserGroups(params *GetUserGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGroupsOK, error)

	GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error)

	RemoveUser(params *RemoveUserParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveUserOK, error)

	RemoveUserGroup(params *RemoveUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveUserGroupOK, error)

	SearchTenants(params *SearchTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*SearchTenantsOK, error)

	UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserOK, error)

	UpdateUserGroup(params *UpdateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUser creates a user

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/tenants/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateUserGroup creates a user group

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) CreateUserGroup(params *CreateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUserGroup",
		Method:             "POST",
		PathPattern:        "/tenants/user-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUser gets a user

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/tenants/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserGroup gets a user group

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetUserGroup(params *GetUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserGroup",
		Method:             "GET",
		PathPattern:        "/tenants/user-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserGroups gets all user groups

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetUserGroups(params *GetUserGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserGroups",
		Method:             "GET",
		PathPattern:        "/tenants/user-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsers gets all users

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/tenants/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveUser deletes a user

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) RemoveUser(params *RemoveUserParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeUser",
		Method:             "DELETE",
		PathPattern:        "/tenants/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveUserGroup deletes a user group

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) RemoveUserGroup(params *RemoveUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveUserGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUserGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeUserGroup",
		Method:             "DELETE",
		PathPattern:        "/tenants/user-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveUserGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveUserGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchTenants searches for a tenant with the specified identity

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SearchTenants(params *SearchTenantsParams, authInfo runtime.ClientAuthInfoWriter) (*SearchTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchTenants",
		Method:             "GET",
		PathPattern:        "/tenants/search-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchTenantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchTenantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchTenants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUser updates a user

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUser",
		Method:             "PUT",
		PathPattern:        "/tenants/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUserGroup updates a user group

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) UpdateUserGroup(params *UpdateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserGroup",
		Method:             "PUT",
		PathPattern:        "/tenants/user-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
