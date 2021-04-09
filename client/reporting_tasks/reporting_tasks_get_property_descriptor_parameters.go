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

// NewReportingTasksGetPropertyDescriptorParams creates a new ReportingTasksGetPropertyDescriptorParams object
// with the default values initialized.
func NewReportingTasksGetPropertyDescriptorParams() *ReportingTasksGetPropertyDescriptorParams {
	var ()
	return &ReportingTasksGetPropertyDescriptorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReportingTasksGetPropertyDescriptorParamsWithTimeout creates a new ReportingTasksGetPropertyDescriptorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReportingTasksGetPropertyDescriptorParamsWithTimeout(timeout time.Duration) *ReportingTasksGetPropertyDescriptorParams {
	var ()
	return &ReportingTasksGetPropertyDescriptorParams{

		timeout: timeout,
	}
}

// NewReportingTasksGetPropertyDescriptorParamsWithContext creates a new ReportingTasksGetPropertyDescriptorParams object
// with the default values initialized, and the ability to set a context for a request
func NewReportingTasksGetPropertyDescriptorParamsWithContext(ctx context.Context) *ReportingTasksGetPropertyDescriptorParams {
	var ()
	return &ReportingTasksGetPropertyDescriptorParams{

		Context: ctx,
	}
}

// NewReportingTasksGetPropertyDescriptorParamsWithHTTPClient creates a new ReportingTasksGetPropertyDescriptorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReportingTasksGetPropertyDescriptorParamsWithHTTPClient(client *http.Client) *ReportingTasksGetPropertyDescriptorParams {
	var ()
	return &ReportingTasksGetPropertyDescriptorParams{
		HTTPClient: client,
	}
}

/*ReportingTasksGetPropertyDescriptorParams contains all the parameters to send to the API endpoint
for the reporting tasks get property descriptor operation typically these are written to a http.Request
*/
type ReportingTasksGetPropertyDescriptorParams struct {

	/*ID
	  The reporting task id.

	*/
	ID string
	/*PropertyName
	  The property name.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) WithTimeout(timeout time.Duration) *ReportingTasksGetPropertyDescriptorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) WithContext(ctx context.Context) *ReportingTasksGetPropertyDescriptorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) WithHTTPClient(client *http.Client) *ReportingTasksGetPropertyDescriptorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) WithID(id string) *ReportingTasksGetPropertyDescriptorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) SetID(id string) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) WithPropertyName(propertyName string) *ReportingTasksGetPropertyDescriptorParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the reporting tasks get property descriptor params
func (o *ReportingTasksGetPropertyDescriptorParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *ReportingTasksGetPropertyDescriptorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param propertyName
	qrPropertyName := o.PropertyName
	qPropertyName := qrPropertyName
	if qPropertyName != "" {
		if err := r.SetQueryParam("propertyName", qPropertyName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}