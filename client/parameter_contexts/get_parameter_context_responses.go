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

// GetParameterContextReader is a Reader for the GetParameterContext structure.
type GetParameterContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParameterContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParameterContextUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParameterContextForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetParameterContextNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParameterContextConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParameterContextOK creates a GetParameterContextOK with default headers values
func NewGetParameterContextOK() *GetParameterContextOK {
	return &GetParameterContextOK{}
}

/*GetParameterContextOK handles this case with default header values.

successful operation
*/
type GetParameterContextOK struct {
	Payload *models.ParameterContextEntity
}

func (o *GetParameterContextOK) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextOK) GetPayload() *models.ParameterContextEntity {
	return o.Payload
}

func (o *GetParameterContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterContextBadRequest creates a GetParameterContextBadRequest with default headers values
func NewGetParameterContextBadRequest() *GetParameterContextBadRequest {
	return &GetParameterContextBadRequest{}
}

/*GetParameterContextBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterContextBadRequest struct {
}

func (o *GetParameterContextBadRequest) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextBadRequest ", 400)
}

func (o *GetParameterContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUnauthorized creates a GetParameterContextUnauthorized with default headers values
func NewGetParameterContextUnauthorized() *GetParameterContextUnauthorized {
	return &GetParameterContextUnauthorized{}
}

/*GetParameterContextUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetParameterContextUnauthorized struct {
}

func (o *GetParameterContextUnauthorized) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextUnauthorized ", 401)
}

func (o *GetParameterContextUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextForbidden creates a GetParameterContextForbidden with default headers values
func NewGetParameterContextForbidden() *GetParameterContextForbidden {
	return &GetParameterContextForbidden{}
}

/*GetParameterContextForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetParameterContextForbidden struct {
}

func (o *GetParameterContextForbidden) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextForbidden ", 403)
}

func (o *GetParameterContextForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextNotFound creates a GetParameterContextNotFound with default headers values
func NewGetParameterContextNotFound() *GetParameterContextNotFound {
	return &GetParameterContextNotFound{}
}

/*GetParameterContextNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetParameterContextNotFound struct {
}

func (o *GetParameterContextNotFound) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextNotFound ", 404)
}

func (o *GetParameterContextNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextConflict creates a GetParameterContextConflict with default headers values
func NewGetParameterContextConflict() *GetParameterContextConflict {
	return &GetParameterContextConflict{}
}

/*GetParameterContextConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterContextConflict struct {
}

func (o *GetParameterContextConflict) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{id}][%d] getParameterContextConflict ", 409)
}

func (o *GetParameterContextConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
