// Code generated by go-swagger; DO NOT EDIT.

package funnel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// RemoveFunnelReader is a Reader for the RemoveFunnel structure.
type RemoveFunnelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveFunnelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveFunnelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveFunnelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveFunnelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveFunnelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveFunnelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveFunnelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveFunnelOK creates a RemoveFunnelOK with default headers values
func NewRemoveFunnelOK() *RemoveFunnelOK {
	return &RemoveFunnelOK{}
}

/*RemoveFunnelOK handles this case with default header values.

successful operation
*/
type RemoveFunnelOK struct {
	Payload *models.FunnelEntity
}

func (o *RemoveFunnelOK) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelOK  %+v", 200, o.Payload)
}

func (o *RemoveFunnelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveFunnelBadRequest creates a RemoveFunnelBadRequest with default headers values
func NewRemoveFunnelBadRequest() *RemoveFunnelBadRequest {
	return &RemoveFunnelBadRequest{}
}

/*RemoveFunnelBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveFunnelBadRequest struct {
}

func (o *RemoveFunnelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelBadRequest ", 400)
}

func (o *RemoveFunnelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveFunnelUnauthorized creates a RemoveFunnelUnauthorized with default headers values
func NewRemoveFunnelUnauthorized() *RemoveFunnelUnauthorized {
	return &RemoveFunnelUnauthorized{}
}

/*RemoveFunnelUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveFunnelUnauthorized struct {
}

func (o *RemoveFunnelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelUnauthorized ", 401)
}

func (o *RemoveFunnelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveFunnelForbidden creates a RemoveFunnelForbidden with default headers values
func NewRemoveFunnelForbidden() *RemoveFunnelForbidden {
	return &RemoveFunnelForbidden{}
}

/*RemoveFunnelForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveFunnelForbidden struct {
}

func (o *RemoveFunnelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelForbidden ", 403)
}

func (o *RemoveFunnelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveFunnelNotFound creates a RemoveFunnelNotFound with default headers values
func NewRemoveFunnelNotFound() *RemoveFunnelNotFound {
	return &RemoveFunnelNotFound{}
}

/*RemoveFunnelNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveFunnelNotFound struct {
}

func (o *RemoveFunnelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelNotFound ", 404)
}

func (o *RemoveFunnelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveFunnelConflict creates a RemoveFunnelConflict with default headers values
func NewRemoveFunnelConflict() *RemoveFunnelConflict {
	return &RemoveFunnelConflict{}
}

/*RemoveFunnelConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveFunnelConflict struct {
}

func (o *RemoveFunnelConflict) Error() string {
	return fmt.Sprintf("[DELETE /funnels/{id}][%d] removeFunnelConflict ", 409)
}

func (o *RemoveFunnelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
