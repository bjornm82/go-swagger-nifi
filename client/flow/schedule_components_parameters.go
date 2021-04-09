// Code generated by go-swagger; DO NOT EDIT.

package flow

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

	models "github.com/bjornm82/swagger-nifi/models"
)

// NewScheduleComponentsParams creates a new ScheduleComponentsParams object
// with the default values initialized.
func NewScheduleComponentsParams() *ScheduleComponentsParams {
	var ()
	return &ScheduleComponentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleComponentsParamsWithTimeout creates a new ScheduleComponentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScheduleComponentsParamsWithTimeout(timeout time.Duration) *ScheduleComponentsParams {
	var ()
	return &ScheduleComponentsParams{

		timeout: timeout,
	}
}

// NewScheduleComponentsParamsWithContext creates a new ScheduleComponentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewScheduleComponentsParamsWithContext(ctx context.Context) *ScheduleComponentsParams {
	var ()
	return &ScheduleComponentsParams{

		Context: ctx,
	}
}

// NewScheduleComponentsParamsWithHTTPClient creates a new ScheduleComponentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScheduleComponentsParamsWithHTTPClient(client *http.Client) *ScheduleComponentsParams {
	var ()
	return &ScheduleComponentsParams{
		HTTPClient: client,
	}
}

/*ScheduleComponentsParams contains all the parameters to send to the API endpoint
for the schedule components operation typically these are written to a http.Request
*/
type ScheduleComponentsParams struct {

	/*Body
	  The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered.

	*/
	Body *models.ScheduleComponentsEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schedule components params
func (o *ScheduleComponentsParams) WithTimeout(timeout time.Duration) *ScheduleComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule components params
func (o *ScheduleComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule components params
func (o *ScheduleComponentsParams) WithContext(ctx context.Context) *ScheduleComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule components params
func (o *ScheduleComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule components params
func (o *ScheduleComponentsParams) WithHTTPClient(client *http.Client) *ScheduleComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule components params
func (o *ScheduleComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schedule components params
func (o *ScheduleComponentsParams) WithBody(body *models.ScheduleComponentsEntity) *ScheduleComponentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schedule components params
func (o *ScheduleComponentsParams) SetBody(body *models.ScheduleComponentsEntity) {
	o.Body = body
}

// WithID adds the id to the schedule components params
func (o *ScheduleComponentsParams) WithID(id string) *ScheduleComponentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schedule components params
func (o *ScheduleComponentsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
