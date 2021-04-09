// Code generated by go-swagger; DO NOT EDIT.

package provenance_events

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

// NewGetInputContentParams creates a new GetInputContentParams object
// with the default values initialized.
func NewGetInputContentParams() *GetInputContentParams {
	var ()
	return &GetInputContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInputContentParamsWithTimeout creates a new GetInputContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInputContentParamsWithTimeout(timeout time.Duration) *GetInputContentParams {
	var ()
	return &GetInputContentParams{

		timeout: timeout,
	}
}

// NewGetInputContentParamsWithContext creates a new GetInputContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInputContentParamsWithContext(ctx context.Context) *GetInputContentParams {
	var ()
	return &GetInputContentParams{

		Context: ctx,
	}
}

// NewGetInputContentParamsWithHTTPClient creates a new GetInputContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInputContentParamsWithHTTPClient(client *http.Client) *GetInputContentParams {
	var ()
	return &GetInputContentParams{
		HTTPClient: client,
	}
}

/*GetInputContentParams contains all the parameters to send to the API endpoint
for the get input content operation typically these are written to a http.Request
*/
type GetInputContentParams struct {

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

// WithTimeout adds the timeout to the get input content params
func (o *GetInputContentParams) WithTimeout(timeout time.Duration) *GetInputContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get input content params
func (o *GetInputContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get input content params
func (o *GetInputContentParams) WithContext(ctx context.Context) *GetInputContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get input content params
func (o *GetInputContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get input content params
func (o *GetInputContentParams) WithHTTPClient(client *http.Client) *GetInputContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get input content params
func (o *GetInputContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get input content params
func (o *GetInputContentParams) WithClusterNodeID(clusterNodeID *string) *GetInputContentParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get input content params
func (o *GetInputContentParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get input content params
func (o *GetInputContentParams) WithID(id string) *GetInputContentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get input content params
func (o *GetInputContentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInputContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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