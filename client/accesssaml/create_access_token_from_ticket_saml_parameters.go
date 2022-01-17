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

// NewCreateAccessTokenFromTicketSamlParams creates a new CreateAccessTokenFromTicketSamlParams object
// with the default values initialized.
func NewCreateAccessTokenFromTicketSamlParams() *CreateAccessTokenFromTicketSamlParams {

	return &CreateAccessTokenFromTicketSamlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccessTokenFromTicketSamlParamsWithTimeout creates a new CreateAccessTokenFromTicketSamlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccessTokenFromTicketSamlParamsWithTimeout(timeout time.Duration) *CreateAccessTokenFromTicketSamlParams {

	return &CreateAccessTokenFromTicketSamlParams{

		timeout: timeout,
	}
}

// NewCreateAccessTokenFromTicketSamlParamsWithContext creates a new CreateAccessTokenFromTicketSamlParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccessTokenFromTicketSamlParamsWithContext(ctx context.Context) *CreateAccessTokenFromTicketSamlParams {

	return &CreateAccessTokenFromTicketSamlParams{

		Context: ctx,
	}
}

// NewCreateAccessTokenFromTicketSamlParamsWithHTTPClient creates a new CreateAccessTokenFromTicketSamlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccessTokenFromTicketSamlParamsWithHTTPClient(client *http.Client) *CreateAccessTokenFromTicketSamlParams {

	return &CreateAccessTokenFromTicketSamlParams{
		HTTPClient: client,
	}
}

/*CreateAccessTokenFromTicketSamlParams contains all the parameters to send to the API endpoint
for the create access token from ticket saml operation typically these are written to a http.Request
*/
type CreateAccessTokenFromTicketSamlParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) WithTimeout(timeout time.Duration) *CreateAccessTokenFromTicketSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) WithContext(ctx context.Context) *CreateAccessTokenFromTicketSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) WithHTTPClient(client *http.Client) *CreateAccessTokenFromTicketSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create access token from ticket saml params
func (o *CreateAccessTokenFromTicketSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccessTokenFromTicketSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}