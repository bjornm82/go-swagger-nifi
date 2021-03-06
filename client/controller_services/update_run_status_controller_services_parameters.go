// Code generated by go-swagger; DO NOT EDIT.

package controller_services

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

	"github.com/bjornm82/go-swagger-nifi/models"
)

// NewUpdateRunStatusControllerServicesParams creates a new UpdateRunStatusControllerServicesParams object
// with the default values initialized.
func NewUpdateRunStatusControllerServicesParams() *UpdateRunStatusControllerServicesParams {
	var ()
	return &UpdateRunStatusControllerServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRunStatusControllerServicesParamsWithTimeout creates a new UpdateRunStatusControllerServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRunStatusControllerServicesParamsWithTimeout(timeout time.Duration) *UpdateRunStatusControllerServicesParams {
	var ()
	return &UpdateRunStatusControllerServicesParams{

		timeout: timeout,
	}
}

// NewUpdateRunStatusControllerServicesParamsWithContext creates a new UpdateRunStatusControllerServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRunStatusControllerServicesParamsWithContext(ctx context.Context) *UpdateRunStatusControllerServicesParams {
	var ()
	return &UpdateRunStatusControllerServicesParams{

		Context: ctx,
	}
}

// NewUpdateRunStatusControllerServicesParamsWithHTTPClient creates a new UpdateRunStatusControllerServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRunStatusControllerServicesParamsWithHTTPClient(client *http.Client) *UpdateRunStatusControllerServicesParams {
	var ()
	return &UpdateRunStatusControllerServicesParams{
		HTTPClient: client,
	}
}

/*UpdateRunStatusControllerServicesParams contains all the parameters to send to the API endpoint
for the update run status controller services operation typically these are written to a http.Request
*/
type UpdateRunStatusControllerServicesParams struct {

	/*Body
	  The controller service run status.

	*/
	Body *models.ControllerServiceRunStatusEntity
	/*ID
	  The controller service id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) WithTimeout(timeout time.Duration) *UpdateRunStatusControllerServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) WithContext(ctx context.Context) *UpdateRunStatusControllerServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) WithHTTPClient(client *http.Client) *UpdateRunStatusControllerServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) WithBody(body *models.ControllerServiceRunStatusEntity) *UpdateRunStatusControllerServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) SetBody(body *models.ControllerServiceRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) WithID(id string) *UpdateRunStatusControllerServicesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update run status controller services params
func (o *UpdateRunStatusControllerServicesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRunStatusControllerServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
