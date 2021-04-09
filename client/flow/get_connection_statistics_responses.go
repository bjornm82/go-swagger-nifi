// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-nifi/models"
)

// GetConnectionStatisticsReader is a Reader for the GetConnectionStatistics structure.
type GetConnectionStatisticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionStatisticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConnectionStatisticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetConnectionStatisticsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetConnectionStatisticsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetConnectionStatisticsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetConnectionStatisticsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewGetConnectionStatisticsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConnectionStatisticsOK creates a GetConnectionStatisticsOK with default headers values
func NewGetConnectionStatisticsOK() *GetConnectionStatisticsOK {
	return &GetConnectionStatisticsOK{}
}

/*GetConnectionStatisticsOK handles this case with default header values.

successful operation
*/
type GetConnectionStatisticsOK struct {
	Payload *models.ConnectionStatisticsEntity
}

func (o *GetConnectionStatisticsOK) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsOK  %+v", 200, o.Payload)
}

func (o *GetConnectionStatisticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionStatisticsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionStatisticsBadRequest creates a GetConnectionStatisticsBadRequest with default headers values
func NewGetConnectionStatisticsBadRequest() *GetConnectionStatisticsBadRequest {
	return &GetConnectionStatisticsBadRequest{}
}

/*GetConnectionStatisticsBadRequest handles this case with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetConnectionStatisticsBadRequest struct {
}

func (o *GetConnectionStatisticsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsBadRequest ", 400)
}

func (o *GetConnectionStatisticsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatisticsUnauthorized creates a GetConnectionStatisticsUnauthorized with default headers values
func NewGetConnectionStatisticsUnauthorized() *GetConnectionStatisticsUnauthorized {
	return &GetConnectionStatisticsUnauthorized{}
}

/*GetConnectionStatisticsUnauthorized handles this case with default header values.

Client could not be authenticated.
*/
type GetConnectionStatisticsUnauthorized struct {
}

func (o *GetConnectionStatisticsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsUnauthorized ", 401)
}

func (o *GetConnectionStatisticsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatisticsForbidden creates a GetConnectionStatisticsForbidden with default headers values
func NewGetConnectionStatisticsForbidden() *GetConnectionStatisticsForbidden {
	return &GetConnectionStatisticsForbidden{}
}

/*GetConnectionStatisticsForbidden handles this case with default header values.

Client is not authorized to make this request.
*/
type GetConnectionStatisticsForbidden struct {
}

func (o *GetConnectionStatisticsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsForbidden ", 403)
}

func (o *GetConnectionStatisticsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatisticsNotFound creates a GetConnectionStatisticsNotFound with default headers values
func NewGetConnectionStatisticsNotFound() *GetConnectionStatisticsNotFound {
	return &GetConnectionStatisticsNotFound{}
}

/*GetConnectionStatisticsNotFound handles this case with default header values.

The specified resource could not be found.
*/
type GetConnectionStatisticsNotFound struct {
}

func (o *GetConnectionStatisticsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsNotFound ", 404)
}

func (o *GetConnectionStatisticsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionStatisticsConflict creates a GetConnectionStatisticsConflict with default headers values
func NewGetConnectionStatisticsConflict() *GetConnectionStatisticsConflict {
	return &GetConnectionStatisticsConflict{}
}

/*GetConnectionStatisticsConflict handles this case with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetConnectionStatisticsConflict struct {
}

func (o *GetConnectionStatisticsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/connections/{id}/statistics][%d] getConnectionStatisticsConflict ", 409)
}

func (o *GetConnectionStatisticsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
