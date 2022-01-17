// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accesssaml API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accesssaml API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccessTokenFromTicketSaml(params *CreateAccessTokenFromTicketSamlParams) (*CreateAccessTokenFromTicketSamlOK, error)

	CreateAccessTokenSaml(params *CreateAccessTokenSamlParams) (*CreateAccessTokenSamlOK, error)

	CreateDownloadTokenSaml(params *CreateDownloadTokenSamlParams) (*CreateDownloadTokenSamlOK, error)

	CreateUIExtensionTokenSaml(params *CreateUIExtensionTokenSamlParams) (*CreateUIExtensionTokenSamlOK, error)

	GetAccessStatusSaml(params *GetAccessStatusSamlParams) (*GetAccessStatusSamlOK, error)

	GetLoginConfigSaml(params *GetLoginConfigSamlParams) (*GetLoginConfigSamlOK, error)

	KnoxCallbackSaml(params *KnoxCallbackSamlParams) error

	KnoxLogoutSaml(params *KnoxLogoutSamlParams) error

	KnoxRequestSaml(params *KnoxRequestSamlParams) error

	LogOutCompleteSaml(params *LogOutCompleteSamlParams) (*LogOutCompleteSamlOK, error)

	LogOutSaml(params *LogOutSamlParams) (*LogOutSamlOK, error)

	SamlLocalLogout(params *SamlLocalLogoutParams) error

	SamlLoginExchange(params *SamlLoginExchangeParams) (*SamlLoginExchangeOK, error)

	SamlLoginHTTPPostConsumer(params *SamlLoginHTTPPostConsumerParams) error

	SamlLoginHTTPRedirectConsumer(params *SamlLoginHTTPRedirectConsumerParams) error

	SamlLoginRequest(params *SamlLoginRequestParams) error

	SamlMetadata(params *SamlMetadataParams) error

	SamlSingleLogoutHTTPPostConsumer(params *SamlSingleLogoutHTTPPostConsumerParams) error

	SamlSingleLogoutHTTPRedirectConsumer(params *SamlSingleLogoutHTTPRedirectConsumerParams) error

	SamlSingleLogoutRequest(params *SamlSingleLogoutRequestParams) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAccessTokenFromTicketSaml creates a token for accessing the r e s t API via kerberos ticket exchange s p n e g o negotiation

  The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'. It is also stored in the browser as a cookie.
