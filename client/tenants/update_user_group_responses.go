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

// UpdateUserGroupReader is a Reader for the UpdateUserGroup structure.
type UpdateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserGroupOK creates a UpdateUserGroupOK with default headers values
func NewUpdateUserGroupOK() *UpdateUserGroupOK {
	return &UpdateUserGroupOK{}
}

/*UpdateUserGroupOK handles this case with default header values.

successful operation
*/
type UpdateUserGroupOK struct {
	Payload *models.UserGroupEntity
}

func (o *UpdateUserGroupOK) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateUserGroupOK) GetPayload() *models.UserGroupEntity {
	return o.Payload
}

func (o *UpdateUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupBadRequest creates a UpdateUserGroupBadRequest with default headers values
func NewUpdateUserGroupBadRequest() *UpdateUserGroupBadRequest {
	return &UpdateUserGroupBadRequest{}
}

/*UpdateUserGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateUserGroupBadRequest struct {
}

func (o *UpdateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupBadRequest ", 400)
}

func (o *UpdateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupUnauthorized creates a UpdateUserGroupUnauthorized with default headers values
func NewUpdateUserGroupUnauthorized() *UpdateUserGroupUnauthorized {
	return &UpdateUserGroupUnauthorized{}
}

/*UpdateUserGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateUserGroupUnauthorized struct {
}

func (o *UpdateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupUnauthorized ", 401)
}

func (o *UpdateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupForbidden creates a UpdateUserGroupForbidden with default headers values
func NewUpdateUserGroupForbidden() *UpdateUserGroupForbidden {
	return &UpdateUserGroupForbidden{}
}

/*UpdateUserGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateUserGroupForbidden struct {
}

func (o *UpdateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupForbidden ", 403)
}

func (o *UpdateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupNotFound creates a UpdateUserGroupNotFound with default headers values
func NewUpdateUserGroupNotFound() *UpdateUserGroupNotFound {
	return &UpdateUserGroupNotFound{}
}

/*UpdateUserGroupNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateUserGroupNotFound struct {
}

func (o *UpdateUserGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupNotFound ", 404)
}

func (o *UpdateUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupConflict creates a UpdateUserGroupConflict with default headers values
func NewUpdateUserGroupConflict() *UpdateUserGroupConflict {
	return &UpdateUserGroupConflict{}
}

/*UpdateUserGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateUserGroupConflict struct {
}

func (o *UpdateUserGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupConflict ", 409)
}

func (o *UpdateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
