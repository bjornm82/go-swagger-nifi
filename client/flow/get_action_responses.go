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

// GetActionReader is a Reader for the GetAction structure.
type GetActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActionOK creates a GetActionOK with default headers values
func NewGetActionOK() *GetActionOK {
	return &GetActionOK{}
}

/*GetActionOK handles this case with default header values.

successful operation
*/
type GetActionOK struct {
	Payload *models.ActionEntity
}

func (o *GetActionOK) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionOK  %+v", 200, o.Payload)
}

func (o *GetActionOK) GetPayload() *models.ActionEntity {
	return o.Payload
}

func (o *GetActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionBadRequest creates a GetActionBadRequest with default headers values
func NewGetActionBadRequest() *GetActionBadRequest {
	return &GetActionBadRequest{}
}

/*GetActionBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetActionBadRequest struct {
}

func (o *GetActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionBadRequest ", 400)
}

func (o *GetActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionUnauthorized creates a GetActionUnauthorized with default headers values
func NewGetActionUnauthorized() *GetActionUnauthorized {
	return &GetActionUnauthorized{}
}

/*GetActionUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetActionUnauthorized struct {
}

func (o *GetActionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionUnauthorized ", 401)
}

func (o *GetActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionForbidden creates a GetActionForbidden with default headers values
func NewGetActionForbidden() *GetActionForbidden {
	return &GetActionForbidden{}
}

/*GetActionForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetActionForbidden struct {
}

func (o *GetActionForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionForbidden ", 403)
}

func (o *GetActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionNotFound creates a GetActionNotFound with default headers values
func NewGetActionNotFound() *GetActionNotFound {
	return &GetActionNotFound{}
}

/*GetActionNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetActionNotFound struct {
}

func (o *GetActionNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionNotFound ", 404)
}

func (o *GetActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionConflict creates a GetActionConflict with default headers values
func NewGetActionConflict() *GetActionConflict {
	return &GetActionConflict{}
}

/*GetActionConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetActionConflict struct {
}

func (o *GetActionConflict) Error() string {
	return fmt.Sprintf("[GET /flow/history/{id}][%d] getActionConflict ", 409)
}

func (o *GetActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
