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

// NewKnoxRequestOidcParams creates a new KnoxRequestOidcParams object
// with the default values initialized.
func NewKnoxRequestOidcParams() *KnoxRequestOidcParams {

	return &KnoxRequestOidcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKnoxRequestOidcParamsWithTimeout creates a new KnoxRequestOidcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKnoxRequestOidcParamsWithTimeout(timeout time.Duration) *KnoxRequestOidcParams {

	return &KnoxRequestOidcParams{

		timeout: timeout,
	}
}

// NewKnoxRequestOidcParamsWithContext creates a new KnoxRequestOidcParams object
// with the default values initialized, and the ability to set a context for a request
func NewKnoxRequestOidcParamsWithContext(ctx context.Context) *KnoxRequestOidcParams {

	return &KnoxRequestOidcParams{

		Context: ctx,
	}
}

// NewKnoxRequestOidcParamsWithHTTPClient creates a new KnoxRequestOidcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKnoxRequestOidcParamsWithHTTPClient(client *http.Client) *KnoxRequestOidcParams {

	return &KnoxRequestOidcParams{
		HTTPClient: client,
	}
}

/*KnoxRequestOidcParams contains all the parameters to send to the API endpoint
for the knox request oidc operation typically these are written to a http.Request
*/
type KnoxRequestOidcParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the knox request oidc params
func (o *KnoxRequestOidcParams) WithTimeout(timeout time.Duration) *KnoxRequestOidcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the knox request oidc params
func (o *KnoxRequestOidcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the knox request oidc params
func (o *KnoxRequestOidcParams) WithContext(ctx context.Context) *KnoxRequestOidcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the knox request oidc params
func (o *KnoxRequestOidcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the knox request oidc params
func (o *KnoxRequestOidcParams) WithHTTPClient(client *http.Client) *KnoxRequestOidcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the knox request oidc params
func (o *KnoxRequestOidcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *KnoxRequestOidcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
