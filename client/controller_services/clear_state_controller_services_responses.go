// Code generated by go-swagger; DO NOT EDIT.

package controller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// ClearStateControllerServicesReader is a Reader for the ClearStateControllerServices structure.
type ClearStateControllerServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearStateControllerServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearStateControllerServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewClearStateControllerServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewClearStateControllerServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewClearStateControllerServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClearStateControllerServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewClearStateControllerServicesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClearStateControllerServicesOK creates a ClearStateControllerServicesOK with default headers values
func NewClearStateControllerServicesOK() *ClearStateControllerServicesOK {
	return &ClearStateControllerServicesOK{}
}

/*ClearStateControllerServicesOK handles this case with default header values.

successful operation
*/
type ClearStateControllerServicesOK struct {
	Payload *models.ComponentStateEntity
}

func (o *ClearStateControllerServicesOK) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesOK  %+v", 200, o.Payload)
}

func (o *ClearStateControllerServicesOK) GetPayload() *models.ComponentStateEntity {
	return o.Payload
}

func (o *ClearStateControllerServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearStateControllerServicesBadRequest creates a ClearStateControllerServicesBadRequest with default headers values
func NewClearStateControllerServicesBadRequest() *ClearStateControllerServicesBadRequest {
	return &ClearStateControllerServicesBadRequest{}
}

/*ClearStateControllerServicesBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ClearStateControllerServicesBadRequest struct {
}

func (o *ClearStateControllerServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesBadRequest ", 400)
}

func (o *ClearStateControllerServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateControllerServicesUnauthorized creates a ClearStateControllerServicesUnauthorized with default headers values
func NewClearStateControllerServicesUnauthorized() *ClearStateControllerServicesUnauthorized {
	return &ClearStateControllerServicesUnauthorized{}
}

/*ClearStateControllerServicesUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type ClearStateControllerServicesUnauthorized struct {
}

func (o *ClearStateControllerServicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesUnauthorized ", 401)
}

func (o *ClearStateControllerServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateControllerServicesForbidden creates a ClearStateControllerServicesForbidden with default headers values
func NewClearStateControllerServicesForbidden() *ClearStateControllerServicesForbidden {
	return &ClearStateControllerServicesForbidden{}
}

/*ClearStateControllerServicesForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type ClearStateControllerServicesForbidden struct {
}

func (o *ClearStateControllerServicesForbidden) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesForbidden ", 403)
}

func (o *ClearStateControllerServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateControllerServicesNotFound creates a ClearStateControllerServicesNotFound with default headers values
func NewClearStateControllerServicesNotFound() *ClearStateControllerServicesNotFound {
	return &ClearStateControllerServicesNotFound{}
}

/*ClearStateControllerServicesNotFound handles this case with default header values.

The specified resource could not be found.
*/
type ClearStateControllerServicesNotFound struct {
}

func (o *ClearStateControllerServicesNotFound) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesNotFound ", 404)
}

func (o *ClearStateControllerServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateControllerServicesConflict creates a ClearStateControllerServicesConflict with default headers values
func NewClearStateControllerServicesConflict() *ClearStateControllerServicesConflict {
	return &ClearStateControllerServicesConflict{}
}

/*ClearStateControllerServicesConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ClearStateControllerServicesConflict struct {
}

func (o *ClearStateControllerServicesConflict) Error() string {
	return fmt.Sprintf("[POST /controller-services/{id}/state/clear-requests][%d] clearStateControllerServicesConflict ", 409)
}

func (o *ClearStateControllerServicesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
