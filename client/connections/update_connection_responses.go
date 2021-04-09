// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateConnectionReader is a Reader for the UpdateConnection structure.
type UpdateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConnectionOK creates a UpdateConnectionOK with default headers values
func NewUpdateConnectionOK() *UpdateConnectionOK {
	return &UpdateConnectionOK{}
}

/*UpdateConnectionOK handles this case with default header values.

successful operation
*/
type UpdateConnectionOK struct {
	Payload *models.ConnectionEntity
}

func (o *UpdateConnectionOK) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionBadRequest creates a UpdateConnectionBadRequest with default headers values
func NewUpdateConnectionBadRequest() *UpdateConnectionBadRequest {
	return &UpdateConnectionBadRequest{}
}

/*UpdateConnectionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateConnectionBadRequest struct {
}

func (o *UpdateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionBadRequest ", 400)
}

func (o *UpdateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectionUnauthorized creates a UpdateConnectionUnauthorized with default headers values
func NewUpdateConnectionUnauthorized() *UpdateConnectionUnauthorized {
	return &UpdateConnectionUnauthorized{}
}

/*UpdateConnectionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateConnectionUnauthorized struct {
}

func (o *UpdateConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionUnauthorized ", 401)
}

func (o *UpdateConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectionForbidden creates a UpdateConnectionForbidden with default headers values
func NewUpdateConnectionForbidden() *UpdateConnectionForbidden {
	return &UpdateConnectionForbidden{}
}

/*UpdateConnectionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateConnectionForbidden struct {
}

func (o *UpdateConnectionForbidden) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionForbidden ", 403)
}

func (o *UpdateConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectionNotFound creates a UpdateConnectionNotFound with default headers values
func NewUpdateConnectionNotFound() *UpdateConnectionNotFound {
	return &UpdateConnectionNotFound{}
}

/*UpdateConnectionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateConnectionNotFound struct {
}

func (o *UpdateConnectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionNotFound ", 404)
}

func (o *UpdateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectionConflict creates a UpdateConnectionConflict with default headers values
func NewUpdateConnectionConflict() *UpdateConnectionConflict {
	return &UpdateConnectionConflict{}
}

/*UpdateConnectionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateConnectionConflict struct {
}

func (o *UpdateConnectionConflict) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionConflict ", 409)
}

func (o *UpdateConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
