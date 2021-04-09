// Code generated by go-swagger; DO NOT EDIT.

package connections

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

// NewDeleteConnectionParams creates a new DeleteConnectionParams object
// with the default values initialized.
func NewDeleteConnectionParams() *DeleteConnectionParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteConnectionParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConnectionParamsWithTimeout creates a new DeleteConnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConnectionParamsWithTimeout(timeout time.Duration) *DeleteConnectionParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteConnectionParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		timeout: timeout,
	}
}

// NewDeleteConnectionParamsWithContext creates a new DeleteConnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConnectionParamsWithContext(ctx context.Context) *DeleteConnectionParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteConnectionParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,

		Context: ctx,
	}
}

// NewDeleteConnectionParamsWithHTTPClient creates a new DeleteConnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConnectionParamsWithHTTPClient(client *http.Client) *DeleteConnectionParams {
	var (
		disconnectedNodeAcknowledgedDefault = bool(false)
	)
	return &DeleteConnectionParams{
		DisconnectedNodeAcknowledged: &disconnectedNodeAcknowledgedDefault,
		HTTPClient:                   client,
	}
}

/*DeleteConnectionParams contains all the parameters to send to the API endpoint
for the delete connection operation typically these are written to a http.Request
*/
type DeleteConnectionParams struct {

	/*ClientID
	  If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.

	*/
	ClientID *string
	/*DisconnectedNodeAcknowledged
	  Acknowledges that this node is disconnected to allow for mutable requests to proceed.

	*/
	DisconnectedNodeAcknowledged *bool
	/*ID
	  The connection id.

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

// WithTimeout adds the timeout to the delete connection params
func (o *DeleteConnectionParams) WithTimeout(timeout time.Duration) *DeleteConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete connection params
func (o *DeleteConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete connection params
func (o *DeleteConnectionParams) WithContext(ctx context.Context) *DeleteConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete connection params
func (o *DeleteConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete connection params
func (o *DeleteConnectionParams) WithHTTPClient(client *http.Client) *DeleteConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete connection params
func (o *DeleteConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete connection params
func (o *DeleteConnectionParams) WithClientID(clientID *string) *DeleteConnectionParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete connection params
func (o *DeleteConnectionParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete connection params
func (o *DeleteConnectionParams) WithDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) *DeleteConnectionParams {
	o.SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged)
	return o
}

// SetDisconnectedNodeAcknowledged adds the disconnectedNodeAcknowledged to the delete connection params
func (o *DeleteConnectionParams) SetDisconnectedNodeAcknowledged(disconnectedNodeAcknowledged *bool) {
	o.DisconnectedNodeAcknowledged = disconnectedNodeAcknowledged
}

// WithID adds the id to the delete connection params
func (o *DeleteConnectionParams) WithID(id string) *DeleteConnectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete connection params
func (o *DeleteConnectionParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete connection params
func (o *DeleteConnectionParams) WithVersion(version *string) *DeleteConnectionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete connection params
func (o *DeleteConnectionParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
