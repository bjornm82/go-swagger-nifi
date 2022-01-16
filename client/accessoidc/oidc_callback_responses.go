// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OidcCallbackReader is a Reader for the OidcCallback structure.
type OidcCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewOidcCallbackDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewOidcCallbackDefault creates a OidcCallbackDefault with default headers values
func NewOidcCallbackDefault(code int) *OidcCallbackDefault {
	return &OidcCallbackDefault{
		_statusCode: code,
	}
}

/*OidcCallbackDefault handles this case with default header values.

successful operation
*/
type OidcCallbackDefault struct {
	_statusCode int
}

// Code gets the status code for the oidc callback default response
func (o *OidcCallbackDefault) Code() int {
	return o._statusCode
}

func (o *OidcCallbackDefault) Error() string {
	return fmt.Sprintf("[GET /access/oidc/callback][%d] oidcCallback default ", o._statusCode)
}

func (o *OidcCallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
