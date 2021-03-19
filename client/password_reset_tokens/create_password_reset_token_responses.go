// Code generated by go-swagger; DO NOT EDIT.

package password_reset_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePasswordResetTokenReader is a Reader for the CreatePasswordResetToken structure.
type CreatePasswordResetTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePasswordResetTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePasswordResetTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePasswordResetTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePasswordResetTokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePasswordResetTokenCreated creates a CreatePasswordResetTokenCreated with default headers values
func NewCreatePasswordResetTokenCreated() *CreatePasswordResetTokenCreated {
	return &CreatePasswordResetTokenCreated{}
}

/* CreatePasswordResetTokenCreated describes a response with status code 201, with default header values.

created
*/
type CreatePasswordResetTokenCreated struct {
}

func (o *CreatePasswordResetTokenCreated) Error() string {
	return fmt.Sprintf("[POST /reset-password][%d] createPasswordResetTokenCreated ", 201)
}

func (o *CreatePasswordResetTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePasswordResetTokenUnauthorized creates a CreatePasswordResetTokenUnauthorized with default headers values
func NewCreatePasswordResetTokenUnauthorized() *CreatePasswordResetTokenUnauthorized {
	return &CreatePasswordResetTokenUnauthorized{}
}

/* CreatePasswordResetTokenUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type CreatePasswordResetTokenUnauthorized struct {
}

func (o *CreatePasswordResetTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reset-password][%d] createPasswordResetTokenUnauthorized ", 401)
}

func (o *CreatePasswordResetTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePasswordResetTokenUnprocessableEntity creates a CreatePasswordResetTokenUnprocessableEntity with default headers values
func NewCreatePasswordResetTokenUnprocessableEntity() *CreatePasswordResetTokenUnprocessableEntity {
	return &CreatePasswordResetTokenUnprocessableEntity{}
}

/* CreatePasswordResetTokenUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type CreatePasswordResetTokenUnprocessableEntity struct {
}

func (o *CreatePasswordResetTokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /reset-password][%d] createPasswordResetTokenUnprocessableEntity ", 422)
}

func (o *CreatePasswordResetTokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}