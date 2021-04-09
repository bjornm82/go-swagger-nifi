// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// SubmitParameterContextUpdateReader is a Reader for the SubmitParameterContextUpdate structure.
type SubmitParameterContextUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitParameterContextUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSubmitParameterContextUpdateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSubmitParameterContextUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSubmitParameterContextUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSubmitParameterContextUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSubmitParameterContextUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewSubmitParameterContextUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitParameterContextUpdateCreated creates a SubmitParameterContextUpdateCreated with default headers values
func NewSubmitParameterContextUpdateCreated() *SubmitParameterContextUpdateCreated {
	return &SubmitParameterContextUpdateCreated{}
}

/*SubmitParameterContextUpdateCreated handles this case with default header values.

successful operation
*/
type SubmitParameterContextUpdateCreated struct {
	Payload *models.ParameterContextUpdateRequestEntity
}

func (o *SubmitParameterContextUpdateCreated) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateCreated  %+v", 201, o.Payload)
}

func (o *SubmitParameterContextUpdateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitParameterContextUpdateBadRequest creates a SubmitParameterContextUpdateBadRequest with default headers values
func NewSubmitParameterContextUpdateBadRequest() *SubmitParameterContextUpdateBadRequest {
	return &SubmitParameterContextUpdateBadRequest{}
}

/*SubmitParameterContextUpdateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SubmitParameterContextUpdateBadRequest struct {
}

func (o *SubmitParameterContextUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateBadRequest ", 400)
}

func (o *SubmitParameterContextUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitParameterContextUpdateUnauthorized creates a SubmitParameterContextUpdateUnauthorized with default headers values
func NewSubmitParameterContextUpdateUnauthorized() *SubmitParameterContextUpdateUnauthorized {
	return &SubmitParameterContextUpdateUnauthorized{}
}

/*SubmitParameterContextUpdateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type SubmitParameterContextUpdateUnauthorized struct {
}

func (o *SubmitParameterContextUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateUnauthorized ", 401)
}

func (o *SubmitParameterContextUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitParameterContextUpdateForbidden creates a SubmitParameterContextUpdateForbidden with default headers values
func NewSubmitParameterContextUpdateForbidden() *SubmitParameterContextUpdateForbidden {
	return &SubmitParameterContextUpdateForbidden{}
}

/*SubmitParameterContextUpdateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type SubmitParameterContextUpdateForbidden struct {
}

func (o *SubmitParameterContextUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateForbidden ", 403)
}

func (o *SubmitParameterContextUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitParameterContextUpdateNotFound creates a SubmitParameterContextUpdateNotFound with default headers values
func NewSubmitParameterContextUpdateNotFound() *SubmitParameterContextUpdateNotFound {
	return &SubmitParameterContextUpdateNotFound{}
}

/*SubmitParameterContextUpdateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type SubmitParameterContextUpdateNotFound struct {
}

func (o *SubmitParameterContextUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateNotFound ", 404)
}

func (o *SubmitParameterContextUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitParameterContextUpdateConflict creates a SubmitParameterContextUpdateConflict with default headers values
func NewSubmitParameterContextUpdateConflict() *SubmitParameterContextUpdateConflict {
	return &SubmitParameterContextUpdateConflict{}
}

/*SubmitParameterContextUpdateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SubmitParameterContextUpdateConflict struct {
}

func (o *SubmitParameterContextUpdateConflict) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts/{contextId}/update-requests][%d] submitParameterContextUpdateConflict ", 409)
}

func (o *SubmitParameterContextUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
