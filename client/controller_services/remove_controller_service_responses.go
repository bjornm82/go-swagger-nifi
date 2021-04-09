// Code generated by go-swagger; DO NOT EDIT.

package controller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// RemoveControllerServiceReader is a Reader for the RemoveControllerService structure.
type RemoveControllerServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveControllerServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveControllerServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveControllerServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveControllerServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveControllerServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveControllerServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveControllerServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveControllerServiceOK creates a RemoveControllerServiceOK with default headers values
func NewRemoveControllerServiceOK() *RemoveControllerServiceOK {
	return &RemoveControllerServiceOK{}
}

/*RemoveControllerServiceOK handles this case with default header values.

successful operation
*/
type RemoveControllerServiceOK struct {
	Payload *models.ControllerServiceEntity
}

func (o *RemoveControllerServiceOK) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceOK  %+v", 200, o.Payload)
}

func (o *RemoveControllerServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveControllerServiceBadRequest creates a RemoveControllerServiceBadRequest with default headers values
func NewRemoveControllerServiceBadRequest() *RemoveControllerServiceBadRequest {
	return &RemoveControllerServiceBadRequest{}
}

/*RemoveControllerServiceBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveControllerServiceBadRequest struct {
}

func (o *RemoveControllerServiceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceBadRequest ", 400)
}

func (o *RemoveControllerServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceUnauthorized creates a RemoveControllerServiceUnauthorized with default headers values
func NewRemoveControllerServiceUnauthorized() *RemoveControllerServiceUnauthorized {
	return &RemoveControllerServiceUnauthorized{}
}

/*RemoveControllerServiceUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveControllerServiceUnauthorized struct {
}

func (o *RemoveControllerServiceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceUnauthorized ", 401)
}

func (o *RemoveControllerServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceForbidden creates a RemoveControllerServiceForbidden with default headers values
func NewRemoveControllerServiceForbidden() *RemoveControllerServiceForbidden {
	return &RemoveControllerServiceForbidden{}
}

/*RemoveControllerServiceForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveControllerServiceForbidden struct {
}

func (o *RemoveControllerServiceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceForbidden ", 403)
}

func (o *RemoveControllerServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceNotFound creates a RemoveControllerServiceNotFound with default headers values
func NewRemoveControllerServiceNotFound() *RemoveControllerServiceNotFound {
	return &RemoveControllerServiceNotFound{}
}

/*RemoveControllerServiceNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveControllerServiceNotFound struct {
}

func (o *RemoveControllerServiceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceNotFound ", 404)
}

func (o *RemoveControllerServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceConflict creates a RemoveControllerServiceConflict with default headers values
func NewRemoveControllerServiceConflict() *RemoveControllerServiceConflict {
	return &RemoveControllerServiceConflict{}
}

/*RemoveControllerServiceConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveControllerServiceConflict struct {
}

func (o *RemoveControllerServiceConflict) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceConflict ", 409)
}

func (o *RemoveControllerServiceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
