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

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetUsersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

successful operation
*/
type GetUsersOK struct {
	Payload *models.UsersEntity
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*GetUsersBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetUsersBadRequest struct {
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*GetUsersUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetUsersUnauthorized struct {
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*GetUsersForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetUsersForbidden struct {
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersNotFound creates a GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {
	return &GetUsersNotFound{}
}

/*GetUsersNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetUsersNotFound struct {
}

func (o *GetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersNotFound ", 404)
}

func (o *GetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersConflict creates a GetUsersConflict with default headers values
func NewGetUsersConflict() *GetUsersConflict {
	return &GetUsersConflict{}
}

/*GetUsersConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetUsersConflict struct {
}

func (o *GetUsersConflict) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersConflict ", 409)
}

func (o *GetUsersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
