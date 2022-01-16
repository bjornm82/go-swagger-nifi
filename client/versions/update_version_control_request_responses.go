// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateVersionControlRequestReader is a Reader for the UpdateVersionControlRequest structure.
type UpdateVersionControlRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVersionControlRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVersionControlRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVersionControlRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateVersionControlRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVersionControlRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVersionControlRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateVersionControlRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVersionControlRequestOK creates a UpdateVersionControlRequestOK with default headers values
func NewUpdateVersionControlRequestOK() *UpdateVersionControlRequestOK {
	return &UpdateVersionControlRequestOK{}
}

/*UpdateVersionControlRequestOK handles this case with default header values.

successful operation
*/
type UpdateVersionControlRequestOK struct {
	Payload *models.VersionControlInformationEntity
}

func (o *UpdateVersionControlRequestOK) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestOK  %+v", 200, o.Payload)
}

func (o *UpdateVersionControlRequestOK) GetPayload() *models.VersionControlInformationEntity {
	return o.Payload
}

func (o *UpdateVersionControlRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionControlInformationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVersionControlRequestBadRequest creates a UpdateVersionControlRequestBadRequest with default headers values
func NewUpdateVersionControlRequestBadRequest() *UpdateVersionControlRequestBadRequest {
	return &UpdateVersionControlRequestBadRequest{}
}

/*UpdateVersionControlRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateVersionControlRequestBadRequest struct {
}

func (o *UpdateVersionControlRequestBadRequest) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestBadRequest ", 400)
}

func (o *UpdateVersionControlRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVersionControlRequestUnauthorized creates a UpdateVersionControlRequestUnauthorized with default headers values
func NewUpdateVersionControlRequestUnauthorized() *UpdateVersionControlRequestUnauthorized {
	return &UpdateVersionControlRequestUnauthorized{}
}

/*UpdateVersionControlRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateVersionControlRequestUnauthorized struct {
}

func (o *UpdateVersionControlRequestUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestUnauthorized ", 401)
}

func (o *UpdateVersionControlRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVersionControlRequestForbidden creates a UpdateVersionControlRequestForbidden with default headers values
func NewUpdateVersionControlRequestForbidden() *UpdateVersionControlRequestForbidden {
	return &UpdateVersionControlRequestForbidden{}
}

/*UpdateVersionControlRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateVersionControlRequestForbidden struct {
}

func (o *UpdateVersionControlRequestForbidden) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestForbidden ", 403)
}

func (o *UpdateVersionControlRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVersionControlRequestNotFound creates a UpdateVersionControlRequestNotFound with default headers values
func NewUpdateVersionControlRequestNotFound() *UpdateVersionControlRequestNotFound {
	return &UpdateVersionControlRequestNotFound{}
}

/*UpdateVersionControlRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateVersionControlRequestNotFound struct {
}

func (o *UpdateVersionControlRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestNotFound ", 404)
}

func (o *UpdateVersionControlRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVersionControlRequestConflict creates a UpdateVersionControlRequestConflict with default headers values
func NewUpdateVersionControlRequestConflict() *UpdateVersionControlRequestConflict {
	return &UpdateVersionControlRequestConflict{}
}

/*UpdateVersionControlRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateVersionControlRequestConflict struct {
}

func (o *UpdateVersionControlRequestConflict) Error() string {
	return fmt.Sprintf("[PUT /versions/active-requests/{id}][%d] updateVersionControlRequestConflict ", 409)
}

func (o *UpdateVersionControlRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
