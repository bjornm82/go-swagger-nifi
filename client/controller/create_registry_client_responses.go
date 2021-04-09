// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// CreateRegistryClientReader is a Reader for the CreateRegistryClient structure.
type CreateRegistryClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRegistryClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateRegistryClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateRegistryClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateRegistryClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateRegistryClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRegistryClientOK creates a CreateRegistryClientOK with default headers values
func NewCreateRegistryClientOK() *CreateRegistryClientOK {
	return &CreateRegistryClientOK{}
}

/*CreateRegistryClientOK handles this case with default header values.

successful operation
*/
type CreateRegistryClientOK struct {
	Payload *models.RegistryClientEntity
}

func (o *CreateRegistryClientOK) Error() string {
	return fmt.Sprintf("[POST /controller/registry-clients][%d] createRegistryClientOK  %+v", 200, o.Payload)
}

func (o *CreateRegistryClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistryClientEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegistryClientBadRequest creates a CreateRegistryClientBadRequest with default headers values
func NewCreateRegistryClientBadRequest() *CreateRegistryClientBadRequest {
	return &CreateRegistryClientBadRequest{}
}

/*CreateRegistryClientBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateRegistryClientBadRequest struct {
}

func (o *CreateRegistryClientBadRequest) Error() string {
	return fmt.Sprintf("[POST /controller/registry-clients][%d] createRegistryClientBadRequest ", 400)
}

func (o *CreateRegistryClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRegistryClientUnauthorized creates a CreateRegistryClientUnauthorized with default headers values
func NewCreateRegistryClientUnauthorized() *CreateRegistryClientUnauthorized {
	return &CreateRegistryClientUnauthorized{}
}

/*CreateRegistryClientUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateRegistryClientUnauthorized struct {
}

func (o *CreateRegistryClientUnauthorized) Error() string {
	return fmt.Sprintf("[POST /controller/registry-clients][%d] createRegistryClientUnauthorized ", 401)
}

func (o *CreateRegistryClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRegistryClientForbidden creates a CreateRegistryClientForbidden with default headers values
func NewCreateRegistryClientForbidden() *CreateRegistryClientForbidden {
	return &CreateRegistryClientForbidden{}
}

/*CreateRegistryClientForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateRegistryClientForbidden struct {
}

func (o *CreateRegistryClientForbidden) Error() string {
	return fmt.Sprintf("[POST /controller/registry-clients][%d] createRegistryClientForbidden ", 403)
}

func (o *CreateRegistryClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRegistryClientConflict creates a CreateRegistryClientConflict with default headers values
func NewCreateRegistryClientConflict() *CreateRegistryClientConflict {
	return &CreateRegistryClientConflict{}
}

/*CreateRegistryClientConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateRegistryClientConflict struct {
}

func (o *CreateRegistryClientConflict) Error() string {
	return fmt.Sprintf("[POST /controller/registry-clients][%d] createRegistryClientConflict ", 409)
}

func (o *CreateRegistryClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
