// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*UpdateUserOK handles this case with default header values.

successful operation
*/
type UpdateUserOK struct {
	Payload *models.UserEntity
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/*UpdateUserBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateUserBadRequest struct {
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserBadRequest ", 400)
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserUnauthorized creates a UpdateUserUnauthorized with default headers values
func NewUpdateUserUnauthorized() *UpdateUserUnauthorized {
	return &UpdateUserUnauthorized{}
}

/*UpdateUserUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateUserUnauthorized struct {
}

func (o *UpdateUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserUnauthorized ", 401)
}

func (o *UpdateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/*UpdateUserForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateUserForbidden struct {
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserForbidden ", 403)
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/*UpdateUserNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateUserNotFound struct {
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserNotFound ", 404)
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserConflict creates a UpdateUserConflict with default headers values
func NewUpdateUserConflict() *UpdateUserConflict {
	return &UpdateUserConflict{}
}

/*UpdateUserConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateUserConflict struct {
}

func (o *UpdateUserConflict) Error() string {
	return fmt.Sprintf("[PUT /tenants/users/{id}][%d] updateUserConflict ", 409)
}

func (o *UpdateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
