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

// KafeidoServiceGetModelInferenceJobReader is a Reader for the KafeidoServiceGetModelInferenceJob structure.
type KafeidoServiceGetModelInferenceJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceGetModelInferenceJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceGetModelInferenceJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceGetModelInferenceJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceGetModelInferenceJobOK creates a KafeidoServiceGetModelInferenceJobOK with default headers values
func NewKafeidoServiceGetModelInferenceJobOK() *KafeidoServiceGetModelInferenceJobOK {
	return &KafeidoServiceGetModelInferenceJobOK{}
}

/*
KafeidoServiceGetModelInferenceJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceGetModelInferenceJobOK struct {
	Payload *models.KafeidoGetModelInferenceJobResponse
}

// IsSuccess returns true when this kafeido service get model inference job o k response has a 2xx status code
func (o *KafeidoServiceGetModelInferenceJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service get model inference job o k response has a 3xx status code
func (o *KafeidoServiceGetModelInferenceJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service get model inference job o k response has a 4xx status code
func (o *KafeidoServiceGetModelInferenceJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service get model inference job o k response has a 5xx status code
func (o *KafeidoServiceGetModelInferenceJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service get model inference job o k response a status code equal to that given
func (o *KafeidoServiceGetModelInferenceJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service get model inference job o k response
func (o *KafeidoServiceGetModelInferenceJobOK) Code() int {
	return 200
}

func (o *KafeidoServiceGetModelInferenceJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] kafeidoServiceGetModelInferenceJobOK %s", 200, payload)
}

func (o *KafeidoServiceGetModelInferenceJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] kafeidoServiceGetModelInferenceJobOK %s", 200, payload)
}

func (o *KafeidoServiceGetModelInferenceJobOK) GetPayload() *models.KafeidoGetModelInferenceJobResponse {
	return o.Payload
}

func (o *KafeidoServiceGetModelInferenceJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoGetModelInferenceJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceGetModelInferenceJobDefault creates a KafeidoServiceGetModelInferenceJobDefault with default headers values
func NewKafeidoServiceGetModelInferenceJobDefault(code int) *KafeidoServiceGetModelInferenceJobDefault {
	return &KafeidoServiceGetModelInferenceJobDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceGetModelInferenceJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceGetModelInferenceJobDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service get model inference job default response has a 2xx status code
func (o *KafeidoServiceGetModelInferenceJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service get model inference job default response has a 3xx status code
func (o *KafeidoServiceGetModelInferenceJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service get model inference job default response has a 4xx status code
func (o *KafeidoServiceGetModelInferenceJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service get model inference job default response has a 5xx status code
func (o *KafeidoServiceGetModelInferenceJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service get model inference job default response a status code equal to that given
func (o *KafeidoServiceGetModelInferenceJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service get model inference job default response
func (o *KafeidoServiceGetModelInferenceJobDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceGetModelInferenceJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] KafeidoService_GetModelInferenceJob default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetModelInferenceJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] KafeidoService_GetModelInferenceJob default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetModelInferenceJobDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceGetModelInferenceJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
