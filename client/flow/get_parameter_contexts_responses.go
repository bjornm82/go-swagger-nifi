// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// GetParameterContextsReader is a Reader for the GetParameterContexts structure.
type GetParameterContextsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterContextsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterContextsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParameterContextsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParameterContextsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParameterContextsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParameterContextsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParameterContextsOK creates a GetParameterContextsOK with default headers values
func NewGetParameterContextsOK() *GetParameterContextsOK {
	return &GetParameterContextsOK{}
}

/*GetParameterContextsOK handles this case with default header values.

successful operation
*/
type GetParameterContextsOK struct {
	Payload *models.ParameterContextsEntity
}

func (o *GetParameterContextsOK) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextsOK) GetPayload() *models.ParameterContextsEntity {
	return o.Payload
}

func (o *GetParameterContextsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterContextsBadRequest creates a GetParameterContextsBadRequest with default headers values
func NewGetParameterContextsBadRequest() *GetParameterContextsBadRequest {
	return &GetParameterContextsBadRequest{}
}

/*GetParameterContextsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterContextsBadRequest struct {
}

func (o *GetParameterContextsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsBadRequest ", 400)
}

func (o *GetParameterContextsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsUnauthorized creates a GetParameterContextsUnauthorized with default headers values
func NewGetParameterContextsUnauthorized() *GetParameterContextsUnauthorized {
	return &GetParameterContextsUnauthorized{}
}

/*GetParameterContextsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetParameterContextsUnauthorized struct {
}

func (o *GetParameterContextsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsUnauthorized ", 401)
}

func (o *GetParameterContextsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsForbidden creates a GetParameterContextsForbidden with default headers values
func NewGetParameterContextsForbidden() *GetParameterContextsForbidden {
	return &GetParameterContextsForbidden{}
}

/*GetParameterContextsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetParameterContextsForbidden struct {
}

func (o *GetParameterContextsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsForbidden ", 403)
}

func (o *GetParameterContextsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsConflict creates a GetParameterContextsConflict with default headers values
func NewGetParameterContextsConflict() *GetParameterContextsConflict {
	return &GetParameterContextsConflict{}
}

/*GetParameterContextsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterContextsConflict struct {
}

func (o *GetParameterContextsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsConflict ", 409)
}

func (o *GetParameterContextsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
