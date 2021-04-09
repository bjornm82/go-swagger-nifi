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

// NewOidcLogoutCallbackParams creates a new OidcLogoutCallbackParams object
// with the default values initialized.
func NewOidcLogoutCallbackParams() *OidcLogoutCallbackParams {

	return &OidcLogoutCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOidcLogoutCallbackParamsWithTimeout creates a new OidcLogoutCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOidcLogoutCallbackParamsWithTimeout(timeout time.Duration) *OidcLogoutCallbackParams {

	return &OidcLogoutCallbackParams{

		timeout: timeout,
	}
}

// NewOidcLogoutCallbackParamsWithContext creates a new OidcLogoutCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewOidcLogoutCallbackParamsWithContext(ctx context.Context) *OidcLogoutCallbackParams {

	return &OidcLogoutCallbackParams{

		Context: ctx,
	}
}

// NewOidcLogoutCallbackParamsWithHTTPClient creates a new OidcLogoutCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOidcLogoutCallbackParamsWithHTTPClient(client *http.Client) *OidcLogoutCallbackParams {

	return &OidcLogoutCallbackParams{
		HTTPClient: client,
	}
}

/*OidcLogoutCallbackParams contains all the parameters to send to the API endpoint
for the oidc logout callback operation typically these are written to a http.Request
*/
type OidcLogoutCallbackParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oidc logout callback params
func (o *OidcLogoutCallbackParams) WithTimeout(timeout time.Duration) *OidcLogoutCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oidc logout callback params
func (o *OidcLogoutCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oidc logout callback params
func (o *OidcLogoutCallbackParams) WithContext(ctx context.Context) *OidcLogoutCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oidc logout callback params
func (o *OidcLogoutCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oidc logout callback params
func (o *OidcLogoutCallbackParams) WithHTTPClient(client *http.Client) *OidcLogoutCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oidc logout callback params
func (o *OidcLogoutCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OidcLogoutCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}