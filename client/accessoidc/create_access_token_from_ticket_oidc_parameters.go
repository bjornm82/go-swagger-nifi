// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

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

// NewCreateAccessTokenFromTicketOidcParams creates a new CreateAccessTokenFromTicketOidcParams object
// with the default values initialized.
func NewCreateAccessTokenFromTicketOidcParams() *CreateAccessTokenFromTicketOidcParams {

	return &CreateAccessTokenFromTicketOidcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccessTokenFromTicketOidcParamsWithTimeout creates a new CreateAccessTokenFromTicketOidcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccessTokenFromTicketOidcParamsWithTimeout(timeout time.Duration) *CreateAccessTokenFromTicketOidcParams {

	return &CreateAccessTokenFromTicketOidcParams{

		timeout: timeout,
	}
}

// NewCreateAccessTokenFromTicketOidcParamsWithContext creates a new CreateAccessTokenFromTicketOidcParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccessTokenFromTicketOidcParamsWithContext(ctx context.Context) *CreateAccessTokenFromTicketOidcParams {

	return &CreateAccessTokenFromTicketOidcParams{

		Context: ctx,
	}
}

// NewCreateAccessTokenFromTicketOidcParamsWithHTTPClient creates a new CreateAccessTokenFromTicketOidcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccessTokenFromTicketOidcParamsWithHTTPClient(client *http.Client) *CreateAccessTokenFromTicketOidcParams {

	return &CreateAccessTokenFromTicketOidcParams{
		HTTPClient: client,
	}
}

/*CreateAccessTokenFromTicketOidcParams contains all the parameters to send to the API endpoint
for the create access token from ticket oidc operation typically these are written to a http.Request
*/
type CreateAccessTokenFromTicketOidcParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) WithTimeout(timeout time.Duration) *CreateAccessTokenFromTicketOidcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) WithContext(ctx context.Context) *CreateAccessTokenFromTicketOidcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) WithHTTPClient(client *http.Client) *CreateAccessTokenFromTicketOidcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create access token from ticket oidc params
func (o *CreateAccessTokenFromTicketOidcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccessTokenFromTicketOidcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
