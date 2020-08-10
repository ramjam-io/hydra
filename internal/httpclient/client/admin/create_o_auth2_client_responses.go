// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// CreateOAuth2ClientReader is a Reader for the CreateOAuth2Client structure.
type CreateOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOAuth2ClientCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOAuth2ClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateOAuth2ClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOAuth2ClientCreated creates a CreateOAuth2ClientCreated with default headers values
func NewCreateOAuth2ClientCreated() *CreateOAuth2ClientCreated {
	return &CreateOAuth2ClientCreated{}
}

/*CreateOAuth2ClientCreated handles this case with default header values.

oAuth2Client
*/
type CreateOAuth2ClientCreated struct {
	Payload *models.OAuth2Client
}

func (o *CreateOAuth2ClientCreated) Error() string {
	return fmt.Sprintf("[POST /clients][%d] createOAuth2ClientCreated  %+v", 201, o.Payload)
}

func (o *CreateOAuth2ClientCreated) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *CreateOAuth2ClientCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOAuth2ClientBadRequest creates a CreateOAuth2ClientBadRequest with default headers values
func NewCreateOAuth2ClientBadRequest() *CreateOAuth2ClientBadRequest {
	return &CreateOAuth2ClientBadRequest{}
}

/*CreateOAuth2ClientBadRequest handles this case with default header values.

genericError
*/
type CreateOAuth2ClientBadRequest struct {
	Payload *models.GenericError
}

func (o *CreateOAuth2ClientBadRequest) Error() string {
	return fmt.Sprintf("[POST /clients][%d] createOAuth2ClientBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOAuth2ClientBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateOAuth2ClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOAuth2ClientConflict creates a CreateOAuth2ClientConflict with default headers values
func NewCreateOAuth2ClientConflict() *CreateOAuth2ClientConflict {
	return &CreateOAuth2ClientConflict{}
}

/*CreateOAuth2ClientConflict handles this case with default header values.

genericError
*/
type CreateOAuth2ClientConflict struct {
	Payload *models.GenericError
}

func (o *CreateOAuth2ClientConflict) Error() string {
	return fmt.Sprintf("[POST /clients][%d] createOAuth2ClientConflict  %+v", 409, o.Payload)
}

func (o *CreateOAuth2ClientConflict) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateOAuth2ClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOAuth2ClientInternalServerError creates a CreateOAuth2ClientInternalServerError with default headers values
func NewCreateOAuth2ClientInternalServerError() *CreateOAuth2ClientInternalServerError {
	return &CreateOAuth2ClientInternalServerError{}
}

/*CreateOAuth2ClientInternalServerError handles this case with default header values.

genericError
*/
type CreateOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *CreateOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clients][%d] createOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOAuth2ClientInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
