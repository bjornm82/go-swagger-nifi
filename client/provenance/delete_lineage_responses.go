// Code generated by go-swagger; DO NOT EDIT.

package provenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// DeleteLineageReader is a Reader for the DeleteLineage structure.
type DeleteLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLineageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLineageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLineageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLineageOK creates a DeleteLineageOK with default headers values
func NewDeleteLineageOK() *DeleteLineageOK {
	return &DeleteLineageOK{}
}

/*DeleteLineageOK handles this case with default header values.

successful operation
*/
type DeleteLineageOK struct {
	Payload *models.LineageEntity
}

func (o *DeleteLineageOK) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageOK  %+v", 200, o.Payload)
}

func (o *DeleteLineageOK) GetPayload() *models.LineageEntity {
	return o.Payload
}

func (o *DeleteLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineageEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLineageBadRequest creates a DeleteLineageBadRequest with default headers values
func NewDeleteLineageBadRequest() *DeleteLineageBadRequest {
	return &DeleteLineageBadRequest{}
}

/*DeleteLineageBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteLineageBadRequest struct {
}

func (o *DeleteLineageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageBadRequest ", 400)
}

func (o *DeleteLineageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLineageUnauthorized creates a DeleteLineageUnauthorized with default headers values
func NewDeleteLineageUnauthorized() *DeleteLineageUnauthorized {
	return &DeleteLineageUnauthorized{}
}

/*DeleteLineageUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type DeleteLineageUnauthorized struct {
}

func (o *DeleteLineageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageUnauthorized ", 401)
}

func (o *DeleteLineageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLineageForbidden creates a DeleteLineageForbidden with default headers values
func NewDeleteLineageForbidden() *DeleteLineageForbidden {
	return &DeleteLineageForbidden{}
}

/*DeleteLineageForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type DeleteLineageForbidden struct {
}

func (o *DeleteLineageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageForbidden ", 403)
}

func (o *DeleteLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLineageNotFound creates a DeleteLineageNotFound with default headers values
func NewDeleteLineageNotFound() *DeleteLineageNotFound {
	return &DeleteLineageNotFound{}
}

/*DeleteLineageNotFound handles this case with default header values.

The specified resource could not be found.
*/
type DeleteLineageNotFound struct {
}

func (o *DeleteLineageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageNotFound ", 404)
}

func (o *DeleteLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLineageConflict creates a DeleteLineageConflict with default headers values
func NewDeleteLineageConflict() *DeleteLineageConflict {
	return &DeleteLineageConflict{}
}

/*DeleteLineageConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteLineageConflict struct {
}

func (o *DeleteLineageConflict) Error() string {
	return fmt.Sprintf("[DELETE /provenance/lineage/{id}][%d] deleteLineageConflict ", 409)
}

func (o *DeleteLineageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
