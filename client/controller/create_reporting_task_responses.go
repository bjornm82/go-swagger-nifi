// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CreateReportingTaskReader is a Reader for the CreateReportingTask structure.
type CreateReportingTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReportingTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateReportingTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReportingTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateReportingTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateReportingTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateReportingTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateReportingTaskOK creates a CreateReportingTaskOK with default headers values
func NewCreateReportingTaskOK() *CreateReportingTaskOK {
	return &CreateReportingTaskOK{}
}

/*CreateReportingTaskOK handles this case with default header values.

successful operation
*/
type CreateReportingTaskOK struct {
	Payload *models.ReportingTaskEntity
}

func (o *CreateReportingTaskOK) Error() string {
	return fmt.Sprintf("[POST /controller/reporting-tasks][%d] createReportingTaskOK  %+v", 200, o.Payload)
}

func (o *CreateReportingTaskOK) GetPayload() *models.ReportingTaskEntity {
	return o.Payload
}

func (o *CreateReportingTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportingTaskEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportingTaskBadRequest creates a CreateReportingTaskBadRequest with default headers values
func NewCreateReportingTaskBadRequest() *CreateReportingTaskBadRequest {
	return &CreateReportingTaskBadRequest{}
}

/*CreateReportingTaskBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateReportingTaskBadRequest struct {
}

func (o *CreateReportingTaskBadRequest) Error() string {
	return fmt.Sprintf("[POST /controller/reporting-tasks][%d] createReportingTaskBadRequest ", 400)
}

func (o *CreateReportingTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateReportingTaskUnauthorized creates a CreateReportingTaskUnauthorized with default headers values
func NewCreateReportingTaskUnauthorized() *CreateReportingTaskUnauthorized {
	return &CreateReportingTaskUnauthorized{}
}

/*CreateReportingTaskUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateReportingTaskUnauthorized struct {
}

func (o *CreateReportingTaskUnauthorized) Error() string {
	return fmt.Sprintf("[POST /controller/reporting-tasks][%d] createReportingTaskUnauthorized ", 401)
}

func (o *CreateReportingTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateReportingTaskForbidden creates a CreateReportingTaskForbidden with default headers values
func NewCreateReportingTaskForbidden() *CreateReportingTaskForbidden {
	return &CreateReportingTaskForbidden{}
}

/*CreateReportingTaskForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateReportingTaskForbidden struct {
}

func (o *CreateReportingTaskForbidden) Error() string {
	return fmt.Sprintf("[POST /controller/reporting-tasks][%d] createReportingTaskForbidden ", 403)
}

func (o *CreateReportingTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateReportingTaskConflict creates a CreateReportingTaskConflict with default headers values
func NewCreateReportingTaskConflict() *CreateReportingTaskConflict {
	return &CreateReportingTaskConflict{}
}

/*CreateReportingTaskConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateReportingTaskConflict struct {
}

func (o *CreateReportingTaskConflict) Error() string {
	return fmt.Sprintf("[POST /controller/reporting-tasks][%d] createReportingTaskConflict ", 409)
}

func (o *CreateReportingTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
