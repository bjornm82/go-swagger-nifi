// Code generated by go-swagger; DO NOT EDIT.

package provenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// SubmitLineageRequestReader is a Reader for the SubmitLineageRequest structure.
type SubmitLineageRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitLineageRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubmitLineageRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSubmitLineageRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSubmitLineageRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSubmitLineageRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSubmitLineageRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewSubmitLineageRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitLineageRequestOK creates a SubmitLineageRequestOK with default headers values
func NewSubmitLineageRequestOK() *SubmitLineageRequestOK {
	return &SubmitLineageRequestOK{}
}

/*SubmitLineageRequestOK handles this case with default header values.

successful operation
*/
type SubmitLineageRequestOK struct {
	Payload *models.LineageEntity
}

func (o *SubmitLineageRequestOK) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestOK  %+v", 200, o.Payload)
}

func (o *SubmitLineageRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineageEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitLineageRequestBadRequest creates a SubmitLineageRequestBadRequest with default headers values
func NewSubmitLineageRequestBadRequest() *SubmitLineageRequestBadRequest {
	return &SubmitLineageRequestBadRequest{}
}

/*SubmitLineageRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SubmitLineageRequestBadRequest struct {
}

func (o *SubmitLineageRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestBadRequest ", 400)
}

func (o *SubmitLineageRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitLineageRequestUnauthorized creates a SubmitLineageRequestUnauthorized with default headers values
func NewSubmitLineageRequestUnauthorized() *SubmitLineageRequestUnauthorized {
	return &SubmitLineageRequestUnauthorized{}
}

/*SubmitLineageRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type SubmitLineageRequestUnauthorized struct {
}

func (o *SubmitLineageRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestUnauthorized ", 401)
}

func (o *SubmitLineageRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitLineageRequestForbidden creates a SubmitLineageRequestForbidden with default headers values
func NewSubmitLineageRequestForbidden() *SubmitLineageRequestForbidden {
	return &SubmitLineageRequestForbidden{}
}

/*SubmitLineageRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type SubmitLineageRequestForbidden struct {
}

func (o *SubmitLineageRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestForbidden ", 403)
}

func (o *SubmitLineageRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitLineageRequestNotFound creates a SubmitLineageRequestNotFound with default headers values
func NewSubmitLineageRequestNotFound() *SubmitLineageRequestNotFound {
	return &SubmitLineageRequestNotFound{}
}

/*SubmitLineageRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type SubmitLineageRequestNotFound struct {
}

func (o *SubmitLineageRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestNotFound ", 404)
}

func (o *SubmitLineageRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitLineageRequestConflict creates a SubmitLineageRequestConflict with default headers values
func NewSubmitLineageRequestConflict() *SubmitLineageRequestConflict {
	return &SubmitLineageRequestConflict{}
}

/*SubmitLineageRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SubmitLineageRequestConflict struct {
}

func (o *SubmitLineageRequestConflict) Error() string {
	return fmt.Sprintf("[POST /provenance/lineage][%d] submitLineageRequestConflict ", 409)
}

func (o *SubmitLineageRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
