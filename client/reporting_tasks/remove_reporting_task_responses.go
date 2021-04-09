// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// RemoveReportingTaskReader is a Reader for the RemoveReportingTask structure.
type RemoveReportingTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveReportingTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveReportingTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveReportingTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveReportingTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveReportingTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveReportingTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewRemoveReportingTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveReportingTaskOK creates a RemoveReportingTaskOK with default headers values
func NewRemoveReportingTaskOK() *RemoveReportingTaskOK {
	return &RemoveReportingTaskOK{}
}

/*RemoveReportingTaskOK handles this case with default header values.

successful operation
*/
type RemoveReportingTaskOK struct {
	Payload *models.ReportingTaskEntity
}

func (o *RemoveReportingTaskOK) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskOK  %+v", 200, o.Payload)
}

func (o *RemoveReportingTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingTaskEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveReportingTaskBadRequest creates a RemoveReportingTaskBadRequest with default headers values
func NewRemoveReportingTaskBadRequest() *RemoveReportingTaskBadRequest {
	return &RemoveReportingTaskBadRequest{}
}

/*RemoveReportingTaskBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveReportingTaskBadRequest struct {
}

func (o *RemoveReportingTaskBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskBadRequest ", 400)
}

func (o *RemoveReportingTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReportingTaskUnauthorized creates a RemoveReportingTaskUnauthorized with default headers values
func NewRemoveReportingTaskUnauthorized() *RemoveReportingTaskUnauthorized {
	return &RemoveReportingTaskUnauthorized{}
}

/*RemoveReportingTaskUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type RemoveReportingTaskUnauthorized struct {
}

func (o *RemoveReportingTaskUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskUnauthorized ", 401)
}

func (o *RemoveReportingTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReportingTaskForbidden creates a RemoveReportingTaskForbidden with default headers values
func NewRemoveReportingTaskForbidden() *RemoveReportingTaskForbidden {
	return &RemoveReportingTaskForbidden{}
}

/*RemoveReportingTaskForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type RemoveReportingTaskForbidden struct {
}

func (o *RemoveReportingTaskForbidden) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskForbidden ", 403)
}

func (o *RemoveReportingTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReportingTaskNotFound creates a RemoveReportingTaskNotFound with default headers values
func NewRemoveReportingTaskNotFound() *RemoveReportingTaskNotFound {
	return &RemoveReportingTaskNotFound{}
}

/*RemoveReportingTaskNotFound handles this case with default header values.

The specified resource could not be found.
*/
type RemoveReportingTaskNotFound struct {
}

func (o *RemoveReportingTaskNotFound) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskNotFound ", 404)
}

func (o *RemoveReportingTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReportingTaskConflict creates a RemoveReportingTaskConflict with default headers values
func NewRemoveReportingTaskConflict() *RemoveReportingTaskConflict {
	return &RemoveReportingTaskConflict{}
}

/*RemoveReportingTaskConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveReportingTaskConflict struct {
}

func (o *RemoveReportingTaskConflict) Error() string {
	return fmt.Sprintf("[DELETE /reporting-tasks/{id}][%d] removeReportingTaskConflict ", 409)
}

func (o *RemoveReportingTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
