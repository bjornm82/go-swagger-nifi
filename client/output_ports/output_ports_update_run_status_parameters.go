// Code generated by go-swagger; DO NOT EDIT.

package output_ports

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

// NewOutputPortsUpdateRunStatusParams creates a new OutputPortsUpdateRunStatusParams object
// with the default values initialized.
func NewOutputPortsUpdateRunStatusParams() *OutputPortsUpdateRunStatusParams {
	var ()
	return &OutputPortsUpdateRunStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOutputPortsUpdateRunStatusParamsWithTimeout creates a new OutputPortsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOutputPortsUpdateRunStatusParamsWithTimeout(timeout time.Duration) *OutputPortsUpdateRunStatusParams {
	var ()
	return &OutputPortsUpdateRunStatusParams{

		timeout: timeout,
	}
}

// NewOutputPortsUpdateRunStatusParamsWithContext creates a new OutputPortsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewOutputPortsUpdateRunStatusParamsWithContext(ctx context.Context) *OutputPortsUpdateRunStatusParams {
	var ()
	return &OutputPortsUpdateRunStatusParams{

		Context: ctx,
	}
}

// NewOutputPortsUpdateRunStatusParamsWithHTTPClient creates a new OutputPortsUpdateRunStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOutputPortsUpdateRunStatusParamsWithHTTPClient(client *http.Client) *OutputPortsUpdateRunStatusParams {
	var ()
	return &OutputPortsUpdateRunStatusParams{
		HTTPClient: client,
	}
}

/*OutputPortsUpdateRunStatusParams contains all the parameters to send to the API endpoint
for the output ports update run status operation typically these are written to a http.Request
*/
type OutputPortsUpdateRunStatusParams struct {

	/*Body
	  The port run status.

	*/
	Body *models.PortRunStatusEntity
	/*ID
	  The port id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) WithTimeout(timeout time.Duration) *OutputPortsUpdateRunStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) WithContext(ctx context.Context) *OutputPortsUpdateRunStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) WithHTTPClient(client *http.Client) *OutputPortsUpdateRunStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) WithBody(body *models.PortRunStatusEntity) *OutputPortsUpdateRunStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) SetBody(body *models.PortRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) WithID(id string) *OutputPortsUpdateRunStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the output ports update run status params
func (o *OutputPortsUpdateRunStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OutputPortsUpdateRunStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
