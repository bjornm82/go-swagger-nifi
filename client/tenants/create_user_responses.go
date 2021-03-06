// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*CreateUserCreated handles this case with default header values.

successful operation
*/
type CreateUserCreated struct {
	Payload *models.UserEntity
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) GetPayload() *models.UserEntity {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*CreateUserBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateUserBadRequest struct {
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserBadRequest ", 400)
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/*CreateUserUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateUserUnauthorized struct {
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserUnauthorized ", 401)
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/*CreateUserForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateUserForbidden struct {
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserForbidden ", 403)
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserNotFound creates a CreateUserNotFound with default headers values
func NewCreateUserNotFound() *CreateUserNotFound {
	return &CreateUserNotFound{}
}

/*CreateUserNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateUserNotFound struct {
}

func (o *CreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserNotFound ", 404)
}

func (o *CreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*CreateUserConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateUserConflict struct {
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /tenants/users][%d] createUserConflict ", 409)
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
