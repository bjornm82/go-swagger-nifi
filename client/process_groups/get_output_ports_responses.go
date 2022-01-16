// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetOutputPortsReader is a Reader for the GetOutputPorts structure.
type GetOutputPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutputPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutputPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutputPortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutputPortsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutputPortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutputPortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetOutputPortsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOutputPortsOK creates a GetOutputPortsOK with default headers values
func NewGetOutputPortsOK() *GetOutputPortsOK {
	return &GetOutputPortsOK{}
}

/*GetOutputPortsOK handles this case with default header values.

successful operation
*/
type GetOutputPortsOK struct {
	Payload *models.OutputPortsEntity
}

func (o *GetOutputPortsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsOK  %+v", 200, o.Payload)
}

func (o *GetOutputPortsOK) GetPayload() *models.OutputPortsEntity {
	return o.Payload
}

func (o *GetOutputPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutputPortsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutputPortsBadRequest creates a GetOutputPortsBadRequest with default headers values
func NewGetOutputPortsBadRequest() *GetOutputPortsBadRequest {
	return &GetOutputPortsBadRequest{}
}

/*GetOutputPortsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetOutputPortsBadRequest struct {
}

func (o *GetOutputPortsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsBadRequest ", 400)
}

func (o *GetOutputPortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortsUnauthorized creates a GetOutputPortsUnauthorized with default headers values
func NewGetOutputPortsUnauthorized() *GetOutputPortsUnauthorized {
	return &GetOutputPortsUnauthorized{}
}

/*GetOutputPortsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetOutputPortsUnauthorized struct {
}

func (o *GetOutputPortsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsUnauthorized ", 401)
}

func (o *GetOutputPortsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortsForbidden creates a GetOutputPortsForbidden with default headers values
func NewGetOutputPortsForbidden() *GetOutputPortsForbidden {
	return &GetOutputPortsForbidden{}
}

/*GetOutputPortsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetOutputPortsForbidden struct {
}

func (o *GetOutputPortsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsForbidden ", 403)
}

func (o *GetOutputPortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortsNotFound creates a GetOutputPortsNotFound with default headers values
func NewGetOutputPortsNotFound() *GetOutputPortsNotFound {
	return &GetOutputPortsNotFound{}
}

/*GetOutputPortsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetOutputPortsNotFound struct {
}

func (o *GetOutputPortsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsNotFound ", 404)
}

func (o *GetOutputPortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortsConflict creates a GetOutputPortsConflict with default headers values
func NewGetOutputPortsConflict() *GetOutputPortsConflict {
	return &GetOutputPortsConflict{}
}

/*GetOutputPortsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetOutputPortsConflict struct {
}

func (o *GetOutputPortsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/output-ports][%d] getOutputPortsConflict ", 409)
}

func (o *GetOutputPortsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
