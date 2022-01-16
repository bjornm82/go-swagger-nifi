// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

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

// NewUpdateParameterContextParams creates a new UpdateParameterContextParams object
// with the default values initialized.
func NewUpdateParameterContextParams() *UpdateParameterContextParams {
	var ()
	return &UpdateParameterContextParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParameterContextParamsWithTimeout creates a new UpdateParameterContextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateParameterContextParamsWithTimeout(timeout time.Duration) *UpdateParameterContextParams {
	var ()
	return &UpdateParameterContextParams{

		timeout: timeout,
	}
}

// NewUpdateParameterContextParamsWithContext creates a new UpdateParameterContextParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateParameterContextParamsWithContext(ctx context.Context) *UpdateParameterContextParams {
	var ()
	return &UpdateParameterContextParams{

		Context: ctx,
	}
}

// NewUpdateParameterContextParamsWithHTTPClient creates a new UpdateParameterContextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateParameterContextParamsWithHTTPClient(client *http.Client) *UpdateParameterContextParams {
	var ()
	return &UpdateParameterContextParams{
		HTTPClient: client,
	}
}

/*UpdateParameterContextParams contains all the parameters to send to the API endpoint
for the update parameter context operation typically these are written to a http.Request
*/
type UpdateParameterContextParams struct {

	/*Body
	  The updated Parameter Context

	*/
	Body *models.ParameterContextEntity
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update parameter context params
func (o *UpdateParameterContextParams) WithTimeout(timeout time.Duration) *UpdateParameterContextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update parameter context params
func (o *UpdateParameterContextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update parameter context params
func (o *UpdateParameterContextParams) WithContext(ctx context.Context) *UpdateParameterContextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update parameter context params
func (o *UpdateParameterContextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update parameter context params
func (o *UpdateParameterContextParams) WithHTTPClient(client *http.Client) *UpdateParameterContextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update parameter context params
func (o *UpdateParameterContextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update parameter context params
func (o *UpdateParameterContextParams) WithBody(body *models.ParameterContextEntity) *UpdateParameterContextParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update parameter context params
func (o *UpdateParameterContextParams) SetBody(body *models.ParameterContextEntity) {
	o.Body = body
}

// WithID adds the id to the update parameter context params
func (o *UpdateParameterContextParams) WithID(id string) *UpdateParameterContextParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update parameter context params
func (o *UpdateParameterContextParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParameterContextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
