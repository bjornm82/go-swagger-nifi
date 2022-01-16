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

// NewSamlLoginHTTPPostConsumerParams creates a new SamlLoginHTTPPostConsumerParams object
// with the default values initialized.
func NewSamlLoginHTTPPostConsumerParams() *SamlLoginHTTPPostConsumerParams {

	return &SamlLoginHTTPPostConsumerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSamlLoginHTTPPostConsumerParamsWithTimeout creates a new SamlLoginHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSamlLoginHTTPPostConsumerParamsWithTimeout(timeout time.Duration) *SamlLoginHTTPPostConsumerParams {

	return &SamlLoginHTTPPostConsumerParams{

		timeout: timeout,
	}
}

// NewSamlLoginHTTPPostConsumerParamsWithContext creates a new SamlLoginHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a context for a request
func NewSamlLoginHTTPPostConsumerParamsWithContext(ctx context.Context) *SamlLoginHTTPPostConsumerParams {

	return &SamlLoginHTTPPostConsumerParams{

		Context: ctx,
	}
}

// NewSamlLoginHTTPPostConsumerParamsWithHTTPClient creates a new SamlLoginHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSamlLoginHTTPPostConsumerParamsWithHTTPClient(client *http.Client) *SamlLoginHTTPPostConsumerParams {

	return &SamlLoginHTTPPostConsumerParams{
		HTTPClient: client,
	}
}

/*SamlLoginHTTPPostConsumerParams contains all the parameters to send to the API endpoint
for the saml login Http post consumer operation typically these are written to a http.Request
*/
type SamlLoginHTTPPostConsumerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) WithTimeout(timeout time.Duration) *SamlLoginHTTPPostConsumerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) WithContext(ctx context.Context) *SamlLoginHTTPPostConsumerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) WithHTTPClient(client *http.Client) *SamlLoginHTTPPostConsumerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the saml login Http post consumer params
func (o *SamlLoginHTTPPostConsumerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SamlLoginHTTPPostConsumerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
