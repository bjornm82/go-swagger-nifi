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

// UpdateControllerServiceReferencesReader is a Reader for the UpdateControllerServiceReferences structure.
type UpdateControllerServiceReferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateControllerServiceReferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateControllerServiceReferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateControllerServiceReferencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateControllerServiceReferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateControllerServiceReferencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateControllerServiceReferencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateControllerServiceReferencesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateControllerServiceReferencesOK creates a UpdateControllerServiceReferencesOK with default headers values
func NewUpdateControllerServiceReferencesOK() *UpdateControllerServiceReferencesOK {
	return &UpdateControllerServiceReferencesOK{}
}

/*UpdateControllerServiceReferencesOK handles this case with default header values.

successful operation
*/
type UpdateControllerServiceReferencesOK struct {
	Payload *models.ControllerServiceReferencingComponentsEntity
}

func (o *UpdateControllerServiceReferencesOK) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesOK  %+v", 200, o.Payload)
}

func (o *UpdateControllerServiceReferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceReferencingComponentsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateControllerServiceReferencesBadRequest creates a UpdateControllerServiceReferencesBadRequest with default headers values
func NewUpdateControllerServiceReferencesBadRequest() *UpdateControllerServiceReferencesBadRequest {
	return &UpdateControllerServiceReferencesBadRequest{}
}

/*UpdateControllerServiceReferencesBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateControllerServiceReferencesBadRequest struct {
}

func (o *UpdateControllerServiceReferencesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesBadRequest ", 400)
}

func (o *UpdateControllerServiceReferencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerServiceReferencesUnauthorized creates a UpdateControllerServiceReferencesUnauthorized with default headers values
func NewUpdateControllerServiceReferencesUnauthorized() *UpdateControllerServiceReferencesUnauthorized {
	return &UpdateControllerServiceReferencesUnauthorized{}
}

/*UpdateControllerServiceReferencesUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateControllerServiceReferencesUnauthorized struct {
}

func (o *UpdateControllerServiceReferencesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesUnauthorized ", 401)
}

func (o *UpdateControllerServiceReferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerServiceReferencesForbidden creates a UpdateControllerServiceReferencesForbidden with default headers values
func NewUpdateControllerServiceReferencesForbidden() *UpdateControllerServiceReferencesForbidden {
	return &UpdateControllerServiceReferencesForbidden{}
}

/*UpdateControllerServiceReferencesForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateControllerServiceReferencesForbidden struct {
}

func (o *UpdateControllerServiceReferencesForbidden) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesForbidden ", 403)
}

func (o *UpdateControllerServiceReferencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerServiceReferencesNotFound creates a UpdateControllerServiceReferencesNotFound with default headers values
func NewUpdateControllerServiceReferencesNotFound() *UpdateControllerServiceReferencesNotFound {
	return &UpdateControllerServiceReferencesNotFound{}
}

/*UpdateControllerServiceReferencesNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateControllerServiceReferencesNotFound struct {
}

func (o *UpdateControllerServiceReferencesNotFound) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesNotFound ", 404)
}

func (o *UpdateControllerServiceReferencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerServiceReferencesConflict creates a UpdateControllerServiceReferencesConflict with default headers values
func NewUpdateControllerServiceReferencesConflict() *UpdateControllerServiceReferencesConflict {
	return &UpdateControllerServiceReferencesConflict{}
}

/*UpdateControllerServiceReferencesConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateControllerServiceReferencesConflict struct {
}

func (o *UpdateControllerServiceReferencesConflict) Error() string {
	return fmt.Sprintf("[PUT /controller-services/{id}/references][%d] updateControllerServiceReferencesConflict ", 409)
}

func (o *UpdateControllerServiceReferencesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}