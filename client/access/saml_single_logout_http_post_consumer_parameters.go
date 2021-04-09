// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSamlSingleLogoutHTTPPostConsumerParams creates a new SamlSingleLogoutHTTPPostConsumerParams object
// with the default values initialized.
func NewSamlSingleLogoutHTTPPostConsumerParams() *SamlSingleLogoutHTTPPostConsumerParams {

	return &SamlSingleLogoutHTTPPostConsumerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSamlSingleLogoutHTTPPostConsumerParamsWithTimeout creates a new SamlSingleLogoutHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSamlSingleLogoutHTTPPostConsumerParamsWithTimeout(timeout time.Duration) *SamlSingleLogoutHTTPPostConsumerParams {

	return &SamlSingleLogoutHTTPPostConsumerParams{

		timeout: timeout,
	}
}

// NewSamlSingleLogoutHTTPPostConsumerParamsWithContext creates a new SamlSingleLogoutHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a context for a request
func NewSamlSingleLogoutHTTPPostConsumerParamsWithContext(ctx context.Context) *SamlSingleLogoutHTTPPostConsumerParams {

	return &SamlSingleLogoutHTTPPostConsumerParams{

		Context: ctx,
	}
}

// NewSamlSingleLogoutHTTPPostConsumerParamsWithHTTPClient creates a new SamlSingleLogoutHTTPPostConsumerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSamlSingleLogoutHTTPPostConsumerParamsWithHTTPClient(client *http.Client) *SamlSingleLogoutHTTPPostConsumerParams {

	return &SamlSingleLogoutHTTPPostConsumerParams{
		HTTPClient: client,
	}
}

/*SamlSingleLogoutHTTPPostConsumerParams contains all the parameters to send to the API endpoint
for the saml single logout Http post consumer operation typically these are written to a http.Request
*/
type SamlSingleLogoutHTTPPostConsumerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) WithTimeout(timeout time.Duration) *SamlSingleLogoutHTTPPostConsumerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) WithContext(ctx context.Context) *SamlSingleLogoutHTTPPostConsumerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) WithHTTPClient(client *http.Client) *SamlSingleLogoutHTTPPostConsumerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the saml single logout Http post consumer params
func (o *SamlSingleLogoutHTTPPostConsumerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SamlSingleLogoutHTTPPostConsumerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
