// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// DeleteProcessorReader is a Reader for the DeleteProcessor structure.
type DeleteProcessorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProcessorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProcessorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteProcessorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteProcessorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteProcessorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteProcessorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteProcessorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProcessorOK creates a DeleteProcessorOK with default headers values
func NewDeleteProcessorOK() *DeleteProcessorOK {
	return &DeleteProcessorOK{}
}

/*DeleteProcessorOK handles this case with default header values.

successful operation
*/
type DeleteProcessorOK struct {
	Payload *models.ProcessorEntity
}

func (o *DeleteProcessorOK) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorOK  %+v", 200, o.Payload)
}

func (o *DeleteProcessorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProcessorBadRequest creates a DeleteProcessorBadRequest with default headers values
func NewDeleteProcessorBadRequest() *DeleteProcessorBadRequest {
	return &DeleteProcessorBadRequest{}
}

/*DeleteProcessorBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteProcessorBadRequest struct {
}

func (o *DeleteProcessorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorBadRequest ", 400)
}

func (o *DeleteProcessorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProcessorUnauthorized creates a DeleteProcessorUnauthorized with default headers values
func NewDeleteProcessorUnauthorized() *DeleteProcessorUnauthorized {
	return &DeleteProcessorUnauthorized{}
}

/*DeleteProcessorUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteProcessorUnauthorized struct {
}

func (o *DeleteProcessorUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorUnauthorized ", 401)
}

func (o *DeleteProcessorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProcessorForbidden creates a DeleteProcessorForbidden with default headers values
func NewDeleteProcessorForbidden() *DeleteProcessorForbidden {
	return &DeleteProcessorForbidden{}
}

/*DeleteProcessorForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteProcessorForbidden struct {
}

func (o *DeleteProcessorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorForbidden ", 403)
}

func (o *DeleteProcessorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProcessorNotFound creates a DeleteProcessorNotFound with default headers values
func NewDeleteProcessorNotFound() *DeleteProcessorNotFound {
	return &DeleteProcessorNotFound{}
}

/*DeleteProcessorNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteProcessorNotFound struct {
}

func (o *DeleteProcessorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorNotFound ", 404)
}

func (o *DeleteProcessorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProcessorConflict creates a DeleteProcessorConflict with default headers values
func NewDeleteProcessorConflict() *DeleteProcessorConflict {
	return &DeleteProcessorConflict{}
}

/*DeleteProcessorConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteProcessorConflict struct {
}

func (o *DeleteProcessorConflict) Error() string {
	return fmt.Sprintf("[DELETE /processors/{id}][%d] deleteProcessorConflict ", 409)
}

func (o *DeleteProcessorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
