// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOutputPortStatusParams creates a new GetOutputPortStatusParams object
// with the default values initialized.
func NewGetOutputPortStatusParams() *GetOutputPortStatusParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetOutputPortStatusParams{
		Nodewise: &nodewiseDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutputPortStatusParamsWithTimeout creates a new GetOutputPortStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutputPortStatusParamsWithTimeout(timeout time.Duration) *GetOutputPortStatusParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetOutputPortStatusParams{
		Nodewise: &nodewiseDefault,

		timeout: timeout,
	}
}

// NewGetOutputPortStatusParamsWithContext creates a new GetOutputPortStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutputPortStatusParamsWithContext(ctx context.Context) *GetOutputPortStatusParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetOutputPortStatusParams{
		Nodewise: &nodewiseDefault,

		Context: ctx,
	}
}

// NewGetOutputPortStatusParamsWithHTTPClient creates a new GetOutputPortStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutputPortStatusParamsWithHTTPClient(client *http.Client) *GetOutputPortStatusParams {
	var (
		nodewiseDefault = bool(false)
	)
	return &GetOutputPortStatusParams{
		Nodewise:   &nodewiseDefault,
		HTTPClient: client,
	}
}

/*GetOutputPortStatusParams contains all the parameters to send to the API endpoint
for the get output port status operation typically these are written to a http.Request
*/
type GetOutputPortStatusParams struct {

	/*ClusterNodeID
	  The id of the node where to get the status.

	*/
	ClusterNodeID *string
	/*ID
	  The output port id.

	*/
	ID string
	/*Nodewise
	  Whether or not to include the breakdown per node. Optional, defaults to false

	*/
	Nodewise *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get output port status params
func (o *GetOutputPortStatusParams) WithTimeout(timeout time.Duration) *GetOutputPortStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get output port status params
func (o *GetOutputPortStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get output port status params
func (o *GetOutputPortStatusParams) WithContext(ctx context.Context) *GetOutputPortStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get output port status params
func (o *GetOutputPortStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get output port status params
func (o *GetOutputPortStatusParams) WithHTTPClient(client *http.Client) *GetOutputPortStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get output port status params
func (o *GetOutputPortStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get output port status params
func (o *GetOutputPortStatusParams) WithClusterNodeID(clusterNodeID *string) *GetOutputPortStatusParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get output port status params
func (o *GetOutputPortStatusParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get output port status params
func (o *GetOutputPortStatusParams) WithID(id string) *GetOutputPortStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get output port status params
func (o *GetOutputPortStatusParams) SetID(id string) {
	o.ID = id
}

// WithNodewise adds the nodewise to the get output port status params
func (o *GetOutputPortStatusParams) WithNodewise(nodewise *bool) *GetOutputPortStatusParams {
	o.SetNodewise(nodewise)
	return o
}

// SetNodewise adds the nodewise to the get output port status params
func (o *GetOutputPortStatusParams) SetNodewise(nodewise *bool) {
	o.Nodewise = nodewise
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutputPortStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Nodewise != nil {

		// query param nodewise
		var qrNodewise bool
		if o.Nodewise != nil {
			qrNodewise = *o.Nodewise
		}
		qNodewise := swag.FormatBool(qrNodewise)
		if qNodewise != "" {
			if err := r.SetQueryParam("nodewise", qNodewise); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}