// Code generated by go-swagger; DO NOT EDIT.

package kafeido_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoServiceListProjectsReader is a Reader for the KafeidoServiceListProjects structure.
type KafeidoServiceListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceListProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceListProjectsOK creates a KafeidoServiceListProjectsOK with default headers values
func NewKafeidoServiceListProjectsOK() *KafeidoServiceListProjectsOK {
	return &KafeidoServiceListProjectsOK{}
}

/*
KafeidoServiceListProjectsOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceListProjectsOK struct {
	Payload *models.KafeidoListProjectsResponse
}

// IsSuccess returns true when this kafeido service list projects o k response has a 2xx status code
func (o *KafeidoServiceListProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service list projects o k response has a 3xx status code
func (o *KafeidoServiceListProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service list projects o k response has a 4xx status code
func (o *KafeidoServiceListProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service list projects o k response has a 5xx status code
func (o *KafeidoServiceListProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service list projects o k response a status code equal to that given
func (o *KafeidoServiceListProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service list projects o k response
func (o *KafeidoServiceListProjectsOK) Code() int {
	return 200
}

func (o *KafeidoServiceListProjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects][%d] kafeidoServiceListProjectsOK %s", 200, payload)
}

func (o *KafeidoServiceListProjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects][%d] kafeidoServiceListProjectsOK %s", 200, payload)
}

func (o *KafeidoServiceListProjectsOK) GetPayload() *models.KafeidoListProjectsResponse {
	return o.Payload
}

func (o *KafeidoServiceListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoListProjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceListProjectsDefault creates a KafeidoServiceListProjectsDefault with default headers values
func NewKafeidoServiceListProjectsDefault(code int) *KafeidoServiceListProjectsDefault {
	return &KafeidoServiceListProjectsDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceListProjectsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceListProjectsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service list projects default response has a 2xx status code
func (o *KafeidoServiceListProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service list projects default response has a 3xx status code
func (o *KafeidoServiceListProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service list projects default response has a 4xx status code
func (o *KafeidoServiceListProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service list projects default response has a 5xx status code
func (o *KafeidoServiceListProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service list projects default response a status code equal to that given
func (o *KafeidoServiceListProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service list projects default response
func (o *KafeidoServiceListProjectsDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceListProjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects][%d] KafeidoService_ListProjects default %s", o._statusCode, payload)
}

func (o *KafeidoServiceListProjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects][%d] KafeidoService_ListProjects default %s", o._statusCode, payload)
}

func (o *KafeidoServiceListProjectsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceListProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}