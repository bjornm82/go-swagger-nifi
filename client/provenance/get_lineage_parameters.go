// Code generated by go-swagger; DO NOT EDIT.

package provenance

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

// NewGetLineageParams creates a new GetLineageParams object
// with the default values initialized.
func NewGetLineageParams() *GetLineageParams {
	var ()
	return &GetLineageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLineageParamsWithTimeout creates a new GetLineageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLineageParamsWithTimeout(timeout time.Duration) *GetLineageParams {
	var ()
	return &GetLineageParams{

		timeout: timeout,
	}
}

// NewGetLineageParamsWithContext creates a new GetLineageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLineageParamsWithContext(ctx context.Context) *GetLineageParams {
	var ()
	return &GetLineageParams{

		Context: ctx,
	}
}

// NewGetLineageParamsWithHTTPClient creates a new GetLineageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLineageParamsWithHTTPClient(client *http.Client) *GetLineageParams {
	var ()
	return &GetLineageParams{
		HTTPClient: client,
	}
}

/*GetLineageParams contains all the parameters to send to the API endpoint
for the get lineage operation typically these are written to a http.Request
*/
type GetLineageParams struct {

	/*ClusterNodeID
	  The id of the node where this query exists if clustered.

	*/
	ClusterNodeID *string
	/*ID
	  The id of the lineage query.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lineage params
func (o *GetLineageParams) WithTimeout(timeout time.Duration) *GetLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lineage params
func (o *GetLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lineage params
func (o *GetLineageParams) WithContext(ctx context.Context) *GetLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lineage params
func (o *GetLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lineage params
func (o *GetLineageParams) WithHTTPClient(client *http.Client) *GetLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lineage params
func (o *GetLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get lineage params
func (o *GetLineageParams) WithClusterNodeID(clusterNodeID *string) *GetLineageParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get lineage params
func (o *GetLineageParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get lineage params
func (o *GetLineageParams) WithID(id string) *GetLineageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lineage params
func (o *GetLineageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
