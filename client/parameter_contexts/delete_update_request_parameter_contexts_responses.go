// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// DeleteUpdateRequestParameterContextsReader is a Reader for the DeleteUpdateRequestParameterContexts structure.
type DeleteUpdateRequestParameterContextsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUpdateRequestParameterContextsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUpdateRequestParameterContextsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUpdateRequestParameterContextsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUpdateRequestParameterContextsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUpdateRequestParameterContextsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUpdateRequestParameterContextsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteUpdateRequestParameterContextsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUpdateRequestParameterContextsOK creates a DeleteUpdateRequestParameterContextsOK with default headers values
func NewDeleteUpdateRequestParameterContextsOK() *DeleteUpdateRequestParameterContextsOK {
	return &DeleteUpdateRequestParameterContextsOK{}
}

/*DeleteUpdateRequestParameterContextsOK handles this case with default header values.

successful operation
*/
type DeleteUpdateRequestParameterContextsOK struct {
	Payload *models.ParameterContextUpdateRequestEntity
}

func (o *DeleteUpdateRequestParameterContextsOK) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsOK  %+v", 200, o.Payload)
}

func (o *DeleteUpdateRequestParameterContextsOK) GetPayload() *models.ParameterContextUpdateRequestEntity {
	return o.Payload
}

func (o *DeleteUpdateRequestParameterContextsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUpdateRequestParameterContextsBadRequest creates a DeleteUpdateRequestParameterContextsBadRequest with default headers values
func NewDeleteUpdateRequestParameterContextsBadRequest() *DeleteUpdateRequestParameterContextsBadRequest {
	return &DeleteUpdateRequestParameterContextsBadRequest{}
}

/*DeleteUpdateRequestParameterContextsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteUpdateRequestParameterContextsBadRequest struct {
}

func (o *DeleteUpdateRequestParameterContextsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsBadRequest ", 400)
}

func (o *DeleteUpdateRequestParameterContextsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestParameterContextsUnauthorized creates a DeleteUpdateRequestParameterContextsUnauthorized with default headers values
func NewDeleteUpdateRequestParameterContextsUnauthorized() *DeleteUpdateRequestParameterContextsUnauthorized {
	return &DeleteUpdateRequestParameterContextsUnauthorized{}
}

/*DeleteUpdateRequestParameterContextsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteUpdateRequestParameterContextsUnauthorized struct {
}

func (o *DeleteUpdateRequestParameterContextsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsUnauthorized ", 401)
}

func (o *DeleteUpdateRequestParameterContextsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestParameterContextsForbidden creates a DeleteUpdateRequestParameterContextsForbidden with default headers values
func NewDeleteUpdateRequestParameterContextsForbidden() *DeleteUpdateRequestParameterContextsForbidden {
	return &DeleteUpdateRequestParameterContextsForbidden{}
}

/*DeleteUpdateRequestParameterContextsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteUpdateRequestParameterContextsForbidden struct {
}

func (o *DeleteUpdateRequestParameterContextsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsForbidden ", 403)
}

func (o *DeleteUpdateRequestParameterContextsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestParameterContextsNotFound creates a DeleteUpdateRequestParameterContextsNotFound with default headers values
func NewDeleteUpdateRequestParameterContextsNotFound() *DeleteUpdateRequestParameterContextsNotFound {
	return &DeleteUpdateRequestParameterContextsNotFound{}
}

/*DeleteUpdateRequestParameterContextsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteUpdateRequestParameterContextsNotFound struct {
}

func (o *DeleteUpdateRequestParameterContextsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsNotFound ", 404)
}

func (o *DeleteUpdateRequestParameterContextsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestParameterContextsConflict creates a DeleteUpdateRequestParameterContextsConflict with default headers values
func NewDeleteUpdateRequestParameterContextsConflict() *DeleteUpdateRequestParameterContextsConflict {
	return &DeleteUpdateRequestParameterContextsConflict{}
}

/*DeleteUpdateRequestParameterContextsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteUpdateRequestParameterContextsConflict struct {
}

func (o *DeleteUpdateRequestParameterContextsConflict) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/update-requests/{requestId}][%d] deleteUpdateRequestParameterContextsConflict ", 409)
}

func (o *DeleteUpdateRequestParameterContextsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}