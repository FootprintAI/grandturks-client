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

// KafeidoGetProjectPipelineReader is a Reader for the KafeidoGetProjectPipeline structure.
type KafeidoGetProjectPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoGetProjectPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoGetProjectPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoGetProjectPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoGetProjectPipelineOK creates a KafeidoGetProjectPipelineOK with default headers values
func NewKafeidoGetProjectPipelineOK() *KafeidoGetProjectPipelineOK {
	return &KafeidoGetProjectPipelineOK{}
}

/*
KafeidoGetProjectPipelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoGetProjectPipelineOK struct {
	Payload *models.KafeidoGetProjectPipelineResponse
}

// IsSuccess returns true when this kafeido get project pipeline o k response has a 2xx status code
func (o *KafeidoGetProjectPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido get project pipeline o k response has a 3xx status code
func (o *KafeidoGetProjectPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido get project pipeline o k response has a 4xx status code
func (o *KafeidoGetProjectPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido get project pipeline o k response has a 5xx status code
func (o *KafeidoGetProjectPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido get project pipeline o k response a status code equal to that given
func (o *KafeidoGetProjectPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido get project pipeline o k response
func (o *KafeidoGetProjectPipelineOK) Code() int {
	return 200
}

func (o *KafeidoGetProjectPipelineOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoGetProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoGetProjectPipelineOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoGetProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoGetProjectPipelineOK) GetPayload() *models.KafeidoGetProjectPipelineResponse {
	return o.Payload
}

func (o *KafeidoGetProjectPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoGetProjectPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoGetProjectPipelineDefault creates a KafeidoGetProjectPipelineDefault with default headers values
func NewKafeidoGetProjectPipelineDefault(code int) *KafeidoGetProjectPipelineDefault {
	return &KafeidoGetProjectPipelineDefault{
		_statusCode: code,
	}
}

/*
KafeidoGetProjectPipelineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoGetProjectPipelineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido get project pipeline default response has a 2xx status code
func (o *KafeidoGetProjectPipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido get project pipeline default response has a 3xx status code
func (o *KafeidoGetProjectPipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido get project pipeline default response has a 4xx status code
func (o *KafeidoGetProjectPipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido get project pipeline default response has a 5xx status code
func (o *KafeidoGetProjectPipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido get project pipeline default response a status code equal to that given
func (o *KafeidoGetProjectPipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido get project pipeline default response
func (o *KafeidoGetProjectPipelineDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoGetProjectPipelineDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeido_GetProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoGetProjectPipelineDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeido_GetProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoGetProjectPipelineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoGetProjectPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
