// Code generated by go-swagger; DO NOT EDIT.

package processors

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
	"github.com/go-openapi/swag"
)

// NewDeleteProcessorParams creates a new DeleteProcessorParams object
// with the default values initialized.
func NewDeleteProcessorParams() *DeleteProcessorParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteProcessorParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProcessorParamsWithTimeout creates a new DeleteProcessorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProcessorParamsWithTimeout(timeout time.Duration) *DeleteProcessorParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteProcessorParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: timeout,
	}
}

// NewDeleteProcessorParamsWithContext creates a new DeleteProcessorParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProcessorParamsWithContext(ctx context.Context) *DeleteProcessorParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteProcessorParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		Context: ctx,
	}
}

// NewDeleteProcessorParamsWithHTTPClient creates a new DeleteProcessorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProcessorParamsWithHTTPClient(client *http.Client) *DeleteProcessorParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteProcessorParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
		HTTPClient:                   client,
	}
}

/*DeleteProcessorParams contains all the parameters to send to the API endpoint
for the delete processor operation typically these are written to a http.Request
*/
type DeleteProcessorParams struct {

	/*ClientID
	  If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.

	*/
	ClientID *string
	/*DisconnectedNodeAcknowledged
	  Acknowledges that this node is disconnected to allow for mutable requests to proceed.

	*/
	DisconnectedNodeAcknowledged *bool
	/*ID
	  The processor id.

	*/
	ID string
	/*Version
	  The revision is used to verify the client is working with the latest version of the flow.

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete processor params
func (o *DeleteProcessorParams) WithTimeout(timeout time.Duration) *DeleteProcessorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete processor params
func (o *DeleteProcessorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete processor params
func (o *DeleteProcessorParams) WithContext(ctx context.Context) *DeleteProcessorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete processor params
func (o *DeleteProcessorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete processor params
func (o *DeleteProcessorParams) WithHTTPClient(client *http.Client) *DeleteProcessorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete processor params
func (o *DeleteProcessorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete processor params
func (o *DeleteProcessorParams) WithClientID(clientID *string) *DeleteProcessorParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete processor params
func (o *DeleteProcessorParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete processor params
func (o *DeleteProcessorParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *DeleteProcessorParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete processor params
func (o *DeleteProcessorParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the delete processor params
func (o *DeleteProcessorParams) WithID(id string) *DeleteProcessorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete processor params
func (o *DeleteProcessorParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete processor params
func (o *DeleteProcessorParams) WithVersion(version *string) *DeleteProcessorParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete processor params
func (o *DeleteProcessorParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProcessorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string
		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {
			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}

	}

	if o.DisconnectedNodeAcknowledged != nil {

		// query param disconnectedNodeAcknowledged
		var qrDisconnectedNodeAcknowledged bool
		if o.DisconnectedNodeAcknowledged != nil {
			qrDisconnectedNodeAcknowledged = *o.DisconnectedNodeAcknowledged
		}
		qDisconnectedNodeAcknowledged := swag.FormatBool(qrDisconnectedNodeAcknowledged)
		if qDisconnectedNodeAcknowledged != "" {
			if err := r.SetQueryParam("disconnectedNodeAcknowledged", qDisconnectedNodeAcknowledged); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
