// Code generated by go-swagger; DO NOT EDIT.

package kafeido

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoListProjectPipelineReader is a Reader for the KafeidoListProjectPipeline structure.
type KafeidoListProjectPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoListProjectPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoListProjectPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoListProjectPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoListProjectPipelineOK creates a KafeidoListProjectPipelineOK with default headers values
func NewKafeidoListProjectPipelineOK() *KafeidoListProjectPipelineOK {
	return &KafeidoListProjectPipelineOK{}
}

/*
KafeidoListProjectPipelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoListProjectPipelineOK struct {
	Payload *models.KafeidoListProjectPipelineResponse
}

// IsSuccess returns true when this kafeido list project pipeline o k response has a 2xx status code
func (o *KafeidoListProjectPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido list project pipeline o k response has a 3xx status code
func (o *KafeidoListProjectPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido list project pipeline o k response has a 4xx status code
func (o *KafeidoListProjectPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido list project pipeline o k response has a 5xx status code
func (o *KafeidoListProjectPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido list project pipeline o k response a status code equal to that given
func (o *KafeidoListProjectPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido list project pipeline o k response
func (o *KafeidoListProjectPipelineOK) Code() int {
	return 200
}

func (o *KafeidoListProjectPipelineOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines][%d] kafeidoListProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoListProjectPipelineOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines][%d] kafeidoListProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoListProjectPipelineOK) GetPayload() *models.KafeidoListProjectPipelineResponse {
	return o.Payload
}

func (o *KafeidoListProjectPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoListProjectPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoListProjectPipelineDefault creates a KafeidoListProjectPipelineDefault with default headers values
func NewKafeidoListProjectPipelineDefault(code int) *KafeidoListProjectPipelineDefault {
	return &KafeidoListProjectPipelineDefault{
		_statusCode: code,
	}
}

/*
KafeidoListProjectPipelineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoListProjectPipelineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido list project pipeline default response has a 2xx status code
func (o *KafeidoListProjectPipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido list project pipeline default response has a 3xx status code
func (o *KafeidoListProjectPipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido list project pipeline default response has a 4xx status code
func (o *KafeidoListProjectPipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido list project pipeline default response has a 5xx status code
func (o *KafeidoListProjectPipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido list project pipeline default response a status code equal to that given
func (o *KafeidoListProjectPipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido list project pipeline default response
func (o *KafeidoListProjectPipelineDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoListProjectPipelineDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines][%d] kafeido_ListProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoListProjectPipelineDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines][%d] kafeido_ListProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoListProjectPipelineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoListProjectPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
