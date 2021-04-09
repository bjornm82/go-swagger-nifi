// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateAccessPolicyReader is a Reader for the UpdateAccessPolicy structure.
type UpdateAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateAccessPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateAccessPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateAccessPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateAccessPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateAccessPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAccessPolicyOK creates a UpdateAccessPolicyOK with default headers values
func NewUpdateAccessPolicyOK() *UpdateAccessPolicyOK {
	return &UpdateAccessPolicyOK{}
}

/*UpdateAccessPolicyOK handles this case with default header values.

successful operation
*/
type UpdateAccessPolicyOK struct {
	Payload *models.AccessPolicyEntity
}

func (o *UpdateAccessPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccessPolicyBadRequest creates a UpdateAccessPolicyBadRequest with default headers values
func NewUpdateAccessPolicyBadRequest() *UpdateAccessPolicyBadRequest {
	return &UpdateAccessPolicyBadRequest{}
}

/*UpdateAccessPolicyBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateAccessPolicyBadRequest struct {
}

func (o *UpdateAccessPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyBadRequest ", 400)
}

func (o *UpdateAccessPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyUnauthorized creates a UpdateAccessPolicyUnauthorized with default headers values
func NewUpdateAccessPolicyUnauthorized() *UpdateAccessPolicyUnauthorized {
	return &UpdateAccessPolicyUnauthorized{}
}

/*UpdateAccessPolicyUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateAccessPolicyUnauthorized struct {
}

func (o *UpdateAccessPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyUnauthorized ", 401)
}

func (o *UpdateAccessPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyForbidden creates a UpdateAccessPolicyForbidden with default headers values
func NewUpdateAccessPolicyForbidden() *UpdateAccessPolicyForbidden {
	return &UpdateAccessPolicyForbidden{}
}

/*UpdateAccessPolicyForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateAccessPolicyForbidden struct {
}

func (o *UpdateAccessPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyForbidden ", 403)
}

func (o *UpdateAccessPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyNotFound creates a UpdateAccessPolicyNotFound with default headers values
func NewUpdateAccessPolicyNotFound() *UpdateAccessPolicyNotFound {
	return &UpdateAccessPolicyNotFound{}
}

/*UpdateAccessPolicyNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateAccessPolicyNotFound struct {
}

func (o *UpdateAccessPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyNotFound ", 404)
}

func (o *UpdateAccessPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyConflict creates a UpdateAccessPolicyConflict with default headers values
func NewUpdateAccessPolicyConflict() *UpdateAccessPolicyConflict {
	return &UpdateAccessPolicyConflict{}
}

/*UpdateAccessPolicyConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateAccessPolicyConflict struct {
}

func (o *UpdateAccessPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyConflict ", 409)
}

func (o *UpdateAccessPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
