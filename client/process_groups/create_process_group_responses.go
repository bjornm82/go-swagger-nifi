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

// CreateProcessGroupReader is a Reader for the CreateProcessGroup structure.
type CreateProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateProcessGroupOK creates a CreateProcessGroupOK with default headers values
func NewCreateProcessGroupOK() *CreateProcessGroupOK {
	return &CreateProcessGroupOK{}
}

/*CreateProcessGroupOK handles this case with default header values.

successful operation
*/
type CreateProcessGroupOK struct {
	Payload *models.ProcessGroupEntity
}

func (o *CreateProcessGroupOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupOK  %+v", 200, o.Payload)
}

func (o *CreateProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProcessGroupBadRequest creates a CreateProcessGroupBadRequest with default headers values
func NewCreateProcessGroupBadRequest() *CreateProcessGroupBadRequest {
	return &CreateProcessGroupBadRequest{}
}

/*CreateProcessGroupBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateProcessGroupBadRequest struct {
}

func (o *CreateProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupBadRequest ", 400)
}

func (o *CreateProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProcessGroupUnauthorized creates a CreateProcessGroupUnauthorized with default headers values
func NewCreateProcessGroupUnauthorized() *CreateProcessGroupUnauthorized {
	return &CreateProcessGroupUnauthorized{}
}

/*CreateProcessGroupUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateProcessGroupUnauthorized struct {
}

func (o *CreateProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupUnauthorized ", 401)
}

func (o *CreateProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProcessGroupForbidden creates a CreateProcessGroupForbidden with default headers values
func NewCreateProcessGroupForbidden() *CreateProcessGroupForbidden {
	return &CreateProcessGroupForbidden{}
}

/*CreateProcessGroupForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateProcessGroupForbidden struct {
}

func (o *CreateProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupForbidden ", 403)
}

func (o *CreateProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProcessGroupNotFound creates a CreateProcessGroupNotFound with default headers values
func NewCreateProcessGroupNotFound() *CreateProcessGroupNotFound {
	return &CreateProcessGroupNotFound{}
}

/*CreateProcessGroupNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateProcessGroupNotFound struct {
}

func (o *CreateProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupNotFound ", 404)
}

func (o *CreateProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProcessGroupConflict creates a CreateProcessGroupConflict with default headers values
func NewCreateProcessGroupConflict() *CreateProcessGroupConflict {
	return &CreateProcessGroupConflict{}
}

/*CreateProcessGroupConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateProcessGroupConflict struct {
}

func (o *CreateProcessGroupConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups][%d] createProcessGroupConflict ", 409)
}

func (o *CreateProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
