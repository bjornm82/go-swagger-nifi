// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// KnoxRequestReader is a Reader for the KnoxRequest structure.
type KnoxRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KnoxRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewKnoxRequestDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewKnoxRequestDefault creates a KnoxRequestDefault with default headers values
func NewKnoxRequestDefault(code int) *KnoxRequestDefault {
	return &KnoxRequestDefault{
		_statusCode: code,
	}
}

/*KnoxRequestDefault handles this case with default header values.

successful operation
*/
type KnoxRequestDefault struct {
	_statusCode int
}

// Code gets the status code for the knox request default response
func (o *KnoxRequestDefault) Code() int {
	return o._statusCode
}

func (o *KnoxRequestDefault) Error() string {
	return fmt.Sprintf("[GET /access/knox/request][%d] knoxRequest default ", o._statusCode)
}

func (o *KnoxRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
