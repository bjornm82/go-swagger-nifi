// Code generated by go-swagger; DO NOT EDIT.

package input_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// RemoveInputPortReader is a Reader for the RemoveInputPort structure.
type RemoveInputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveInputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveInputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveInputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveInputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveInputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveInputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveInputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveInputPortOK creates a RemoveInputPortOK with default headers values
func NewRemoveInputPortOK() *RemoveInputPortOK {
	return &RemoveInputPortOK{}
}

/*RemoveInputPortOK handles this case with default header values.

successful operation
*/
type RemoveInputPortOK struct {
	Payload *models.PortEntity
}

func (o *RemoveInputPortOK) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortOK  %+v", 200, o.Payload)
}

func (o *RemoveInputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveInputPortBadRequest creates a RemoveInputPortBadRequest with default headers values
func NewRemoveInputPortBadRequest() *RemoveInputPortBadRequest {
	return &RemoveInputPortBadRequest{}
}

/*RemoveInputPortBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveInputPortBadRequest struct {
}

func (o *RemoveInputPortBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortBadRequest ", 400)
}

func (o *RemoveInputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveInputPortUnauthorized creates a RemoveInputPortUnauthorized with default headers values
func NewRemoveInputPortUnauthorized() *RemoveInputPortUnauthorized {
	return &RemoveInputPortUnauthorized{}
}

/*RemoveInputPortUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveInputPortUnauthorized struct {
}

func (o *RemoveInputPortUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortUnauthorized ", 401)
}

func (o *RemoveInputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveInputPortForbidden creates a RemoveInputPortForbidden with default headers values
func NewRemoveInputPortForbidden() *RemoveInputPortForbidden {
	return &RemoveInputPortForbidden{}
}

/*RemoveInputPortForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveInputPortForbidden struct {
}

func (o *RemoveInputPortForbidden) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortForbidden ", 403)
}

func (o *RemoveInputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveInputPortNotFound creates a RemoveInputPortNotFound with default headers values
func NewRemoveInputPortNotFound() *RemoveInputPortNotFound {
	return &RemoveInputPortNotFound{}
}

/*RemoveInputPortNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveInputPortNotFound struct {
}

func (o *RemoveInputPortNotFound) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortNotFound ", 404)
}

func (o *RemoveInputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveInputPortConflict creates a RemoveInputPortConflict with default headers values
func NewRemoveInputPortConflict() *RemoveInputPortConflict {
	return &RemoveInputPortConflict{}
}

/*RemoveInputPortConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveInputPortConflict struct {
}

func (o *RemoveInputPortConflict) Error() string {
	return fmt.Sprintf("[DELETE /input-ports/{id}][%d] removeInputPortConflict ", 409)
}

func (o *RemoveInputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
