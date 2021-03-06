// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetPropertyDescriptorReportingTasksReader is a Reader for the GetPropertyDescriptorReportingTasks structure.
type GetPropertyDescriptorReportingTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPropertyDescriptorReportingTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPropertyDescriptorReportingTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPropertyDescriptorReportingTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPropertyDescriptorReportingTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPropertyDescriptorReportingTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPropertyDescriptorReportingTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPropertyDescriptorReportingTasksConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPropertyDescriptorReportingTasksOK creates a GetPropertyDescriptorReportingTasksOK with default headers values
func NewGetPropertyDescriptorReportingTasksOK() *GetPropertyDescriptorReportingTasksOK {
	return &GetPropertyDescriptorReportingTasksOK{}
}

/*GetPropertyDescriptorReportingTasksOK handles this case with default header values.

successful operation
*/
type GetPropertyDescriptorReportingTasksOK struct {
	Payload *models.PropertyDescriptorEntity
}

func (o *GetPropertyDescriptorReportingTasksOK) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksOK  %+v", 200, o.Payload)
}

func (o *GetPropertyDescriptorReportingTasksOK) GetPayload() *models.PropertyDescriptorEntity {
	return o.Payload
}

func (o *GetPropertyDescriptorReportingTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyDescriptorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPropertyDescriptorReportingTasksBadRequest creates a GetPropertyDescriptorReportingTasksBadRequest with default headers values
func NewGetPropertyDescriptorReportingTasksBadRequest() *GetPropertyDescriptorReportingTasksBadRequest {
	return &GetPropertyDescriptorReportingTasksBadRequest{}
}

/*GetPropertyDescriptorReportingTasksBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetPropertyDescriptorReportingTasksBadRequest struct {
}

func (o *GetPropertyDescriptorReportingTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksBadRequest ", 400)
}

func (o *GetPropertyDescriptorReportingTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyDescriptorReportingTasksUnauthorized creates a GetPropertyDescriptorReportingTasksUnauthorized with default headers values
func NewGetPropertyDescriptorReportingTasksUnauthorized() *GetPropertyDescriptorReportingTasksUnauthorized {
	return &GetPropertyDescriptorReportingTasksUnauthorized{}
}

/*GetPropertyDescriptorReportingTasksUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetPropertyDescriptorReportingTasksUnauthorized struct {
}

func (o *GetPropertyDescriptorReportingTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksUnauthorized ", 401)
}

func (o *GetPropertyDescriptorReportingTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyDescriptorReportingTasksForbidden creates a GetPropertyDescriptorReportingTasksForbidden with default headers values
func NewGetPropertyDescriptorReportingTasksForbidden() *GetPropertyDescriptorReportingTasksForbidden {
	return &GetPropertyDescriptorReportingTasksForbidden{}
}

/*GetPropertyDescriptorReportingTasksForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetPropertyDescriptorReportingTasksForbidden struct {
}

func (o *GetPropertyDescriptorReportingTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksForbidden ", 403)
}

func (o *GetPropertyDescriptorReportingTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyDescriptorReportingTasksNotFound creates a GetPropertyDescriptorReportingTasksNotFound with default headers values
func NewGetPropertyDescriptorReportingTasksNotFound() *GetPropertyDescriptorReportingTasksNotFound {
	return &GetPropertyDescriptorReportingTasksNotFound{}
}

/*GetPropertyDescriptorReportingTasksNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetPropertyDescriptorReportingTasksNotFound struct {
}

func (o *GetPropertyDescriptorReportingTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksNotFound ", 404)
}

func (o *GetPropertyDescriptorReportingTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyDescriptorReportingTasksConflict creates a GetPropertyDescriptorReportingTasksConflict with default headers values
func NewGetPropertyDescriptorReportingTasksConflict() *GetPropertyDescriptorReportingTasksConflict {
	return &GetPropertyDescriptorReportingTasksConflict{}
}

/*GetPropertyDescriptorReportingTasksConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetPropertyDescriptorReportingTasksConflict struct {
}

func (o *GetPropertyDescriptorReportingTasksConflict) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/descriptors][%d] getPropertyDescriptorReportingTasksConflict ", 409)
}

func (o *GetPropertyDescriptorReportingTasksConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
