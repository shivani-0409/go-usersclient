// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/shivani-0409/go-usersclient/models"
)

// FindUsersReader is a Reader for the FindUsers structure.
type FindUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindUsersOK creates a FindUsersOK with default headers values
func NewFindUsersOK() *FindUsersOK {
	return &FindUsersOK{}
}

/* FindUsersOK describes a response with status code 200, with default header values.

list the User operations
*/
type FindUsersOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this find users o k response has a 2xx status code
func (o *FindUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find users o k response has a 3xx status code
func (o *FindUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find users o k response has a 4xx status code
func (o *FindUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find users o k response has a 5xx status code
func (o *FindUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find users o k response a status code equal to that given
func (o *FindUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] findUsersOK  %+v", 200, o.Payload)
}

func (o *FindUsersOK) String() string {
	return fmt.Sprintf("[GET /users][%d] findUsersOK  %+v", 200, o.Payload)
}

func (o *FindUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *FindUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersDefault creates a FindUsersDefault with default headers values
func NewFindUsersDefault(code int) *FindUsersDefault {
	return &FindUsersDefault{
		_statusCode: code,
	}
}

/* FindUsersDefault describes a response with status code -1, with default header values.

generic error response
*/
type FindUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find users default response
func (o *FindUsersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find users default response has a 2xx status code
func (o *FindUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find users default response has a 3xx status code
func (o *FindUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find users default response has a 4xx status code
func (o *FindUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find users default response has a 5xx status code
func (o *FindUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find users default response a status code equal to that given
func (o *FindUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] find_Users default  %+v", o._statusCode, o.Payload)
}

func (o *FindUsersDefault) String() string {
	return fmt.Sprintf("[GET /users][%d] find_Users default  %+v", o._statusCode, o.Payload)
}

func (o *FindUsersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
