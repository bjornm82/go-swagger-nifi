// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// UpdateVariableRegistryReader is a Reader for the UpdateVariableRegistry structure.
type UpdateVariableRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVariableRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateVariableRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateVariableRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateVariableRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateVariableRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateVariableRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateVariableRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVariableRegistryOK creates a UpdateVariableRegistryOK with default headers values
func NewUpdateVariableRegistryOK() *UpdateVariableRegistryOK {
	return &UpdateVariableRegistryOK{}
}

/*UpdateVariableRegistryOK handles this case with default header values.

successful operation
*/
type UpdateVariableRegistryOK struct {
	Payload *models.VariableRegistryEntity
}

func (o *UpdateVariableRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateVariableRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVariableRegistryBadRequest creates a UpdateVariableRegistryBadRequest with default headers values
func NewUpdateVariableRegistryBadRequest() *UpdateVariableRegistryBadRequest {
	return &UpdateVariableRegistryBadRequest{}
}

/*UpdateVariableRegistryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateVariableRegistryBadRequest struct {
}

func (o *UpdateVariableRegistryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryBadRequest ", 400)
}

func (o *UpdateVariableRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableRegistryUnauthorized creates a UpdateVariableRegistryUnauthorized with default headers values
func NewUpdateVariableRegistryUnauthorized() *UpdateVariableRegistryUnauthorized {
	return &UpdateVariableRegistryUnauthorized{}
}

/*UpdateVariableRegistryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type UpdateVariableRegistryUnauthorized struct {
}

func (o *UpdateVariableRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryUnauthorized ", 401)
}

func (o *UpdateVariableRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableRegistryForbidden creates a UpdateVariableRegistryForbidden with default headers values
func NewUpdateVariableRegistryForbidden() *UpdateVariableRegistryForbidden {
	return &UpdateVariableRegistryForbidden{}
}

/*UpdateVariableRegistryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type UpdateVariableRegistryForbidden struct {
}

func (o *UpdateVariableRegistryForbidden) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryForbidden ", 403)
}

func (o *UpdateVariableRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableRegistryNotFound creates a UpdateVariableRegistryNotFound with default headers values
func NewUpdateVariableRegistryNotFound() *UpdateVariableRegistryNotFound {
	return &UpdateVariableRegistryNotFound{}
}

/*UpdateVariableRegistryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type UpdateVariableRegistryNotFound struct {
}

func (o *UpdateVariableRegistryNotFound) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryNotFound ", 404)
}

func (o *UpdateVariableRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableRegistryConflict creates a UpdateVariableRegistryConflict with default headers values
func NewUpdateVariableRegistryConflict() *UpdateVariableRegistryConflict {
	return &UpdateVariableRegistryConflict{}
}

/*UpdateVariableRegistryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateVariableRegistryConflict struct {
}

func (o *UpdateVariableRegistryConflict) Error() string {
	return fmt.Sprintf("[PUT /process-groups/{id}/variable-registry][%d] updateVariableRegistryConflict ", 409)
}

func (o *UpdateVariableRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
