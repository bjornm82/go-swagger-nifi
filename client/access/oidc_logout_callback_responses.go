// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OidcLogoutCallbackReader is a Reader for the OidcLogoutCallback structure.
type OidcLogoutCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcLogoutCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewOidcLogoutCallbackDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewOidcLogoutCallbackDefault creates a OidcLogoutCallbackDefault with default headers values
func NewOidcLogoutCallbackDefault(code int) *OidcLogoutCallbackDefault {
	return &OidcLogoutCallbackDefault{
		_statusCode: code,
	}
}

/*OidcLogoutCallbackDefault handles this case with default header values.

successful operation
*/
type OidcLogoutCallbackDefault struct {
	_statusCode int
}

// Code gets the status code for the oidc logout callback default response
func (o *OidcLogoutCallbackDefault) Code() int {
	return o._statusCode
}

func (o *OidcLogoutCallbackDefault) Error() string {
	return fmt.Sprintf("[GET /access/oidc/logoutCallback][%d] oidcLogoutCallback default ", o._statusCode)
}

func (o *OidcLogoutCallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}