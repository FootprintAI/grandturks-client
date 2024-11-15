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

// KafeidoServiceGetUserReader is a Reader for the KafeidoServiceGetUser structure.
type KafeidoServiceGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceGetUserOK creates a KafeidoServiceGetUserOK with default headers values
func NewKafeidoServiceGetUserOK() *KafeidoServiceGetUserOK {
	return &KafeidoServiceGetUserOK{}
}

/*
KafeidoServiceGetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceGetUserOK struct {
	Payload *models.KafeidoGetUserResponse
}

// IsSuccess returns true when this kafeido service get user o k response has a 2xx status code
func (o *KafeidoServiceGetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service get user o k response has a 3xx status code
func (o *KafeidoServiceGetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service get user o k response has a 4xx status code
func (o *KafeidoServiceGetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service get user o k response has a 5xx status code
func (o *KafeidoServiceGetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service get user o k response a status code equal to that given
func (o *KafeidoServiceGetUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service get user o k response
func (o *KafeidoServiceGetUserOK) Code() int {
	return 200
}

func (o *KafeidoServiceGetUserOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/user/{userId}][%d] kafeidoServiceGetUserOK %s", 200, payload)
}

func (o *KafeidoServiceGetUserOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/user/{userId}][%d] kafeidoServiceGetUserOK %s", 200, payload)
}

func (o *KafeidoServiceGetUserOK) GetPayload() *models.KafeidoGetUserResponse {
	return o.Payload
}

func (o *KafeidoServiceGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoGetUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceGetUserDefault creates a KafeidoServiceGetUserDefault with default headers values
func NewKafeidoServiceGetUserDefault(code int) *KafeidoServiceGetUserDefault {
	return &KafeidoServiceGetUserDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceGetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceGetUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service get user default response has a 2xx status code
func (o *KafeidoServiceGetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service get user default response has a 3xx status code
func (o *KafeidoServiceGetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service get user default response has a 4xx status code
func (o *KafeidoServiceGetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service get user default response has a 5xx status code
func (o *KafeidoServiceGetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service get user default response a status code equal to that given
func (o *KafeidoServiceGetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service get user default response
func (o *KafeidoServiceGetUserDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceGetUserDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/user/{userId}][%d] KafeidoService_GetUser default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetUserDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/user/{userId}][%d] KafeidoService_GetUser default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
