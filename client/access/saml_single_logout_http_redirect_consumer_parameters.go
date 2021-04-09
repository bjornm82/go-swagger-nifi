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

// NewSamlSingleLogoutHTTPRedirectConsumerParams creates a new SamlSingleLogoutHTTPRedirectConsumerParams object
// with the default values initialized.
func NewSamlSingleLogoutHTTPRedirectConsumerParams() *SamlSingleLogoutHTTPRedirectConsumerParams {

	return &SamlSingleLogoutHTTPRedirectConsumerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSamlSingleLogoutHTTPRedirectConsumerParamsWithTimeout creates a new SamlSingleLogoutHTTPRedirectConsumerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSamlSingleLogoutHTTPRedirectConsumerParamsWithTimeout(timeout time.Duration) *SamlSingleLogoutHTTPRedirectConsumerParams {

	return &SamlSingleLogoutHTTPRedirectConsumerParams{

		timeout: timeout,
	}
}

// NewSamlSingleLogoutHTTPRedirectConsumerParamsWithContext creates a new SamlSingleLogoutHTTPRedirectConsumerParams object
// with the default values initialized, and the ability to set a context for a request
func NewSamlSingleLogoutHTTPRedirectConsumerParamsWithContext(ctx context.Context) *SamlSingleLogoutHTTPRedirectConsumerParams {

	return &SamlSingleLogoutHTTPRedirectConsumerParams{

		Context: ctx,
	}
}

// NewSamlSingleLogoutHTTPRedirectConsumerParamsWithHTTPClient creates a new SamlSingleLogoutHTTPRedirectConsumerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSamlSingleLogoutHTTPRedirectConsumerParamsWithHTTPClient(client *http.Client) *SamlSingleLogoutHTTPRedirectConsumerParams {

	return &SamlSingleLogoutHTTPRedirectConsumerParams{
		HTTPClient: client,
	}
}

/*SamlSingleLogoutHTTPRedirectConsumerParams contains all the parameters to send to the API endpoint
for the saml single logout Http redirect consumer operation typically these are written to a http.Request
*/
type SamlSingleLogoutHTTPRedirectConsumerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) WithTimeout(timeout time.Duration) *SamlSingleLogoutHTTPRedirectConsumerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) WithContext(ctx context.Context) *SamlSingleLogoutHTTPRedirectConsumerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) WithHTTPClient(client *http.Client) *SamlSingleLogoutHTTPRedirectConsumerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the saml single logout Http redirect consumer params
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SamlSingleLogoutHTTPRedirectConsumerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}