// Code generated by go-swagger; DO NOT EDIT.

package funnel

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

// NewUpdateFunnelParams creates a new UpdateFunnelParams object
// with the default values initialized.
func NewUpdateFunnelParams() *UpdateFunnelParams {
	var ()
	return &UpdateFunnelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFunnelParamsWithTimeout creates a new UpdateFunnelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFunnelParamsWithTimeout(timeout time.Duration) *UpdateFunnelParams {
	var ()
	return &UpdateFunnelParams{

		timeout: timeout,
	}
}

// NewUpdateFunnelParamsWithContext creates a new UpdateFunnelParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFunnelParamsWithContext(ctx context.Context) *UpdateFunnelParams {
	var ()
	return &UpdateFunnelParams{

		Context: ctx,
	}
}

// NewUpdateFunnelParamsWithHTTPClient creates a new UpdateFunnelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFunnelParamsWithHTTPClient(client *http.Client) *UpdateFunnelParams {
	var ()
	return &UpdateFunnelParams{
		HTTPClient: client,
	}
}

/*UpdateFunnelParams contains all the parameters to send to the API endpoint
for the update funnel operation typically these are written to a http.Request
*/
type UpdateFunnelParams struct {

	/*Body
	  The funnel configuration details.

	*/
	Body *models.FunnelEntity
	/*ID
	  The funnel id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update funnel params
func (o *UpdateFunnelParams) WithTimeout(timeout time.Duration) *UpdateFunnelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update funnel params
func (o *UpdateFunnelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update funnel params
func (o *UpdateFunnelParams) WithContext(ctx context.Context) *UpdateFunnelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update funnel params
func (o *UpdateFunnelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update funnel params
func (o *UpdateFunnelParams) WithHTTPClient(client *http.Client) *UpdateFunnelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update funnel params
func (o *UpdateFunnelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update funnel params
func (o *UpdateFunnelParams) WithBody(body *models.FunnelEntity) *UpdateFunnelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update funnel params
func (o *UpdateFunnelParams) SetBody(body *models.FunnelEntity) {
	o.Body = body
}

// WithID adds the id to the update funnel params
func (o *UpdateFunnelParams) WithID(id string) *UpdateFunnelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update funnel params
func (o *UpdateFunnelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFunnelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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