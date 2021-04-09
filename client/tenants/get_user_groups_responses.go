// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetUserGroupsReader is a Reader for the GetUserGroups structure.
type GetUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetUserGroupsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserGroupsOK creates a GetUserGroupsOK with default headers values
func NewGetUserGroupsOK() *GetUserGroupsOK {
	return &GetUserGroupsOK{}
}

/*GetUserGroupsOK handles this case with default header values.

successful operation
*/
type GetUserGroupsOK struct {
	Payload *models.UserGroupsEntity
}

func (o *GetUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsOK  %+v", 200, o.Payload)
}

func (o *GetUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupsBadRequest creates a GetUserGroupsBadRequest with default headers values
func NewGetUserGroupsBadRequest() *GetUserGroupsBadRequest {
	return &GetUserGroupsBadRequest{}
}

/*GetUserGroupsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetUserGroupsBadRequest struct {
}

func (o *GetUserGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsBadRequest ", 400)
}

func (o *GetUserGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupsUnauthorized creates a GetUserGroupsUnauthorized with default headers values
func NewGetUserGroupsUnauthorized() *GetUserGroupsUnauthorized {
	return &GetUserGroupsUnauthorized{}
}

/*GetUserGroupsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetUserGroupsUnauthorized struct {
}

func (o *GetUserGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsUnauthorized ", 401)
}

func (o *GetUserGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupsForbidden creates a GetUserGroupsForbidden with default headers values
func NewGetUserGroupsForbidden() *GetUserGroupsForbidden {
	return &GetUserGroupsForbidden{}
}

/*GetUserGroupsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetUserGroupsForbidden struct {
}

func (o *GetUserGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsForbidden ", 403)
}

func (o *GetUserGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupsNotFound creates a GetUserGroupsNotFound with default headers values
func NewGetUserGroupsNotFound() *GetUserGroupsNotFound {
	return &GetUserGroupsNotFound{}
}

/*GetUserGroupsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetUserGroupsNotFound struct {
}

func (o *GetUserGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsNotFound ", 404)
}

func (o *GetUserGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupsConflict creates a GetUserGroupsConflict with default headers values
func NewGetUserGroupsConflict() *GetUserGroupsConflict {
	return &GetUserGroupsConflict{}
}

/*GetUserGroupsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetUserGroupsConflict struct {
}

func (o *GetUserGroupsConflict) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups][%d] getUserGroupsConflict ", 409)
}

func (o *GetUserGroupsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
