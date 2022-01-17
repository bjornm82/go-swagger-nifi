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

// GetStateProcessorsReader is a Reader for the GetStateProcessors structure.
type GetStateProcessorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStateProcessorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStateProcessorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStateProcessorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStateProcessorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStateProcessorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStateProcessorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetStateProcessorsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStateProcessorsOK creates a GetStateProcessorsOK with default headers values
func NewGetStateProcessorsOK() *GetStateProcessorsOK {
	return &GetStateProcessorsOK{}
}

/*GetStateProcessorsOK handles this case with default header values.

successful operation
*/
type GetStateProcessorsOK struct {
	Payload *models.ComponentStateEntity
}

func (o *GetStateProcessorsOK) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsOK  %+v", 200, o.Payload)
}

func (o *GetStateProcessorsOK) GetPayload() *models.ComponentStateEntity {
	return o.Payload
}

func (o *GetStateProcessorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStateProcessorsBadRequest creates a GetStateProcessorsBadRequest with default headers values
func NewGetStateProcessorsBadRequest() *GetStateProcessorsBadRequest {
	return &GetStateProcessorsBadRequest{}
}

/*GetStateProcessorsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetStateProcessorsBadRequest struct {
}

func (o *GetStateProcessorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsBadRequest ", 400)
}

func (o *GetStateProcessorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateProcessorsUnauthorized creates a GetStateProcessorsUnauthorized with default headers values
func NewGetStateProcessorsUnauthorized() *GetStateProcessorsUnauthorized {
	return &GetStateProcessorsUnauthorized{}
}

/*GetStateProcessorsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetStateProcessorsUnauthorized struct {
}

func (o *GetStateProcessorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsUnauthorized ", 401)
}

func (o *GetStateProcessorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateProcessorsForbidden creates a GetStateProcessorsForbidden with default headers values
func NewGetStateProcessorsForbidden() *GetStateProcessorsForbidden {
	return &GetStateProcessorsForbidden{}
}

/*GetStateProcessorsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetStateProcessorsForbidden struct {
}

func (o *GetStateProcessorsForbidden) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsForbidden ", 403)
}

func (o *GetStateProcessorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateProcessorsNotFound creates a GetStateProcessorsNotFound with default headers values
func NewGetStateProcessorsNotFound() *GetStateProcessorsNotFound {
	return &GetStateProcessorsNotFound{}
}

/*GetStateProcessorsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetStateProcessorsNotFound struct {
}

func (o *GetStateProcessorsNotFound) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsNotFound ", 404)
}

func (o *GetStateProcessorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStateProcessorsConflict creates a GetStateProcessorsConflict with default headers values
func NewGetStateProcessorsConflict() *GetStateProcessorsConflict {
	return &GetStateProcessorsConflict{}
}

/*GetStateProcessorsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetStateProcessorsConflict struct {
}

func (o *GetStateProcessorsConflict) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/state][%d] getStateProcessorsConflict ", 409)
}

func (o *GetStateProcessorsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}