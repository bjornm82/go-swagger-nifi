// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetProcessorRunStatusDetailsReader is a Reader for the GetProcessorRunStatusDetails structure.
type GetProcessorRunStatusDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorRunStatusDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetProcessorRunStatusDetailsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessorRunStatusDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessorRunStatusDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessorRunStatusDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessorRunStatusDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessorRunStatusDetailsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProcessorRunStatusDetailsCreated creates a GetProcessorRunStatusDetailsCreated with default headers values
func NewGetProcessorRunStatusDetailsCreated() *GetProcessorRunStatusDetailsCreated {
	return &GetProcessorRunStatusDetailsCreated{}
}

/*GetProcessorRunStatusDetailsCreated handles this case with default header values.

successful operation
*/
type GetProcessorRunStatusDetailsCreated struct {
	Payload *models.ProcessorsRunStatusDetailsEntity
}

func (o *GetProcessorRunStatusDetailsCreated) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsCreated  %+v", 201, o.Payload)
}

func (o *GetProcessorRunStatusDetailsCreated) GetPayload() *models.ProcessorsRunStatusDetailsEntity {
	return o.Payload
}

func (o *GetProcessorRunStatusDetailsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorsRunStatusDetailsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorRunStatusDetailsBadRequest creates a GetProcessorRunStatusDetailsBadRequest with default headers values
func NewGetProcessorRunStatusDetailsBadRequest() *GetProcessorRunStatusDetailsBadRequest {
	return &GetProcessorRunStatusDetailsBadRequest{}
}

/*GetProcessorRunStatusDetailsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorRunStatusDetailsBadRequest struct {
}

func (o *GetProcessorRunStatusDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsBadRequest ", 400)
}

func (o *GetProcessorRunStatusDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorRunStatusDetailsUnauthorized creates a GetProcessorRunStatusDetailsUnauthorized with default headers values
func NewGetProcessorRunStatusDetailsUnauthorized() *GetProcessorRunStatusDetailsUnauthorized {
	return &GetProcessorRunStatusDetailsUnauthorized{}
}

/*GetProcessorRunStatusDetailsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetProcessorRunStatusDetailsUnauthorized struct {
}

func (o *GetProcessorRunStatusDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsUnauthorized ", 401)
}

func (o *GetProcessorRunStatusDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorRunStatusDetailsForbidden creates a GetProcessorRunStatusDetailsForbidden with default headers values
func NewGetProcessorRunStatusDetailsForbidden() *GetProcessorRunStatusDetailsForbidden {
	return &GetProcessorRunStatusDetailsForbidden{}
}

/*GetProcessorRunStatusDetailsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetProcessorRunStatusDetailsForbidden struct {
}

func (o *GetProcessorRunStatusDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsForbidden ", 403)
}

func (o *GetProcessorRunStatusDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorRunStatusDetailsNotFound creates a GetProcessorRunStatusDetailsNotFound with default headers values
func NewGetProcessorRunStatusDetailsNotFound() *GetProcessorRunStatusDetailsNotFound {
	return &GetProcessorRunStatusDetailsNotFound{}
}

/*GetProcessorRunStatusDetailsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetProcessorRunStatusDetailsNotFound struct {
}

func (o *GetProcessorRunStatusDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsNotFound ", 404)
}

func (o *GetProcessorRunStatusDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorRunStatusDetailsConflict creates a GetProcessorRunStatusDetailsConflict with default headers values
func NewGetProcessorRunStatusDetailsConflict() *GetProcessorRunStatusDetailsConflict {
	return &GetProcessorRunStatusDetailsConflict{}
}

/*GetProcessorRunStatusDetailsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorRunStatusDetailsConflict struct {
}

func (o *GetProcessorRunStatusDetailsConflict) Error() string {
	return fmt.Sprintf("[POST /processors/run-status-details/queries][%d] getProcessorRunStatusDetailsConflict ", 409)
}

func (o *GetProcessorRunStatusDetailsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
