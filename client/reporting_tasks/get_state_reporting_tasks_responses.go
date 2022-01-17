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

// GetStateReportingTasksReader is a Reader for the GetStateReportingTasks structure.
type GetStateReportingTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStateReportingTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStateReportingTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStateReportingTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStateReportingTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStateReportingTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStateReportingTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetStateReportingTasksConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStateReportingTasksOK creates a GetStateReportingTasksOK with default headers values
func NewGetStateReportingTasksOK() *GetStateReportingTasksOK {
	return &GetStateReportingTasksOK{}
}

/*GetStateReportingTasksOK handles this case with default header values.

successful operation
*/
type GetStateReportingTasksOK struct {
	Payload *models.ComponentStateEntity
}

func (o *GetStateReportingTasksOK) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksOK  %+v", 200, o.Payload)
}

func (o *GetStateReportingTasksOK) GetPayload() *models.ComponentStateEntity {
	return o.Payload
}

func (o *GetStateReportingTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStateReportingTasksBadRequest creates a GetStateReportingTasksBadRequest with default headers values
func NewGetStateReportingTasksBadRequest() *GetStateReportingTasksBadRequest {
	return &GetStateReportingTasksBadRequest{}
}

/*GetStateReportingTasksBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetStateReportingTasksBadRequest struct {
}

func (o *GetStateReportingTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksBadRequest ", 400)
}

func (o *GetStateReportingTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateReportingTasksUnauthorized creates a GetStateReportingTasksUnauthorized with default headers values
func NewGetStateReportingTasksUnauthorized() *GetStateReportingTasksUnauthorized {
	return &GetStateReportingTasksUnauthorized{}
}

/*GetStateReportingTasksUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetStateReportingTasksUnauthorized struct {
}

func (o *GetStateReportingTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksUnauthorized ", 401)
}

func (o *GetStateReportingTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateReportingTasksForbidden creates a GetStateReportingTasksForbidden with default headers values
func NewGetStateReportingTasksForbidden() *GetStateReportingTasksForbidden {
	return &GetStateReportingTasksForbidden{}
}

/*GetStateReportingTasksForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetStateReportingTasksForbidden struct {
}

func (o *GetStateReportingTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksForbidden ", 403)
}

func (o *GetStateReportingTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateReportingTasksNotFound creates a GetStateReportingTasksNotFound with default headers values
func NewGetStateReportingTasksNotFound() *GetStateReportingTasksNotFound {
	return &GetStateReportingTasksNotFound{}
}

/*GetStateReportingTasksNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetStateReportingTasksNotFound struct {
}

func (o *GetStateReportingTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksNotFound ", 404)
}

func (o *GetStateReportingTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateReportingTasksConflict creates a GetStateReportingTasksConflict with default headers values
func NewGetStateReportingTasksConflict() *GetStateReportingTasksConflict {
	return &GetStateReportingTasksConflict{}
}

/*GetStateReportingTasksConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetStateReportingTasksConflict struct {
}

func (o *GetStateReportingTasksConflict) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] getStateReportingTasksConflict ", 409)
}

func (o *GetStateReportingTasksConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}