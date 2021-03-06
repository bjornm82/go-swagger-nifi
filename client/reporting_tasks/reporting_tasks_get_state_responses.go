// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// ReportingTasksGetStateReader is a Reader for the ReportingTasksGetState structure.
type ReportingTasksGetStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportingTasksGetStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReportingTasksGetStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReportingTasksGetStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewReportingTasksGetStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReportingTasksGetStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReportingTasksGetStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewReportingTasksGetStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReportingTasksGetStateOK creates a ReportingTasksGetStateOK with default headers values
func NewReportingTasksGetStateOK() *ReportingTasksGetStateOK {
	return &ReportingTasksGetStateOK{}
}

/*ReportingTasksGetStateOK handles this case with default header values.

successful operation
*/
type ReportingTasksGetStateOK struct {
	Payload *models.ComponentStateEntity
}

func (o *ReportingTasksGetStateOK) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateOK  %+v", 200, o.Payload)
}

func (o *ReportingTasksGetStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportingTasksGetStateBadRequest creates a ReportingTasksGetStateBadRequest with default headers values
func NewReportingTasksGetStateBadRequest() *ReportingTasksGetStateBadRequest {
	return &ReportingTasksGetStateBadRequest{}
}

/*ReportingTasksGetStateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ReportingTasksGetStateBadRequest struct {
}

func (o *ReportingTasksGetStateBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateBadRequest ", 400)
}

func (o *ReportingTasksGetStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportingTasksGetStateUnauthorized creates a ReportingTasksGetStateUnauthorized with default headers values
func NewReportingTasksGetStateUnauthorized() *ReportingTasksGetStateUnauthorized {
	return &ReportingTasksGetStateUnauthorized{}
}

/*ReportingTasksGetStateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type ReportingTasksGetStateUnauthorized struct {
}

func (o *ReportingTasksGetStateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateUnauthorized ", 401)
}

func (o *ReportingTasksGetStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportingTasksGetStateForbidden creates a ReportingTasksGetStateForbidden with default headers values
func NewReportingTasksGetStateForbidden() *ReportingTasksGetStateForbidden {
	return &ReportingTasksGetStateForbidden{}
}

/*ReportingTasksGetStateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type ReportingTasksGetStateForbidden struct {
}

func (o *ReportingTasksGetStateForbidden) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateForbidden ", 403)
}

func (o *ReportingTasksGetStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportingTasksGetStateNotFound creates a ReportingTasksGetStateNotFound with default headers values
func NewReportingTasksGetStateNotFound() *ReportingTasksGetStateNotFound {
	return &ReportingTasksGetStateNotFound{}
}

/*ReportingTasksGetStateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type ReportingTasksGetStateNotFound struct {
}

func (o *ReportingTasksGetStateNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateNotFound ", 404)
}

func (o *ReportingTasksGetStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportingTasksGetStateConflict creates a ReportingTasksGetStateConflict with default headers values
func NewReportingTasksGetStateConflict() *ReportingTasksGetStateConflict {
	return &ReportingTasksGetStateConflict{}
}

/*ReportingTasksGetStateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ReportingTasksGetStateConflict struct {
}

func (o *ReportingTasksGetStateConflict) Error() string {
	return fmt.Sprintf("[GET /reporting-tasks/{id}/state][%d] reportingTasksGetStateConflict ", 409)
}

func (o *ReportingTasksGetStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
