// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// GetRevertRequestReader is a Reader for the GetRevertRequest structure.
type GetRevertRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRevertRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRevertRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRevertRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRevertRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRevertRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRevertRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetRevertRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRevertRequestOK creates a GetRevertRequestOK with default headers values
func NewGetRevertRequestOK() *GetRevertRequestOK {
	return &GetRevertRequestOK{}
}

/*GetRevertRequestOK handles this case with default header values.

successful operation
*/
type GetRevertRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

func (o *GetRevertRequestOK) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestOK  %+v", 200, o.Payload)
}

func (o *GetRevertRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevertRequestBadRequest creates a GetRevertRequestBadRequest with default headers values
func NewGetRevertRequestBadRequest() *GetRevertRequestBadRequest {
	return &GetRevertRequestBadRequest{}
}

/*GetRevertRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRevertRequestBadRequest struct {
}

func (o *GetRevertRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestBadRequest ", 400)
}

func (o *GetRevertRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestUnauthorized creates a GetRevertRequestUnauthorized with default headers values
func NewGetRevertRequestUnauthorized() *GetRevertRequestUnauthorized {
	return &GetRevertRequestUnauthorized{}
}

/*GetRevertRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetRevertRequestUnauthorized struct {
}

func (o *GetRevertRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestUnauthorized ", 401)
}

func (o *GetRevertRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestForbidden creates a GetRevertRequestForbidden with default headers values
func NewGetRevertRequestForbidden() *GetRevertRequestForbidden {
	return &GetRevertRequestForbidden{}
}

/*GetRevertRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetRevertRequestForbidden struct {
}

func (o *GetRevertRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestForbidden ", 403)
}

func (o *GetRevertRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestNotFound creates a GetRevertRequestNotFound with default headers values
func NewGetRevertRequestNotFound() *GetRevertRequestNotFound {
	return &GetRevertRequestNotFound{}
}

/*GetRevertRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetRevertRequestNotFound struct {
}

func (o *GetRevertRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestNotFound ", 404)
}

func (o *GetRevertRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestConflict creates a GetRevertRequestConflict with default headers values
func NewGetRevertRequestConflict() *GetRevertRequestConflict {
	return &GetRevertRequestConflict{}
}

/*GetRevertRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRevertRequestConflict struct {
}

func (o *GetRevertRequestConflict) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestConflict ", 409)
}

func (o *GetRevertRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}