*/
func (a *Client) CreateAccessTokenFromTicketSaml(params *CreateAccessTokenFromTicketSamlParams) (*CreateAccessTokenFromTicketSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenFromTicketSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccessTokenFromTicketSaml",
		Method:             "POST",
		PathPattern:        "/access/saml/kerberos",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenFromTicketSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAccessTokenFromTicketSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenFromTicketSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAccessTokenSaml creates a token for accessing the r e s t API via username password

  The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. It is stored in the browser as a cookie, but also returned inthe response body to be stored/used by third party client scripts.
*/
func (a *Client) CreateAccessTokenSaml(params *CreateAccessTokenSamlParams) (*CreateAccessTokenSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessTokenSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccessTokenSaml",
		Method:             "POST",
		PathPattern:        "/access/saml/token",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAccessTokenSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAccessTokenSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAccessTokenSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDownloadTokenSaml creates a single use access token for downloading flow file content

  The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.
*/
func (a *Client) CreateDownloadTokenSaml(params *CreateDownloadTokenSamlParams) (*CreateDownloadTokenSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDownloadTokenSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDownloadTokenSaml",
		Method:             "POST",
		PathPattern:        "/access/saml/download-token",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDownloadTokenSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDownloadTokenSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDownloadTokenSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateUIExtensionTokenSaml creates a single use access token for accessing a ni fi UI extension

  The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.
*/
func (a *Client) CreateUIExtensionTokenSaml(params *CreateUIExtensionTokenSamlParams) (*CreateUIExtensionTokenSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUIExtensionTokenSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUiExtensionTokenSaml",
		Method:             "POST",
		PathPattern:        "/access/saml/ui-extension-token",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUIExtensionTokenSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUIExtensionTokenSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUiExtensionTokenSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccessStatusSaml gets the status the client s access

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) GetAccessStatusSaml(params *GetAccessStatusSamlParams) (*GetAccessStatusSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessStatusSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessStatusSaml",
		Method:             "GET",
		PathPattern:        "/access/saml",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessStatusSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccessStatusSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccessStatusSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLoginConfigSaml retrieves the access configuration for this ni fi
*/
func (a *Client) GetLoginConfigSaml(params *GetLoginConfigSamlParams) (*GetLoginConfigSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginConfigSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoginConfigSaml",
		Method:             "GET",
		PathPattern:        "/access/saml/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLoginConfigSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLoginConfigSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLoginConfigSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KnoxCallbackSaml redirects callback URI for processing the result of the apache knox login sequence

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) KnoxCallbackSaml(params *KnoxCallbackSamlParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKnoxCallbackSamlParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "knoxCallbackSaml",
		Method:             "GET",
		PathPattern:        "/access/saml/knox/callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KnoxCallbackSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  KnoxLogoutSaml performs a logout in the apache knox

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) KnoxLogoutSaml(params *KnoxLogoutSamlParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKnoxLogoutSamlParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "knoxLogoutSaml",
		Method:             "GET",
		PathPattern:        "/access/saml/knox/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KnoxLogoutSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  KnoxRequestSaml initiates a request to authenticate through apache knox

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) KnoxRequestSaml(params *KnoxRequestSamlParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKnoxRequestSamlParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "knoxRequestSaml",
		Method:             "GET",
		PathPattern:        "/access/saml/knox/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &KnoxRequestSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  LogOutCompleteSaml completes the logout sequence by removing the cached logout request and cookie if they existed and redirects to nifi login

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) LogOutCompleteSaml(params *LogOutCompleteSamlParams) (*LogOutCompleteSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogOutCompleteSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logOutCompleteSaml",
		Method:             "GET",
		PathPattern:        "/access/saml/logout/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LogOutCompleteSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogOutCompleteSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logOutCompleteSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LogOutSaml performs a logout for other providers that have been issued a j w t

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) LogOutSaml(params *LogOutSamlParams) (*LogOutSamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogOutSamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logOutSaml",
		Method:             "DELETE",
		PathPattern:        "/access/saml/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LogOutSamlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogOutSamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logOutSaml: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SamlLocalLogout locals logout when s a m l is enabled does not communicate with the ID p

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlLocalLogout(params *SamlLocalLogoutParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlLocalLogoutParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlLocalLogout",
		Method:             "GET",
		PathPattern:        "/access/saml/local-logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlLocalLogoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlLoginExchange retrieves a j w t following a successful login sequence using the configured s a m l identity provider

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlLoginExchange(params *SamlLoginExchangeParams) (*SamlLoginExchangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlLoginExchangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlLoginExchange",
		Method:             "POST",
		PathPattern:        "/access/saml/login/exchange",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlLoginExchangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SamlLoginExchangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for samlLoginExchange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SamlLoginHTTPPostConsumer processes the s s o response from the s a m l identity provider for HTTP p o s t binding

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlLoginHTTPPostConsumer(params *SamlLoginHTTPPostConsumerParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlLoginHTTPPostConsumerParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlLoginHttpPostConsumer",
		Method:             "POST",
		PathPattern:        "/access/saml/login/consumer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlLoginHTTPPostConsumerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlLoginHTTPRedirectConsumer processes the s s o response from the s a m l identity provider for HTTP r e d i r e c t binding

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlLoginHTTPRedirectConsumer(params *SamlLoginHTTPRedirectConsumerParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlLoginHTTPRedirectConsumerParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlLoginHttpRedirectConsumer",
		Method:             "GET",
		PathPattern:        "/access/saml/login/consumer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlLoginHTTPRedirectConsumerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlLoginRequest initiates an s s o request to the configured s a m l identity provider

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlLoginRequest(params *SamlLoginRequestParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlLoginRequestParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlLoginRequest",
		Method:             "GET",
		PathPattern:        "/access/saml/login/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlMetadata retrieves the service provider metadata

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlMetadata(params *SamlMetadataParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlMetadataParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlMetadata",
		Method:             "GET",
		PathPattern:        "/access/saml/metadata",
		ProducesMediaTypes: []string{"application/samlmetadata+xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlSingleLogoutHTTPPostConsumer processes a single logout message from the configured s a m l identity provider using the HTTP p o s t binding

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlSingleLogoutHTTPPostConsumer(params *SamlSingleLogoutHTTPPostConsumerParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlSingleLogoutHTTPPostConsumerParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlSingleLogoutHttpPostConsumer",
		Method:             "POST",
		PathPattern:        "/access/saml/single-logout/consumer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlSingleLogoutHTTPPostConsumerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlSingleLogoutHTTPRedirectConsumer processes a single logout message from the configured s a m l identity provider using the HTTP r e d i r e c t binding

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlSingleLogoutHTTPRedirectConsumer(params *SamlSingleLogoutHTTPRedirectConsumerParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlSingleLogoutHTTPRedirectConsumerParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlSingleLogoutHttpRedirectConsumer",
		Method:             "GET",
		PathPattern:        "/access/saml/single-logout/consumer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlSingleLogoutHTTPRedirectConsumerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlSingleLogoutRequest initiates a logout request using the single logout service of the configured s a m l identity provider

  Note: This endpoint is subject to change as NiFi and it's REST API evolve.
*/
func (a *Client) SamlSingleLogoutRequest(params *SamlSingleLogoutRequestParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlSingleLogoutRequestParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "samlSingleLogoutRequest",
		Method:             "GET",
		PathPattern:        "/access/saml/single-logout/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SamlSingleLogoutRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}