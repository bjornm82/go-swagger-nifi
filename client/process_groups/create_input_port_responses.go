// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CreateInputPortReader is a Reader for the CreateInputPort structure.
type CreateInputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInputPortCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateInputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateInputPortCreated creates a CreateInputPortCreated with default headers values
func NewCreateInputPortCreated() *CreateInputPortCreated {
	return &CreateInputPortCreated{}
}

/*CreateInputPortCreated handles this case with default header values.

successful operation
*/
type CreateInputPortCreated struct {
	Payload *models.PortEntity
}

func (o *CreateInputPortCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortCreated  %+v", 201, o.Payload)
}

func (o *CreateInputPortCreated) GetPayload() *models.PortEntity {
	return o.Payload
}

func (o *CreateInputPortCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInputPortBadRequest creates a CreateInputPortBadRequest with default headers values
func NewCreateInputPortBadRequest() *CreateInputPortBadRequest {
	return &CreateInputPortBadRequest{}
}

/*CreateInputPortBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateInputPortBadRequest struct {
}

func (o *CreateInputPortBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortBadRequest ", 400)
}

func (o *CreateInputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInputPortUnauthorized creates a CreateInputPortUnauthorized with default headers values
func NewCreateInputPortUnauthorized() *CreateInputPortUnauthorized {
	return &CreateInputPortUnauthorized{}
}

/*CreateInputPortUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateInputPortUnauthorized struct {
}

func (o *CreateInputPortUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortUnauthorized ", 401)
}

func (o *CreateInputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInputPortForbidden creates a CreateInputPortForbidden with default headers values
func NewCreateInputPortForbidden() *CreateInputPortForbidden {
	return &CreateInputPortForbidden{}
}

/*CreateInputPortForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateInputPortForbidden struct {
}

func (o *CreateInputPortForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortForbidden ", 403)
}

func (o *CreateInputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInputPortNotFound creates a CreateInputPortNotFound with default headers values
func NewCreateInputPortNotFound() *CreateInputPortNotFound {
	return &CreateInputPortNotFound{}
}

/*CreateInputPortNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateInputPortNotFound struct {
}

func (o *CreateInputPortNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortNotFound ", 404)
}

func (o *CreateInputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInputPortConflict creates a CreateInputPortConflict with default headers values
func NewCreateInputPortConflict() *CreateInputPortConflict {
	return &CreateInputPortConflict{}
}

/*CreateInputPortConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateInputPortConflict struct {
}

func (o *CreateInputPortConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/input-ports][%d] createInputPortConflict ", 409)
}

func (o *CreateInputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
