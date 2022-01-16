// Code generated by go-swagger; DO NOT EDIT.

package provenance_events

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
)

// NewGetOutputContentParams creates a new GetOutputContentParams object
// with the default values initialized.
func NewGetOutputContentParams() *GetOutputContentParams {
	var ()
	return &GetOutputContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutputContentParamsWithTimeout creates a new GetOutputContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutputContentParamsWithTimeout(timeout time.Duration) *GetOutputContentParams {
	var ()
	return &GetOutputContentParams{

		timeout: timeout,
	}
}

// NewGetOutputContentParamsWithContext creates a new GetOutputContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutputContentParamsWithContext(ctx context.Context) *GetOutputContentParams {
	var ()
	return &GetOutputContentParams{

		Context: ctx,
	}
}

// NewGetOutputContentParamsWithHTTPClient creates a new GetOutputContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutputContentParamsWithHTTPClient(client *http.Client) *GetOutputContentParams {
	var ()
	return &GetOutputContentParams{
		HTTPClient: client,
	}
}

/*GetOutputContentParams contains all the parameters to send to the API endpoint
for the get output content operation typically these are written to a http.Request
*/
type GetOutputContentParams struct {

	/*ClusterNodeID
	  The id of the node where the content exists if clustered.

	*/
	ClusterNodeID *string
	/*ID
	  The provenance event id.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get output content params
func (o *GetOutputContentParams) WithTimeout(timeout time.Duration) *GetOutputContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get output content params
func (o *GetOutputContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get output content params
func (o *GetOutputContentParams) WithContext(ctx context.Context) *GetOutputContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get output content params
func (o *GetOutputContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get output content params
func (o *GetOutputContentParams) WithHTTPClient(client *http.Client) *GetOutputContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get output content params
func (o *GetOutputContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get output content params
func (o *GetOutputContentParams) WithClusterNodeID(clusterNodeID *string) *GetOutputContentParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get output content params
func (o *GetOutputContentParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get output content params
func (o *GetOutputContentParams) WithID(id string) *GetOutputContentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get output content params
func (o *GetOutputContentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterNodeID != nil {

		// query param clusterNodeId
		var qrClusterNodeID string
		if o.ClusterNodeID != nil {
			qrClusterNodeID = *o.ClusterNodeID
		}
		qClusterNodeID := qrClusterNodeID
		if qClusterNodeID != "" {
			if err := r.SetQueryParam("clusterNodeId", qClusterNodeID); err != nil {
				return err
			}
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
