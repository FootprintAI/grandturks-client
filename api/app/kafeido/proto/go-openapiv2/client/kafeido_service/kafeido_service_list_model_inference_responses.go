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

// KafeidoServiceListModelInferenceReader is a Reader for the KafeidoServiceListModelInference structure.
type KafeidoServiceListModelInferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceListModelInferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceListModelInferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceListModelInferenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceListModelInferenceOK creates a KafeidoServiceListModelInferenceOK with default headers values
func NewKafeidoServiceListModelInferenceOK() *KafeidoServiceListModelInferenceOK {
	return &KafeidoServiceListModelInferenceOK{}
}

/*
KafeidoServiceListModelInferenceOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceListModelInferenceOK struct {
	Payload *models.KafeidoListModelInferenceResponse
}

// IsSuccess returns true when this kafeido service list model inference o k response has a 2xx status code
func (o *KafeidoServiceListModelInferenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service list model inference o k response has a 3xx status code
func (o *KafeidoServiceListModelInferenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service list model inference o k response has a 4xx status code
func (o *KafeidoServiceListModelInferenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service list model inference o k response has a 5xx status code
func (o *KafeidoServiceListModelInferenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service list model inference o k response a status code equal to that given
func (o *KafeidoServiceListModelInferenceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service list model inference o k response
func (o *KafeidoServiceListModelInferenceOK) Code() int {
	return 200
}

func (o *KafeidoServiceListModelInferenceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/inference][%d] kafeidoServiceListModelInferenceOK %s", 200, payload)
}

func (o *KafeidoServiceListModelInferenceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/inference][%d] kafeidoServiceListModelInferenceOK %s", 200, payload)
}

func (o *KafeidoServiceListModelInferenceOK) GetPayload() *models.KafeidoListModelInferenceResponse {
	return o.Payload
}

func (o *KafeidoServiceListModelInferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoListModelInferenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceListModelInferenceDefault creates a KafeidoServiceListModelInferenceDefault with default headers values
func NewKafeidoServiceListModelInferenceDefault(code int) *KafeidoServiceListModelInferenceDefault {
	return &KafeidoServiceListModelInferenceDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceListModelInferenceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceListModelInferenceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service list model inference default response has a 2xx status code
func (o *KafeidoServiceListModelInferenceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service list model inference default response has a 3xx status code
func (o *KafeidoServiceListModelInferenceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service list model inference default response has a 4xx status code
func (o *KafeidoServiceListModelInferenceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service list model inference default response has a 5xx status code
func (o *KafeidoServiceListModelInferenceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service list model inference default response a status code equal to that given
func (o *KafeidoServiceListModelInferenceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service list model inference default response
func (o *KafeidoServiceListModelInferenceDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceListModelInferenceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/inference][%d] KafeidoService_ListModelInference default %s", o._statusCode, payload)
}

func (o *KafeidoServiceListModelInferenceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/inference][%d] KafeidoService_ListModelInference default %s", o._statusCode, payload)
}

func (o *KafeidoServiceListModelInferenceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceListModelInferenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
