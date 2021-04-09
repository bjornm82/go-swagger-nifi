// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

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

// NewUpdateRemoteProcessGroupOutputPortRunStatusParams creates a new UpdateRemoteProcessGroupOutputPortRunStatusParams object
// with the default values initialized.
func NewUpdateRemoteProcessGroupOutputPortRunStatusParams() *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupOutputPortRunStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithTimeout creates a new UpdateRemoteProcessGroupOutputPortRunStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupOutputPortRunStatusParams{

		timeout: timeout,
	}
}

// NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithContext creates a new UpdateRemoteProcessGroupOutputPortRunStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithContext(ctx context.Context) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupOutputPortRunStatusParams{

		Context: ctx,
	}
}

// NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithHTTPClient creates a new UpdateRemoteProcessGroupOutputPortRunStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRemoteProcessGroupOutputPortRunStatusParamsWithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	var ()
	return &UpdateRemoteProcessGroupOutputPortRunStatusParams{
		HTTPClient: client,
	}
}

/*UpdateRemoteProcessGroupOutputPortRunStatusParams contains all the parameters to send to the API endpoint
for the update remote process group output port run status operation typically these are written to a http.Request
*/
type UpdateRemoteProcessGroupOutputPortRunStatusParams struct {

	/*Body
	  The remote process group port.

	*/
	Body *models.RemotePortRunStatusEntity
	/*ID
	  The remote process group id.

	*/
	ID string
	/*PortID
	  The remote process group port id.

	*/
	PortID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithTimeout(timeout time.Duration) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithContext(ctx context.Context) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithHTTPClient(client *http.Client) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithBody(body *models.RemotePortRunStatusEntity) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetBody(body *models.RemotePortRunStatusEntity) {
	o.Body = body
}

// WithID adds the id to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithID(id string) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetID(id string) {
	o.ID = id
}

// WithPortID adds the portID to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WithPortID(portID string) *UpdateRemoteProcessGroupOutputPortRunStatusParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the update remote process group output port run status params
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) SetPortID(portID string) {
	o.PortID = portID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRemoteProcessGroupOutputPortRunStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param port-id
	if err := r.SetPathParam("port-id", o.PortID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}