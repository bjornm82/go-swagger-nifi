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

// RemoveAccessPolicyReader is a Reader for the RemoveAccessPolicy structure.
type RemoveAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveAccessPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveAccessPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveAccessPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveAccessPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveAccessPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAccessPolicyOK creates a RemoveAccessPolicyOK with default headers values
func NewRemoveAccessPolicyOK() *RemoveAccessPolicyOK {
	return &RemoveAccessPolicyOK{}
}

/*RemoveAccessPolicyOK handles this case with default header values.

successful operation
*/
type RemoveAccessPolicyOK struct {
	Payload *models.AccessPolicyEntity
}

func (o *RemoveAccessPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *RemoveAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAccessPolicyBadRequest creates a RemoveAccessPolicyBadRequest with default headers values
func NewRemoveAccessPolicyBadRequest() *RemoveAccessPolicyBadRequest {
	return &RemoveAccessPolicyBadRequest{}
}

/*RemoveAccessPolicyBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveAccessPolicyBadRequest struct {
}

func (o *RemoveAccessPolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyBadRequest ", 400)
}

func (o *RemoveAccessPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccessPolicyUnauthorized creates a RemoveAccessPolicyUnauthorized with default headers values
func NewRemoveAccessPolicyUnauthorized() *RemoveAccessPolicyUnauthorized {
	return &RemoveAccessPolicyUnauthorized{}
}

/*RemoveAccessPolicyUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveAccessPolicyUnauthorized struct {
}

func (o *RemoveAccessPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyUnauthorized ", 401)
}

func (o *RemoveAccessPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccessPolicyForbidden creates a RemoveAccessPolicyForbidden with default headers values
func NewRemoveAccessPolicyForbidden() *RemoveAccessPolicyForbidden {
	return &RemoveAccessPolicyForbidden{}
}

/*RemoveAccessPolicyForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveAccessPolicyForbidden struct {
}

func (o *RemoveAccessPolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyForbidden ", 403)
}

func (o *RemoveAccessPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccessPolicyNotFound creates a RemoveAccessPolicyNotFound with default headers values
func NewRemoveAccessPolicyNotFound() *RemoveAccessPolicyNotFound {
	return &RemoveAccessPolicyNotFound{}
}

/*RemoveAccessPolicyNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveAccessPolicyNotFound struct {
}

func (o *RemoveAccessPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyNotFound ", 404)
}

func (o *RemoveAccessPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccessPolicyConflict creates a RemoveAccessPolicyConflict with default headers values
func NewRemoveAccessPolicyConflict() *RemoveAccessPolicyConflict {
	return &RemoveAccessPolicyConflict{}
}

/*RemoveAccessPolicyConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveAccessPolicyConflict struct {
}

func (o *RemoveAccessPolicyConflict) Error() string {
	return fmt.Sprintf("[DELETE /policies/{id}][%d] removeAccessPolicyConflict ", 409)
}

func (o *RemoveAccessPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
