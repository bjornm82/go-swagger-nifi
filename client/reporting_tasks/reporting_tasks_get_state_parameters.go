// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

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

// NewReportingTasksGetStateParams creates a new ReportingTasksGetStateParams object
// with the default values initialized.
func NewReportingTasksGetStateParams() *ReportingTasksGetStateParams {
	var ()
	return &ReportingTasksGetStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReportingTasksGetStateParamsWithTimeout creates a new ReportingTasksGetStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReportingTasksGetStateParamsWithTimeout(timeout time.Duration) *ReportingTasksGetStateParams {
	var ()
	return &ReportingTasksGetStateParams{

		timeout: timeout,
	}
}

// NewReportingTasksGetStateParamsWithContext creates a new ReportingTasksGetStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewReportingTasksGetStateParamsWithContext(ctx context.Context) *ReportingTasksGetStateParams {
	var ()
	return &ReportingTasksGetStateParams{

		Context: ctx,
	}
}

// NewReportingTasksGetStateParamsWithHTTPClient creates a new ReportingTasksGetStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReportingTasksGetStateParamsWithHTTPClient(client *http.Client) *ReportingTasksGetStateParams {
	var ()
	return &ReportingTasksGetStateParams{
		HTTPClient: client,
	}
}

/*ReportingTasksGetStateParams contains all the parameters to send to the API endpoint
for the reporting tasks get state operation typically these are written to a http.Request
*/
type ReportingTasksGetStateParams struct {

	/*ID
	  The reporting task id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) WithTimeout(timeout time.Duration) *ReportingTasksGetStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) WithContext(ctx context.Context) *ReportingTasksGetStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) WithHTTPClient(client *http.Client) *ReportingTasksGetStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) WithID(id string) *ReportingTasksGetStateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reporting tasks get state params
func (o *ReportingTasksGetStateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReportingTasksGetStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
