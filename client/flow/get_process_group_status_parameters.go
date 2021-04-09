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

// NewGetProcessGroupStatusParams creates a new GetProcessGroupStatusParams object
// with the default values initialized.
func NewGetProcessGroupStatusParams() *GetProcessGroupStatusParams {
	var (
		nodewiseDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &GetProcessGroupStatusParams{
		Nodewise:  &nodewiseDefault,
		Recursive: &recursiveDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProcessGroupStatusParamsWithTimeout creates a new GetProcessGroupStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProcessGroupStatusParamsWithTimeout(timeout time.Duration) *GetProcessGroupStatusParams {
	var (
		nodewiseDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &GetProcessGroupStatusParams{
		Nodewise:  &nodewiseDefault,
		Recursive: &recursiveDefault,

		timeout: timeout,
	}
}

// NewGetProcessGroupStatusParamsWithContext creates a new GetProcessGroupStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProcessGroupStatusParamsWithContext(ctx context.Context) *GetProcessGroupStatusParams {
	var (
		nodewiseDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &GetProcessGroupStatusParams{
		Nodewise:  &nodewiseDefault,
		Recursive: &recursiveDefault,

		Context: ctx,
	}
}

// NewGetProcessGroupStatusParamsWithHTTPClient creates a new GetProcessGroupStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProcessGroupStatusParamsWithHTTPClient(client *http.Client) *GetProcessGroupStatusParams {
	var (
		nodewiseDefault  = bool(false)
		recursiveDefault = bool(false)
	)
	return &GetProcessGroupStatusParams{
		Nodewise:   &nodewiseDefault,
		Recursive:  &recursiveDefault,
		HTTPClient: client,
	}
}

/*GetProcessGroupStatusParams contains all the parameters to send to the API endpoint
for the get process group status operation typically these are written to a http.Request
*/
type GetProcessGroupStatusParams struct {

	/*ClusterNodeID
	  The id of the node where to get the status.

	*/
	ClusterNodeID *string
	/*ID
	  The process group id.

	*/
	ID string
	/*Nodewise
	  Whether or not to include the breakdown per node. Optional, defaults to false

	*/
	Nodewise *bool
	/*Recursive
	  Whether all descendant groups and the status of their content will be included. Optional, defaults to false

	*/
	Recursive *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get process group status params
func (o *GetProcessGroupStatusParams) WithTimeout(timeout time.Duration) *GetProcessGroupStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get process group status params
func (o *GetProcessGroupStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get process group status params
func (o *GetProcessGroupStatusParams) WithContext(ctx context.Context) *GetProcessGroupStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get process group status params
func (o *GetProcessGroupStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get process group status params
func (o *GetProcessGroupStatusParams) WithHTTPClient(client *http.Client) *GetProcessGroupStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get process group status params
func (o *GetProcessGroupStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNodeID adds the clusterNodeID to the get process group status params
func (o *GetProcessGroupStatusParams) WithClusterNodeID(clusterNodeID *string) *GetProcessGroupStatusParams {
	o.SetClusterNodeID(clusterNodeID)
	return o
}

// SetClusterNodeID adds the clusterNodeId to the get process group status params
func (o *GetProcessGroupStatusParams) SetClusterNodeID(clusterNodeID *string) {
	o.ClusterNodeID = clusterNodeID
}

// WithID adds the id to the get process group status params
func (o *GetProcessGroupStatusParams) WithID(id string) *GetProcessGroupStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get process group status params
func (o *GetProcessGroupStatusParams) SetID(id string) {
	o.ID = id
}

// WithNodewise adds the nodewise to the get process group status params
func (o *GetProcessGroupStatusParams) WithNodewise(nodewise *bool) *GetProcessGroupStatusParams {
	o.SetNodewise(nodewise)
	return o
}

// SetNodewise adds the nodewise to the get process group status params
func (o *GetProcessGroupStatusParams) SetNodewise(nodewise *bool) {
	o.Nodewise = nodewise
}

// WithRecursive adds the recursive to the get process group status params
func (o *GetProcessGroupStatusParams) WithRecursive(recursive *bool) *GetProcessGroupStatusParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the get process group status params
func (o *GetProcessGroupStatusParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WriteToRequest writes these params to a swagger request
func (o *GetProcessGroupStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}