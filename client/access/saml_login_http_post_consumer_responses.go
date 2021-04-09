// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SamlLoginHTTPPostConsumerReader is a Reader for the SamlLoginHTTPPostConsumer structure.
type SamlLoginHTTPPostConsumerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlLoginHTTPPostConsumerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewSamlLoginHTTPPostConsumerDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewSamlLoginHTTPPostConsumerDefault creates a SamlLoginHTTPPostConsumerDefault with default headers values
func NewSamlLoginHTTPPostConsumerDefault(code int) *SamlLoginHTTPPostConsumerDefault {
	return &SamlLoginHTTPPostConsumerDefault{
		_statusCode: code,
	}
}

/*SamlLoginHTTPPostConsumerDefault handles this case with default header values.

successful operation
*/
type SamlLoginHTTPPostConsumerDefault struct {
	_statusCode int
}

// Code gets the status code for the saml login Http post consumer default response
func (o *SamlLoginHTTPPostConsumerDefault) Code() int {
	return o._statusCode
}

func (o *SamlLoginHTTPPostConsumerDefault) Error() string {
	return fmt.Sprintf("[POST /access/saml/login/consumer][%d] samlLoginHttpPostConsumer default ", o._statusCode)
}

func (o *SamlLoginHTTPPostConsumerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
