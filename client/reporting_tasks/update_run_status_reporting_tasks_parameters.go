// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

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

// NewUpdateRunStatusReportingTasksParams creates a new UpdateRunStatusReportingTasksParams object
// with the default values initialized.
func NewUpdateRunStatusReportingTasksParams() *UpdateRunStatusReportingTasksParams {
	var ()
	return &UpdateRunStatusReportingTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRunStatusReportingTasksParamsWithTimeout creates a new UpdateRunStatusReportingTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRunStatusReportingTasksParamsWithTimeout(timeout time.Duration) *UpdateRunStatusReportingTasksParams {
	var ()
	return &UpdateRunStatusReportingTasksParams{

		timeout: timeout,
	}
}

// NewUpdateRunStatusReportingTasksParamsWithContext creates a new UpdateRunStatusReportingTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRunStatusReportingTasksParamsWithContext(ctx context.Context) *UpdateRunStatusReportingTasksParams {
	var ()
	return &UpdateRunStatusReportingTasksParams{

		Context: ctx,
	}
}

// NewUpdateRunStatusReportingTasksParamsWithHTTPClient creates a new UpdateRunStatusReportingTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRunStatusReportingTasksParamsWithHTTPClient(client *http.Client) *UpdateRunStatusReportingTasksParams {
	var ()
	return &UpdateRunStatusReportingTasksParams{
		HTTPClient: client,
	}
}

/*UpdateRunStatusReportingTasksParams contains all the parameters to send to the API endpoint
for the update run status reporting tasks operation typically these are written to a http.Request
*/
type UpdateRunStatusReportingTasksParams struct {

	/*Body
	  The reporting task run status.

	*/
	Body *models.ReportingTaskRunStatusEntity
	/*ID
	  The reporting task id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) WithTimeout(timeout time.Duration) *UpdateRunStatusReportingTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) WithContext(ctx context.Context) *UpdateRunStatusReportingTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) WithHTTPClient(client *http.Client) *UpdateRunStatusReportingTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) WithBody(body *models.ReportingTaskRunStatusEntity) *UpdateRunStatusReportingTasksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) SetBody(body *models.ReportingTaskRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) WithID(id string) *UpdateRunStatusReportingTasksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update run status reporting tasks params
func (o *UpdateRunStatusReportingTasksParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRunStatusReportingTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
