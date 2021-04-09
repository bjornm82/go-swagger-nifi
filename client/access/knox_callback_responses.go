// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// KnoxCallbackReader is a Reader for the KnoxCallback structure.
type KnoxCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KnoxCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewKnoxCallbackDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewKnoxCallbackDefault creates a KnoxCallbackDefault with default headers values
func NewKnoxCallbackDefault(code int) *KnoxCallbackDefault {
	return &KnoxCallbackDefault{
		_statusCode: code,
	}
}

/*KnoxCallbackDefault handles this case with default header values.

successful operation
*/
type KnoxCallbackDefault struct {
	_statusCode int
}

// Code gets the status code for the knox callback default response
func (o *KnoxCallbackDefault) Code() int {
	return o._statusCode
}

func (o *KnoxCallbackDefault) Error() string {
	return fmt.Sprintf("[GET /access/knox/callback][%d] knoxCallback default ", o._statusCode)
}

func (o *KnoxCallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}