// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateRemoteProcessGroupInputPortReader is a Reader for the UpdateRemoteProcessGroupInputPort structure.
type UpdateRemoteProcessGroupInputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemoteProcessGroupInputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRemoteProcessGroupInputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRemoteProcessGroupInputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRemoteProcessGroupInputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateRemoteProcessGroupInputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateRemoteProcessGroupInputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateRemoteProcessGroupInputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRemoteProcessGroupInputPortOK creates a UpdateRemoteProcessGroupInputPortOK with default headers values
func NewUpdateRemoteProcessGroupInputPortOK() *UpdateRemoteProcessGroupInputPortOK {
	return &UpdateRemoteProcessGroupInputPortOK{}
}

/*UpdateRemoteProcessGroupInputPortOK handles this case with default header values.

successful operation
*/
type UpdateRemoteProcessGroupInputPortOK struct {
	Payload *models.RemoteProcessGroupPortEntity
}

func (o *UpdateRemoteProcessGroupInputPortOK) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupInputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupPortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRemoteProcessGroupInputPortBadRequest creates a UpdateRemoteProcessGroupInputPortBadRequest with default headers values
func NewUpdateRemoteProcessGroupInputPortBadRequest() *UpdateRemoteProcessGroupInputPortBadRequest {
	return &UpdateRemoteProcessGroupInputPortBadRequest{}
}

/*UpdateRemoteProcessGroupInputPortBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRemoteProcessGroupInputPortBadRequest struct {
}

func (o *UpdateRemoteProcessGroupInputPortBadRequest) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupInputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortUnauthorized creates a UpdateRemoteProcessGroupInputPortUnauthorized with default headers values
func NewUpdateRemoteProcessGroupInputPortUnauthorized() *UpdateRemoteProcessGroupInputPortUnauthorized {
	return &UpdateRemoteProcessGroupInputPortUnauthorized{}
}

/*UpdateRemoteProcessGroupInputPortUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateRemoteProcessGroupInputPortUnauthorized struct {
}

func (o *UpdateRemoteProcessGroupInputPortUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupInputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortForbidden creates a UpdateRemoteProcessGroupInputPortForbidden with default headers values
func NewUpdateRemoteProcessGroupInputPortForbidden() *UpdateRemoteProcessGroupInputPortForbidden {
	return &UpdateRemoteProcessGroupInputPortForbidden{}
}

/*UpdateRemoteProcessGroupInputPortForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateRemoteProcessGroupInputPortForbidden struct {
}

func (o *UpdateRemoteProcessGroupInputPortForbidden) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupInputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortNotFound creates a UpdateRemoteProcessGroupInputPortNotFound with default headers values
func NewUpdateRemoteProcessGroupInputPortNotFound() *UpdateRemoteProcessGroupInputPortNotFound {
	return &UpdateRemoteProcessGroupInputPortNotFound{}
}

/*UpdateRemoteProcessGroupInputPortNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateRemoteProcessGroupInputPortNotFound struct {
}

func (o *UpdateRemoteProcessGroupInputPortNotFound) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupInputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortConflict creates a UpdateRemoteProcessGroupInputPortConflict with default headers values
func NewUpdateRemoteProcessGroupInputPortConflict() *UpdateRemoteProcessGroupInputPortConflict {
	return &UpdateRemoteProcessGroupInputPortConflict{}
}

/*UpdateRemoteProcessGroupInputPortConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRemoteProcessGroupInputPortConflict struct {
}

func (o *UpdateRemoteProcessGroupInputPortConflict) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortConflict ", 409)
}

func (o *UpdateRemoteProcessGroupInputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}