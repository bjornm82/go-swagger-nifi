// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjornm82/go-swagger-nifi/models"
)

// CreateEmptyAllConnectionsRequestReader is a Reader for the CreateEmptyAllConnectionsRequest structure.
type CreateEmptyAllConnectionsRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEmptyAllConnectionsRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEmptyAllConnectionsRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateEmptyAllConnectionsRequestAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEmptyAllConnectionsRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEmptyAllConnectionsRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEmptyAllConnectionsRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEmptyAllConnectionsRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEmptyAllConnectionsRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateEmptyAllConnectionsRequestOK creates a CreateEmptyAllConnectionsRequestOK with default headers values
func NewCreateEmptyAllConnectionsRequestOK() *CreateEmptyAllConnectionsRequestOK {
	return &CreateEmptyAllConnectionsRequestOK{}
}

/*CreateEmptyAllConnectionsRequestOK handles this case with default header values.

successful operation
*/
type CreateEmptyAllConnectionsRequestOK struct {
	Payload *models.ProcessGroupEntity
}

func (o *CreateEmptyAllConnectionsRequestOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestOK  %+v", 200, o.Payload)
}

func (o *CreateEmptyAllConnectionsRequestOK) GetPayload() *models.ProcessGroupEntity {
	return o.Payload
}

func (o *CreateEmptyAllConnectionsRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEmptyAllConnectionsRequestAccepted creates a CreateEmptyAllConnectionsRequestAccepted with default headers values
func NewCreateEmptyAllConnectionsRequestAccepted() *CreateEmptyAllConnectionsRequestAccepted {
	return &CreateEmptyAllConnectionsRequestAccepted{}
}

/*CreateEmptyAllConnectionsRequestAccepted handles this case with default header values.

The request has been accepted. An HTTP response header will contain the URI where the status can be polled.
*/
type CreateEmptyAllConnectionsRequestAccepted struct {
}

func (o *CreateEmptyAllConnectionsRequestAccepted) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestAccepted ", 202)
}

func (o *CreateEmptyAllConnectionsRequestAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmptyAllConnectionsRequestBadRequest creates a CreateEmptyAllConnectionsRequestBadRequest with default headers values
func NewCreateEmptyAllConnectionsRequestBadRequest() *CreateEmptyAllConnectionsRequestBadRequest {
	return &CreateEmptyAllConnectionsRequestBadRequest{}
}

/*CreateEmptyAllConnectionsRequestBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateEmptyAllConnectionsRequestBadRequest struct {
}

func (o *CreateEmptyAllConnectionsRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestBadRequest ", 400)
}

func (o *CreateEmptyAllConnectionsRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmptyAllConnectionsRequestUnauthorized creates a CreateEmptyAllConnectionsRequestUnauthorized with default headers values
func NewCreateEmptyAllConnectionsRequestUnauthorized() *CreateEmptyAllConnectionsRequestUnauthorized {
	return &CreateEmptyAllConnectionsRequestUnauthorized{}
}

/*CreateEmptyAllConnectionsRequestUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type CreateEmptyAllConnectionsRequestUnauthorized struct {
}

func (o *CreateEmptyAllConnectionsRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestUnauthorized ", 401)
}

func (o *CreateEmptyAllConnectionsRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmptyAllConnectionsRequestForbidden creates a CreateEmptyAllConnectionsRequestForbidden with default headers values
func NewCreateEmptyAllConnectionsRequestForbidden() *CreateEmptyAllConnectionsRequestForbidden {
	return &CreateEmptyAllConnectionsRequestForbidden{}
}

/*CreateEmptyAllConnectionsRequestForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type CreateEmptyAllConnectionsRequestForbidden struct {
}

func (o *CreateEmptyAllConnectionsRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestForbidden ", 403)
}

func (o *CreateEmptyAllConnectionsRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmptyAllConnectionsRequestNotFound creates a CreateEmptyAllConnectionsRequestNotFound with default headers values
func NewCreateEmptyAllConnectionsRequestNotFound() *CreateEmptyAllConnectionsRequestNotFound {
	return &CreateEmptyAllConnectionsRequestNotFound{}
}

/*CreateEmptyAllConnectionsRequestNotFound handles this case with default header values.

The specified resource could not be found.
*/
type CreateEmptyAllConnectionsRequestNotFound struct {
}

func (o *CreateEmptyAllConnectionsRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestNotFound ", 404)
}

func (o *CreateEmptyAllConnectionsRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEmptyAllConnectionsRequestConflict creates a CreateEmptyAllConnectionsRequestConflict with default headers values
func NewCreateEmptyAllConnectionsRequestConflict() *CreateEmptyAllConnectionsRequestConflict {
	return &CreateEmptyAllConnectionsRequestConflict{}
}

/*CreateEmptyAllConnectionsRequestConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateEmptyAllConnectionsRequestConflict struct {
}

func (o *CreateEmptyAllConnectionsRequestConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/empty-all-connections-requests][%d] createEmptyAllConnectionsRequestConflict ", 409)
}

func (o *CreateEmptyAllConnectionsRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
