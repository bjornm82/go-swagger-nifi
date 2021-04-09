// Code generated by go-swagger; DO NOT EDIT.

package controller

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

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// NewUpdateControllerConfigParams creates a new UpdateControllerConfigParams object
// with the default values initialized.
func NewUpdateControllerConfigParams() *UpdateControllerConfigParams {
	var ()
	return &UpdateControllerConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateControllerConfigParamsWithTimeout creates a new UpdateControllerConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateControllerConfigParamsWithTimeout(timeout time.Duration) *UpdateControllerConfigParams {
	var ()
	return &UpdateControllerConfigParams{

		timeout: timeout,
	}
}

// NewUpdateControllerConfigParamsWithContext creates a new UpdateControllerConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateControllerConfigParamsWithContext(ctx context.Context) *UpdateControllerConfigParams {
	var ()
	return &UpdateControllerConfigParams{

		Context: ctx,
	}
}

// NewUpdateControllerConfigParamsWithHTTPClient creates a new UpdateControllerConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateControllerConfigParamsWithHTTPClient(client *http.Client) *UpdateControllerConfigParams {
	var ()
	return &UpdateControllerConfigParams{
		HTTPClient: client,
	}
}

/*UpdateControllerConfigParams contains all the parameters to send to the API endpoint
for the update controller config operation typically these are written to a http.Request
*/
type UpdateControllerConfigParams struct {

	/*Body
	  The controller configuration.

	*/
	Body *models.ControllerConfigurationEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update controller config params
func (o *UpdateControllerConfigParams) WithTimeout(timeout time.Duration) *UpdateControllerConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update controller config params
func (o *UpdateControllerConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update controller config params
func (o *UpdateControllerConfigParams) WithContext(ctx context.Context) *UpdateControllerConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update controller config params
func (o *UpdateControllerConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update controller config params
func (o *UpdateControllerConfigParams) WithHTTPClient(client *http.Client) *UpdateControllerConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update controller config params
func (o *UpdateControllerConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update controller config params
func (o *UpdateControllerConfigParams) WithBody(body *models.ControllerConfigurationEntity) *UpdateControllerConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update controller config params
func (o *UpdateControllerConfigParams) SetBody(body *models.ControllerConfigurationEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateControllerConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
