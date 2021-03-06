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

// NewLogOutCompleteSamlParams creates a new LogOutCompleteSamlParams object
// with the default values initialized.
func NewLogOutCompleteSamlParams() *LogOutCompleteSamlParams {

	return &LogOutCompleteSamlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogOutCompleteSamlParamsWithTimeout creates a new LogOutCompleteSamlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogOutCompleteSamlParamsWithTimeout(timeout time.Duration) *LogOutCompleteSamlParams {

	return &LogOutCompleteSamlParams{

		timeout: timeout,
	}
}

// NewLogOutCompleteSamlParamsWithContext creates a new LogOutCompleteSamlParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogOutCompleteSamlParamsWithContext(ctx context.Context) *LogOutCompleteSamlParams {

	return &LogOutCompleteSamlParams{

		Context: ctx,
	}
}

// NewLogOutCompleteSamlParamsWithHTTPClient creates a new LogOutCompleteSamlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogOutCompleteSamlParamsWithHTTPClient(client *http.Client) *LogOutCompleteSamlParams {

	return &LogOutCompleteSamlParams{
		HTTPClient: client,
	}
}

/*LogOutCompleteSamlParams contains all the parameters to send to the API endpoint
for the log out complete saml operation typically these are written to a http.Request
*/
type LogOutCompleteSamlParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log out complete saml params
func (o *LogOutCompleteSamlParams) WithTimeout(timeout time.Duration) *LogOutCompleteSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log out complete saml params
func (o *LogOutCompleteSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log out complete saml params
func (o *LogOutCompleteSamlParams) WithContext(ctx context.Context) *LogOutCompleteSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log out complete saml params
func (o *LogOutCompleteSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log out complete saml params
func (o *LogOutCompleteSamlParams) WithHTTPClient(client *http.Client) *LogOutCompleteSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log out complete saml params
func (o *LogOutCompleteSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogOutCompleteSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
