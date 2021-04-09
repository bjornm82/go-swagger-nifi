// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetParameterContextUpdateReader is a Reader for the GetParameterContextUpdate structure.
type GetParameterContextUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterContextUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParameterContextUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetParameterContextUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetParameterContextUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetParameterContextUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetParameterContextUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetParameterContextUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParameterContextUpdateOK creates a GetParameterContextUpdateOK with default headers values
func NewGetParameterContextUpdateOK() *GetParameterContextUpdateOK {
	return &GetParameterContextUpdateOK{}
}

/*GetParameterContextUpdateOK handles this case with default header values.

successful operation
*/
type GetParameterContextUpdateOK struct {
	Payload *models.ParameterContextUpdateRequestEntity
}

func (o *GetParameterContextUpdateOK) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterContextUpdateBadRequest creates a GetParameterContextUpdateBadRequest with default headers values
func NewGetParameterContextUpdateBadRequest() *GetParameterContextUpdateBadRequest {
	return &GetParameterContextUpdateBadRequest{}
}

/*GetParameterContextUpdateBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterContextUpdateBadRequest struct {
}

func (o *GetParameterContextUpdateBadRequest) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateBadRequest ", 400)
}

func (o *GetParameterContextUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateUnauthorized creates a GetParameterContextUpdateUnauthorized with default headers values
func NewGetParameterContextUpdateUnauthorized() *GetParameterContextUpdateUnauthorized {
	return &GetParameterContextUpdateUnauthorized{}
}

/*GetParameterContextUpdateUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetParameterContextUpdateUnauthorized struct {
}

func (o *GetParameterContextUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateUnauthorized ", 401)
}

func (o *GetParameterContextUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateForbidden creates a GetParameterContextUpdateForbidden with default headers values
func NewGetParameterContextUpdateForbidden() *GetParameterContextUpdateForbidden {
	return &GetParameterContextUpdateForbidden{}
}

/*GetParameterContextUpdateForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetParameterContextUpdateForbidden struct {
}

func (o *GetParameterContextUpdateForbidden) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateForbidden ", 403)
}

func (o *GetParameterContextUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateNotFound creates a GetParameterContextUpdateNotFound with default headers values
func NewGetParameterContextUpdateNotFound() *GetParameterContextUpdateNotFound {
	return &GetParameterContextUpdateNotFound{}
}

/*GetParameterContextUpdateNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetParameterContextUpdateNotFound struct {
}

func (o *GetParameterContextUpdateNotFound) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateNotFound ", 404)
}

func (o *GetParameterContextUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateConflict creates a GetParameterContextUpdateConflict with default headers values
func NewGetParameterContextUpdateConflict() *GetParameterContextUpdateConflict {
	return &GetParameterContextUpdateConflict{}
}

/*GetParameterContextUpdateConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterContextUpdateConflict struct {
}

func (o *GetParameterContextUpdateConflict) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateConflict ", 409)
}

func (o *GetParameterContextUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
