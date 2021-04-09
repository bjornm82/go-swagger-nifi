// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateRegistryClientReader is a Reader for the UpdateRegistryClient structure.
type UpdateRegistryClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegistryClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRegistryClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRegistryClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRegistryClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateRegistryClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateRegistryClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateRegistryClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRegistryClientOK creates a UpdateRegistryClientOK with default headers values
func NewUpdateRegistryClientOK() *UpdateRegistryClientOK {
	return &UpdateRegistryClientOK{}
}

/*UpdateRegistryClientOK handles this case with default header values.

successful operation
*/
type UpdateRegistryClientOK struct {
	Payload *models.RegistryClientEntity
}

func (o *UpdateRegistryClientOK) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientOK  %+v", 200, o.Payload)
}

func (o *UpdateRegistryClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistryClientEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegistryClientBadRequest creates a UpdateRegistryClientBadRequest with default headers values
func NewUpdateRegistryClientBadRequest() *UpdateRegistryClientBadRequest {
	return &UpdateRegistryClientBadRequest{}
}

/*UpdateRegistryClientBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRegistryClientBadRequest struct {
}

func (o *UpdateRegistryClientBadRequest) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientBadRequest ", 400)
}

func (o *UpdateRegistryClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRegistryClientUnauthorized creates a UpdateRegistryClientUnauthorized with default headers values
func NewUpdateRegistryClientUnauthorized() *UpdateRegistryClientUnauthorized {
	return &UpdateRegistryClientUnauthorized{}
}

/*UpdateRegistryClientUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateRegistryClientUnauthorized struct {
}

func (o *UpdateRegistryClientUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientUnauthorized ", 401)
}

func (o *UpdateRegistryClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRegistryClientForbidden creates a UpdateRegistryClientForbidden with default headers values
func NewUpdateRegistryClientForbidden() *UpdateRegistryClientForbidden {
	return &UpdateRegistryClientForbidden{}
}

/*UpdateRegistryClientForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateRegistryClientForbidden struct {
}

func (o *UpdateRegistryClientForbidden) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientForbidden ", 403)
}

func (o *UpdateRegistryClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRegistryClientNotFound creates a UpdateRegistryClientNotFound with default headers values
func NewUpdateRegistryClientNotFound() *UpdateRegistryClientNotFound {
	return &UpdateRegistryClientNotFound{}
}

/*UpdateRegistryClientNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateRegistryClientNotFound struct {
}

func (o *UpdateRegistryClientNotFound) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientNotFound ", 404)
}

func (o *UpdateRegistryClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRegistryClientConflict creates a UpdateRegistryClientConflict with default headers values
func NewUpdateRegistryClientConflict() *UpdateRegistryClientConflict {
	return &UpdateRegistryClientConflict{}
}

/*UpdateRegistryClientConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRegistryClientConflict struct {
}

func (o *UpdateRegistryClientConflict) Error() string {
	return fmt.Sprintf("[PUT /controller/registry-clients/{id}][%d] updateRegistryClientConflict ", 409)
}

func (o *UpdateRegistryClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
