// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/swagger-nifi/models"
)

// GetVersionsReader is a Reader for the GetVersions structure.
type GetVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVersionsOK creates a GetVersionsOK with default headers values
func NewGetVersionsOK() *GetVersionsOK {
	return &GetVersionsOK{}
}

/*GetVersionsOK handles this case with default header values.

successful operation
*/
type GetVersionsOK struct {
	Payload *models.VersionedFlowSnapshotMetadataSetEntity
}

func (o *GetVersionsOK) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsOK  %+v", 200, o.Payload)
}

func (o *GetVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshotMetadataSetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionsBadRequest creates a GetVersionsBadRequest with default headers values
func NewGetVersionsBadRequest() *GetVersionsBadRequest {
	return &GetVersionsBadRequest{}
}

/*GetVersionsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetVersionsBadRequest struct {
}

func (o *GetVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsBadRequest ", 400)
}

func (o *GetVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsUnauthorized creates a GetVersionsUnauthorized with default headers values
func NewGetVersionsUnauthorized() *GetVersionsUnauthorized {
	return &GetVersionsUnauthorized{}
}

/*GetVersionsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetVersionsUnauthorized struct {
}

func (o *GetVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsUnauthorized ", 401)
}

func (o *GetVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsForbidden creates a GetVersionsForbidden with default headers values
func NewGetVersionsForbidden() *GetVersionsForbidden {
	return &GetVersionsForbidden{}
}

/*GetVersionsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetVersionsForbidden struct {
}

func (o *GetVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsForbidden ", 403)
}

func (o *GetVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsNotFound creates a GetVersionsNotFound with default headers values
func NewGetVersionsNotFound() *GetVersionsNotFound {
	return &GetVersionsNotFound{}
}

/*GetVersionsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetVersionsNotFound struct {
}

func (o *GetVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsNotFound ", 404)
}

func (o *GetVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsConflict creates a GetVersionsConflict with default headers values
func NewGetVersionsConflict() *GetVersionsConflict {
	return &GetVersionsConflict{}
}

/*GetVersionsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetVersionsConflict struct {
}

func (o *GetVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsConflict ", 409)
}

func (o *GetVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
