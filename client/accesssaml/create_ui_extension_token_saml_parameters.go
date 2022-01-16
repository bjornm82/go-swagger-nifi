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

// NewCreateUIExtensionTokenSamlParams creates a new CreateUIExtensionTokenSamlParams object
// with the default values initialized.
func NewCreateUIExtensionTokenSamlParams() *CreateUIExtensionTokenSamlParams {

	return &CreateUIExtensionTokenSamlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUIExtensionTokenSamlParamsWithTimeout creates a new CreateUIExtensionTokenSamlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUIExtensionTokenSamlParamsWithTimeout(timeout time.Duration) *CreateUIExtensionTokenSamlParams {

	return &CreateUIExtensionTokenSamlParams{

		timeout: timeout,
	}
}

// NewCreateUIExtensionTokenSamlParamsWithContext creates a new CreateUIExtensionTokenSamlParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUIExtensionTokenSamlParamsWithContext(ctx context.Context) *CreateUIExtensionTokenSamlParams {

	return &CreateUIExtensionTokenSamlParams{

		Context: ctx,
	}
}

// NewCreateUIExtensionTokenSamlParamsWithHTTPClient creates a new CreateUIExtensionTokenSamlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUIExtensionTokenSamlParamsWithHTTPClient(client *http.Client) *CreateUIExtensionTokenSamlParams {

	return &CreateUIExtensionTokenSamlParams{
		HTTPClient: client,
	}
}

/*CreateUIExtensionTokenSamlParams contains all the parameters to send to the API endpoint
for the create Ui extension token saml operation typically these are written to a http.Request
*/
type CreateUIExtensionTokenSamlParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) WithTimeout(timeout time.Duration) *CreateUIExtensionTokenSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) WithContext(ctx context.Context) *CreateUIExtensionTokenSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) WithHTTPClient(client *http.Client) *CreateUIExtensionTokenSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Ui extension token saml params
func (o *CreateUIExtensionTokenSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUIExtensionTokenSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
