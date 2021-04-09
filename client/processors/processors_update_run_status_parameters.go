// Code generated by go-swagger; DO NOT EDIT.

package processors

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

// NewProcessorsUpdateRunStatusParams creates a new ProcessorsUpdateRunStatusParams object
// with the default values initialized.
func NewProcessorsUpdateRunStatusParams() *ProcessorsUpdateRunStatusParams {
	var ()
	return &ProcessorsUpdateRunStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProcessorsUpdateRunStatusParamsWithTimeout creates a new ProcessorsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProcessorsUpdateRunStatusParamsWithTimeout(timeout time.Duration) *ProcessorsUpdateRunStatusParams {
	var ()
	return &ProcessorsUpdateRunStatusParams{

		timeout: timeout,
	}
}

// NewProcessorsUpdateRunStatusParamsWithContext creates a new ProcessorsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewProcessorsUpdateRunStatusParamsWithContext(ctx context.Context) *ProcessorsUpdateRunStatusParams {
	var ()
	return &ProcessorsUpdateRunStatusParams{

		Context: ctx,
	}
}

// NewProcessorsUpdateRunStatusParamsWithHTTPClient creates a new ProcessorsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProcessorsUpdateRunStatusParamsWithHTTPClient(client *http.Client) *ProcessorsUpdateRunStatusParams {
	var ()
	return &ProcessorsUpdateRunStatusParams{
		HTTPClient: client,
	}
}

/*ProcessorsUpdateRunStatusParams contains all the parameters to send to the API endpoint
for the processors update run status operation typically these are written to a http.Request
*/
type ProcessorsUpdateRunStatusParams struct {

	/*Body
	  The processor run status.

	*/
	Body *models.ProcessorRunStatusEntity
	/*ID
	  The processor id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) WithTimeout(timeout time.Duration) *ProcessorsUpdateRunStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) WithContext(ctx context.Context) *ProcessorsUpdateRunStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) WithHTTPClient(client *http.Client) *ProcessorsUpdateRunStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) WithBody(body *models.ProcessorRunStatusEntity) *ProcessorsUpdateRunStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) SetBody(body *models.ProcessorRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) WithID(id string) *ProcessorsUpdateRunStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the processors update run status params
func (o *ProcessorsUpdateRunStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessorsUpdateRunStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
