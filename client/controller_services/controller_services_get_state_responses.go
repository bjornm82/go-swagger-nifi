// Code generated by go-swagger; DO NOT EDIT.

package controller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// ControllerServicesGetStateReader is a Reader for the ControllerServicesGetState structure.
type ControllerServicesGetStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllerServicesGetStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllerServicesGetStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewControllerServicesGetStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewControllerServicesGetStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllerServicesGetStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllerServicesGetStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewControllerServicesGetStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllerServicesGetStateOK creates a ControllerServicesGetStateOK with default headers values
func NewControllerServicesGetStateOK() *ControllerServicesGetStateOK {
	return &ControllerServicesGetStateOK{}
}

/*ControllerServicesGetStateOK handles this case with default header values.

successful operation
*/
type ControllerServicesGetStateOK struct {
	Payload *models.ComponentStateEntity
}

func (o *ControllerServicesGetStateOK) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateOK  %+v", 200, o.Payload)
}

func (o *ControllerServicesGetStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllerServicesGetStateBadRequest creates a ControllerServicesGetStateBadRequest with default headers values
func NewControllerServicesGetStateBadRequest() *ControllerServicesGetStateBadRequest {
	return &ControllerServicesGetStateBadRequest{}
}

/*ControllerServicesGetStateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ControllerServicesGetStateBadRequest struct {
}

func (o *ControllerServicesGetStateBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateBadRequest ", 400)
}

func (o *ControllerServicesGetStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllerServicesGetStateUnauthorized creates a ControllerServicesGetStateUnauthorized with default headers values
func NewControllerServicesGetStateUnauthorized() *ControllerServicesGetStateUnauthorized {
	return &ControllerServicesGetStateUnauthorized{}
}

/*ControllerServicesGetStateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type ControllerServicesGetStateUnauthorized struct {
}

func (o *ControllerServicesGetStateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateUnauthorized ", 401)
}

func (o *ControllerServicesGetStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllerServicesGetStateForbidden creates a ControllerServicesGetStateForbidden with default headers values
func NewControllerServicesGetStateForbidden() *ControllerServicesGetStateForbidden {
	return &ControllerServicesGetStateForbidden{}
}

/*ControllerServicesGetStateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type ControllerServicesGetStateForbidden struct {
}

func (o *ControllerServicesGetStateForbidden) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateForbidden ", 403)
}

func (o *ControllerServicesGetStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllerServicesGetStateNotFound creates a ControllerServicesGetStateNotFound with default headers values
func NewControllerServicesGetStateNotFound() *ControllerServicesGetStateNotFound {
	return &ControllerServicesGetStateNotFound{}
}

/*ControllerServicesGetStateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type ControllerServicesGetStateNotFound struct {
}

func (o *ControllerServicesGetStateNotFound) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateNotFound ", 404)
}

func (o *ControllerServicesGetStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllerServicesGetStateConflict creates a ControllerServicesGetStateConflict with default headers values
func NewControllerServicesGetStateConflict() *ControllerServicesGetStateConflict {
	return &ControllerServicesGetStateConflict{}
}

/*ControllerServicesGetStateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ControllerServicesGetStateConflict struct {
}

func (o *ControllerServicesGetStateConflict) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/state][%d] controllerServicesGetStateConflict ", 409)
}

func (o *ControllerServicesGetStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
