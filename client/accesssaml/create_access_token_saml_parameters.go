// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateAccessTokenSamlParams creates a new CreateAccessTokenSamlParams object
// with the default values initialized.
func NewCreateAccessTokenSamlParams() *CreateAccessTokenSamlParams {
	var ()
	return &CreateAccessTokenSamlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccessTokenSamlParamsWithTimeout creates a new CreateAccessTokenSamlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccessTokenSamlParamsWithTimeout(timeout time.Duration) *CreateAccessTokenSamlParams {
	var ()
	return &CreateAccessTokenSamlParams{

		timeout: timeout,
	}
}

// NewCreateAccessTokenSamlParamsWithContext creates a new CreateAccessTokenSamlParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccessTokenSamlParamsWithContext(ctx context.Context) *CreateAccessTokenSamlParams {
	var ()
	return &CreateAccessTokenSamlParams{

		Context: ctx,
	}
}

// NewCreateAccessTokenSamlParamsWithHTTPClient creates a new CreateAccessTokenSamlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccessTokenSamlParamsWithHTTPClient(client *http.Client) *CreateAccessTokenSamlParams {
	var ()
	return &CreateAccessTokenSamlParams{
		HTTPClient: client,
	}
}

/*CreateAccessTokenSamlParams contains all the parameters to send to the API endpoint
for the create access token saml operation typically these are written to a http.Request
*/
type CreateAccessTokenSamlParams struct {

	/*Password*/
	Password *string
	/*Username*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create access token saml params
func (o *CreateAccessTokenSamlParams) WithTimeout(timeout time.Duration) *CreateAccessTokenSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create access token saml params
func (o *CreateAccessTokenSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create access token saml params
func (o *CreateAccessTokenSamlParams) WithContext(ctx context.Context) *CreateAccessTokenSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create access token saml params
func (o *CreateAccessTokenSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create access token saml params
func (o *CreateAccessTokenSamlParams) WithHTTPClient(client *http.Client) *CreateAccessTokenSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create access token saml params
func (o *CreateAccessTokenSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPassword adds the password to the create access token saml params
func (o *CreateAccessTokenSamlParams) WithPassword(password *string) *CreateAccessTokenSamlParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the create access token saml params
func (o *CreateAccessTokenSamlParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the create access token saml params
func (o *CreateAccessTokenSamlParams) WithUsername(username *string) *CreateAccessTokenSamlParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the create access token saml params
func (o *CreateAccessTokenSamlParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccessTokenSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}

	}

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
