// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paralus/paralus/api/def/clients/sentry/models"
)

// BootstrapGetBootstrapAgentsReader is a Reader for the BootstrapGetBootstrapAgents structure.
type BootstrapGetBootstrapAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BootstrapGetBootstrapAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBootstrapGetBootstrapAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBootstrapGetBootstrapAgentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBootstrapGetBootstrapAgentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBootstrapGetBootstrapAgentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBootstrapGetBootstrapAgentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBootstrapGetBootstrapAgentsOK creates a BootstrapGetBootstrapAgentsOK with default headers values
func NewBootstrapGetBootstrapAgentsOK() *BootstrapGetBootstrapAgentsOK {
	return &BootstrapGetBootstrapAgentsOK{}
}

/*
	BootstrapGetBootstrapAgentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type BootstrapGetBootstrapAgentsOK struct {
	Payload *models.SentryBootstrapAgentList
}

func (o *BootstrapGetBootstrapAgentsOK) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{templateScope}/agent][%d] bootstrapGetBootstrapAgentsOK  %+v", 200, o.Payload)
}
func (o *BootstrapGetBootstrapAgentsOK) GetPayload() *models.SentryBootstrapAgentList {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentryBootstrapAgentList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentsForbidden creates a BootstrapGetBootstrapAgentsForbidden with default headers values
func NewBootstrapGetBootstrapAgentsForbidden() *BootstrapGetBootstrapAgentsForbidden {
	return &BootstrapGetBootstrapAgentsForbidden{}
}

/*
	BootstrapGetBootstrapAgentsForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type BootstrapGetBootstrapAgentsForbidden struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{templateScope}/agent][%d] bootstrapGetBootstrapAgentsForbidden  %+v", 403, o.Payload)
}
func (o *BootstrapGetBootstrapAgentsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentsNotFound creates a BootstrapGetBootstrapAgentsNotFound with default headers values
func NewBootstrapGetBootstrapAgentsNotFound() *BootstrapGetBootstrapAgentsNotFound {
	return &BootstrapGetBootstrapAgentsNotFound{}
}

/*
	BootstrapGetBootstrapAgentsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type BootstrapGetBootstrapAgentsNotFound struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{templateScope}/agent][%d] bootstrapGetBootstrapAgentsNotFound  %+v", 404, o.Payload)
}
func (o *BootstrapGetBootstrapAgentsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentsInternalServerError creates a BootstrapGetBootstrapAgentsInternalServerError with default headers values
func NewBootstrapGetBootstrapAgentsInternalServerError() *BootstrapGetBootstrapAgentsInternalServerError {
	return &BootstrapGetBootstrapAgentsInternalServerError{}
}

/*
	BootstrapGetBootstrapAgentsInternalServerError describes a response with status code 500, with default header values.

Returned for internal server error
*/
type BootstrapGetBootstrapAgentsInternalServerError struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{templateScope}/agent][%d] bootstrapGetBootstrapAgentsInternalServerError  %+v", 500, o.Payload)
}
func (o *BootstrapGetBootstrapAgentsInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentsDefault creates a BootstrapGetBootstrapAgentsDefault with default headers values
func NewBootstrapGetBootstrapAgentsDefault(code int) *BootstrapGetBootstrapAgentsDefault {
	return &BootstrapGetBootstrapAgentsDefault{
		_statusCode: code,
	}
}

/*
	BootstrapGetBootstrapAgentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BootstrapGetBootstrapAgentsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the bootstrap get bootstrap agents default response
func (o *BootstrapGetBootstrapAgentsDefault) Code() int {
	return o._statusCode
}

func (o *BootstrapGetBootstrapAgentsDefault) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/{templateScope}/agent][%d] Bootstrap_GetBootstrapAgents default  %+v", o._statusCode, o.Payload)
}
func (o *BootstrapGetBootstrapAgentsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
