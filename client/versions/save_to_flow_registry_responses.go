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

// SaveToFlowRegistryReader is a Reader for the SaveToFlowRegistry structure.
type SaveToFlowRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveToFlowRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSaveToFlowRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSaveToFlowRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSaveToFlowRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSaveToFlowRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSaveToFlowRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewSaveToFlowRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSaveToFlowRegistryOK creates a SaveToFlowRegistryOK with default headers values
func NewSaveToFlowRegistryOK() *SaveToFlowRegistryOK {
	return &SaveToFlowRegistryOK{}
}

/*SaveToFlowRegistryOK handles this case with default header values.

successful operation
*/
type SaveToFlowRegistryOK struct {
	Payload *models.VersionControlInformationEntity
}

func (o *SaveToFlowRegistryOK) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryOK  %+v", 200, o.Payload)
}

func (o *SaveToFlowRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionControlInformationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveToFlowRegistryBadRequest creates a SaveToFlowRegistryBadRequest with default headers values
func NewSaveToFlowRegistryBadRequest() *SaveToFlowRegistryBadRequest {
	return &SaveToFlowRegistryBadRequest{}
}

/*SaveToFlowRegistryBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SaveToFlowRegistryBadRequest struct {
}

func (o *SaveToFlowRegistryBadRequest) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryBadRequest ", 400)
}

func (o *SaveToFlowRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryUnauthorized creates a SaveToFlowRegistryUnauthorized with default headers values
func NewSaveToFlowRegistryUnauthorized() *SaveToFlowRegistryUnauthorized {
	return &SaveToFlowRegistryUnauthorized{}
}

/*SaveToFlowRegistryUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type SaveToFlowRegistryUnauthorized struct {
}

func (o *SaveToFlowRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryUnauthorized ", 401)
}

func (o *SaveToFlowRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryForbidden creates a SaveToFlowRegistryForbidden with default headers values
func NewSaveToFlowRegistryForbidden() *SaveToFlowRegistryForbidden {
	return &SaveToFlowRegistryForbidden{}
}

/*SaveToFlowRegistryForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type SaveToFlowRegistryForbidden struct {
}

func (o *SaveToFlowRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryForbidden ", 403)
}

func (o *SaveToFlowRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryNotFound creates a SaveToFlowRegistryNotFound with default headers values
func NewSaveToFlowRegistryNotFound() *SaveToFlowRegistryNotFound {
	return &SaveToFlowRegistryNotFound{}
}

/*SaveToFlowRegistryNotFound handles this case with default header values.

The specified resource could not be found.
*/
type SaveToFlowRegistryNotFound struct {
}

func (o *SaveToFlowRegistryNotFound) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryNotFound ", 404)
}

func (o *SaveToFlowRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryConflict creates a SaveToFlowRegistryConflict with default headers values
func NewSaveToFlowRegistryConflict() *SaveToFlowRegistryConflict {
	return &SaveToFlowRegistryConflict{}
}

/*SaveToFlowRegistryConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SaveToFlowRegistryConflict struct {
}

func (o *SaveToFlowRegistryConflict) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryConflict ", 409)
}

func (o *SaveToFlowRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
