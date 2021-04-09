// Code generated by go-swagger; DO NOT EDIT.

package process_groups

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

// NewProcessGroupsCreateControllerServiceParams creates a new ProcessGroupsCreateControllerServiceParams object
// with the default values initialized.
func NewProcessGroupsCreateControllerServiceParams() *ProcessGroupsCreateControllerServiceParams {
	var ()
	return &ProcessGroupsCreateControllerServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProcessGroupsCreateControllerServiceParamsWithTimeout creates a new ProcessGroupsCreateControllerServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProcessGroupsCreateControllerServiceParamsWithTimeout(timeout time.Duration) *ProcessGroupsCreateControllerServiceParams {
	var ()
	return &ProcessGroupsCreateControllerServiceParams{

		timeout: timeout,
	}
}

// NewProcessGroupsCreateControllerServiceParamsWithContext creates a new ProcessGroupsCreateControllerServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewProcessGroupsCreateControllerServiceParamsWithContext(ctx context.Context) *ProcessGroupsCreateControllerServiceParams {
	var ()
	return &ProcessGroupsCreateControllerServiceParams{

		Context: ctx,
	}
}

// NewProcessGroupsCreateControllerServiceParamsWithHTTPClient creates a new ProcessGroupsCreateControllerServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProcessGroupsCreateControllerServiceParamsWithHTTPClient(client *http.Client) *ProcessGroupsCreateControllerServiceParams {
	var ()
	return &ProcessGroupsCreateControllerServiceParams{
		HTTPClient: client,
	}
}

/*ProcessGroupsCreateControllerServiceParams contains all the parameters to send to the API endpoint
for the process groups create controller service operation typically these are written to a http.Request
*/
type ProcessGroupsCreateControllerServiceParams struct {

	/*Body
	  The controller service configuration details.

	*/
	Body *models.ControllerServiceEntity
	/*ID
	  The process group id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) WithTimeout(timeout time.Duration) *ProcessGroupsCreateControllerServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) WithContext(ctx context.Context) *ProcessGroupsCreateControllerServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) WithHTTPClient(client *http.Client) *ProcessGroupsCreateControllerServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) WithBody(body *models.ControllerServiceEntity) *ProcessGroupsCreateControllerServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) SetBody(body *models.ControllerServiceEntity) {
	o.Body = body
}

// WithID adds the id to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) WithID(id string) *ProcessGroupsCreateControllerServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the process groups create controller service params
func (o *ProcessGroupsCreateControllerServiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessGroupsCreateControllerServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
