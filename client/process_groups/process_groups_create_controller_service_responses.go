// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// ProcessGroupsCreateControllerServiceReader is a Reader for the ProcessGroupsCreateControllerService structure.
type ProcessGroupsCreateControllerServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProcessGroupsCreateControllerServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProcessGroupsCreateControllerServiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProcessGroupsCreateControllerServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewProcessGroupsCreateControllerServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewProcessGroupsCreateControllerServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewProcessGroupsCreateControllerServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProcessGroupsCreateControllerServiceCreated creates a ProcessGroupsCreateControllerServiceCreated with default headers values
func NewProcessGroupsCreateControllerServiceCreated() *ProcessGroupsCreateControllerServiceCreated {
	return &ProcessGroupsCreateControllerServiceCreated{}
}

/*ProcessGroupsCreateControllerServiceCreated handles this case with default header values.

successful operation
*/
type ProcessGroupsCreateControllerServiceCreated struct {
	Payload *models.ControllerServiceEntity
}

func (o *ProcessGroupsCreateControllerServiceCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/controller-services][%d] processGroupsCreateControllerServiceCreated  %+v", 201, o.Payload)
}

func (o *ProcessGroupsCreateControllerServiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProcessGroupsCreateControllerServiceBadRequest creates a ProcessGroupsCreateControllerServiceBadRequest with default headers values
func NewProcessGroupsCreateControllerServiceBadRequest() *ProcessGroupsCreateControllerServiceBadRequest {
	return &ProcessGroupsCreateControllerServiceBadRequest{}
}

/*ProcessGroupsCreateControllerServiceBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ProcessGroupsCreateControllerServiceBadRequest struct {
}

func (o *ProcessGroupsCreateControllerServiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/controller-services][%d] processGroupsCreateControllerServiceBadRequest ", 400)
}

func (o *ProcessGroupsCreateControllerServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessGroupsCreateControllerServiceUnauthorized creates a ProcessGroupsCreateControllerServiceUnauthorized with default headers values
func NewProcessGroupsCreateControllerServiceUnauthorized() *ProcessGroupsCreateControllerServiceUnauthorized {
	return &ProcessGroupsCreateControllerServiceUnauthorized{}
}

/*ProcessGroupsCreateControllerServiceUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type ProcessGroupsCreateControllerServiceUnauthorized struct {
}

func (o *ProcessGroupsCreateControllerServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/controller-services][%d] processGroupsCreateControllerServiceUnauthorized ", 401)
}

func (o *ProcessGroupsCreateControllerServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessGroupsCreateControllerServiceForbidden creates a ProcessGroupsCreateControllerServiceForbidden with default headers values
func NewProcessGroupsCreateControllerServiceForbidden() *ProcessGroupsCreateControllerServiceForbidden {
	return &ProcessGroupsCreateControllerServiceForbidden{}
}

/*ProcessGroupsCreateControllerServiceForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type ProcessGroupsCreateControllerServiceForbidden struct {
}

func (o *ProcessGroupsCreateControllerServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/controller-services][%d] processGroupsCreateControllerServiceForbidden ", 403)
}

func (o *ProcessGroupsCreateControllerServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessGroupsCreateControllerServiceConflict creates a ProcessGroupsCreateControllerServiceConflict with default headers values
func NewProcessGroupsCreateControllerServiceConflict() *ProcessGroupsCreateControllerServiceConflict {
	return &ProcessGroupsCreateControllerServiceConflict{}
}

/*ProcessGroupsCreateControllerServiceConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ProcessGroupsCreateControllerServiceConflict struct {
}

func (o *ProcessGroupsCreateControllerServiceConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/controller-services][%d] processGroupsCreateControllerServiceConflict ", 409)
}

func (o *ProcessGroupsCreateControllerServiceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
