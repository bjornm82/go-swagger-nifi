// Code generated by go-swagger; DO NOT EDIT.

package funnel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetFunnelReader is a Reader for the GetFunnel structure.
type GetFunnelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunnelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFunnelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFunnelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFunnelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFunnelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFunnelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFunnelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFunnelOK creates a GetFunnelOK with default headers values
func NewGetFunnelOK() *GetFunnelOK {
	return &GetFunnelOK{}
}

/*GetFunnelOK handles this case with default header values.

successful operation
*/
type GetFunnelOK struct {
	Payload *models.FunnelEntity
}

func (o *GetFunnelOK) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelOK  %+v", 200, o.Payload)
}

func (o *GetFunnelOK) GetPayload() *models.FunnelEntity {
	return o.Payload
}

func (o *GetFunnelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunnelBadRequest creates a GetFunnelBadRequest with default headers values
func NewGetFunnelBadRequest() *GetFunnelBadRequest {
	return &GetFunnelBadRequest{}
}

/*GetFunnelBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFunnelBadRequest struct {
}

func (o *GetFunnelBadRequest) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelBadRequest ", 400)
}

func (o *GetFunnelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelUnauthorized creates a GetFunnelUnauthorized with default headers values
func NewGetFunnelUnauthorized() *GetFunnelUnauthorized {
	return &GetFunnelUnauthorized{}
}

/*GetFunnelUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetFunnelUnauthorized struct {
}

func (o *GetFunnelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelUnauthorized ", 401)
}

func (o *GetFunnelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelForbidden creates a GetFunnelForbidden with default headers values
func NewGetFunnelForbidden() *GetFunnelForbidden {
	return &GetFunnelForbidden{}
}

/*GetFunnelForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetFunnelForbidden struct {
}

func (o *GetFunnelForbidden) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelForbidden ", 403)
}

func (o *GetFunnelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelNotFound creates a GetFunnelNotFound with default headers values
func NewGetFunnelNotFound() *GetFunnelNotFound {
	return &GetFunnelNotFound{}
}

/*GetFunnelNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetFunnelNotFound struct {
}

func (o *GetFunnelNotFound) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelNotFound ", 404)
}

func (o *GetFunnelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelConflict creates a GetFunnelConflict with default headers values
func NewGetFunnelConflict() *GetFunnelConflict {
	return &GetFunnelConflict{}
}

/*GetFunnelConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFunnelConflict struct {
}

func (o *GetFunnelConflict) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelConflict ", 409)
}

func (o *GetFunnelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
