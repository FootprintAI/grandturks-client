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

// KafeidoServiceDeleteModelInferenceJobReader is a Reader for the KafeidoServiceDeleteModelInferenceJob structure.
type KafeidoServiceDeleteModelInferenceJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceDeleteModelInferenceJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceDeleteModelInferenceJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceDeleteModelInferenceJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceDeleteModelInferenceJobOK creates a KafeidoServiceDeleteModelInferenceJobOK with default headers values
func NewKafeidoServiceDeleteModelInferenceJobOK() *KafeidoServiceDeleteModelInferenceJobOK {
	return &KafeidoServiceDeleteModelInferenceJobOK{}
}

/*
KafeidoServiceDeleteModelInferenceJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceDeleteModelInferenceJobOK struct {
	Payload models.KafeidoDeleteModelInferenceJobResponse
}

// IsSuccess returns true when this kafeido service delete model inference job o k response has a 2xx status code
func (o *KafeidoServiceDeleteModelInferenceJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service delete model inference job o k response has a 3xx status code
func (o *KafeidoServiceDeleteModelInferenceJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service delete model inference job o k response has a 4xx status code
func (o *KafeidoServiceDeleteModelInferenceJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service delete model inference job o k response has a 5xx status code
func (o *KafeidoServiceDeleteModelInferenceJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service delete model inference job o k response a status code equal to that given
func (o *KafeidoServiceDeleteModelInferenceJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service delete model inference job o k response
func (o *KafeidoServiceDeleteModelInferenceJobOK) Code() int {
	return 200
}

func (o *KafeidoServiceDeleteModelInferenceJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] kafeidoServiceDeleteModelInferenceJobOK %s", 200, payload)
}

func (o *KafeidoServiceDeleteModelInferenceJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] kafeidoServiceDeleteModelInferenceJobOK %s", 200, payload)
}

func (o *KafeidoServiceDeleteModelInferenceJobOK) GetPayload() models.KafeidoDeleteModelInferenceJobResponse {
	return o.Payload
}

func (o *KafeidoServiceDeleteModelInferenceJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceDeleteModelInferenceJobDefault creates a KafeidoServiceDeleteModelInferenceJobDefault with default headers values
func NewKafeidoServiceDeleteModelInferenceJobDefault(code int) *KafeidoServiceDeleteModelInferenceJobDefault {
	return &KafeidoServiceDeleteModelInferenceJobDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceDeleteModelInferenceJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceDeleteModelInferenceJobDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service delete model inference job default response has a 2xx status code
func (o *KafeidoServiceDeleteModelInferenceJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service delete model inference job default response has a 3xx status code
func (o *KafeidoServiceDeleteModelInferenceJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service delete model inference job default response has a 4xx status code
func (o *KafeidoServiceDeleteModelInferenceJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service delete model inference job default response has a 5xx status code
func (o *KafeidoServiceDeleteModelInferenceJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service delete model inference job default response a status code equal to that given
func (o *KafeidoServiceDeleteModelInferenceJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service delete model inference job default response
func (o *KafeidoServiceDeleteModelInferenceJobDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceDeleteModelInferenceJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] KafeidoService_DeleteModelInferenceJob default %s", o._statusCode, payload)
}

func (o *KafeidoServiceDeleteModelInferenceJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/job/{modelInferenceJobId}][%d] KafeidoService_DeleteModelInferenceJob default %s", o._statusCode, payload)
}

func (o *KafeidoServiceDeleteModelInferenceJobDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceDeleteModelInferenceJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
