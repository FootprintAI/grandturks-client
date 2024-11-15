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

// KafeidoServiceDeleteProjectPipelineReader is a Reader for the KafeidoServiceDeleteProjectPipeline structure.
type KafeidoServiceDeleteProjectPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceDeleteProjectPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceDeleteProjectPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceDeleteProjectPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceDeleteProjectPipelineOK creates a KafeidoServiceDeleteProjectPipelineOK with default headers values
func NewKafeidoServiceDeleteProjectPipelineOK() *KafeidoServiceDeleteProjectPipelineOK {
	return &KafeidoServiceDeleteProjectPipelineOK{}
}

/*
KafeidoServiceDeleteProjectPipelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceDeleteProjectPipelineOK struct {
	Payload models.KafeidoDeleteProjectPipelineResponse
}

// IsSuccess returns true when this kafeido service delete project pipeline o k response has a 2xx status code
func (o *KafeidoServiceDeleteProjectPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service delete project pipeline o k response has a 3xx status code
func (o *KafeidoServiceDeleteProjectPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service delete project pipeline o k response has a 4xx status code
func (o *KafeidoServiceDeleteProjectPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service delete project pipeline o k response has a 5xx status code
func (o *KafeidoServiceDeleteProjectPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service delete project pipeline o k response a status code equal to that given
func (o *KafeidoServiceDeleteProjectPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service delete project pipeline o k response
func (o *KafeidoServiceDeleteProjectPipelineOK) Code() int {
	return 200
}

func (o *KafeidoServiceDeleteProjectPipelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoServiceDeleteProjectPipelineOK %s", 200, payload)
}

func (o *KafeidoServiceDeleteProjectPipelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoServiceDeleteProjectPipelineOK %s", 200, payload)
}

func (o *KafeidoServiceDeleteProjectPipelineOK) GetPayload() models.KafeidoDeleteProjectPipelineResponse {
	return o.Payload
}

func (o *KafeidoServiceDeleteProjectPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceDeleteProjectPipelineDefault creates a KafeidoServiceDeleteProjectPipelineDefault with default headers values
func NewKafeidoServiceDeleteProjectPipelineDefault(code int) *KafeidoServiceDeleteProjectPipelineDefault {
	return &KafeidoServiceDeleteProjectPipelineDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceDeleteProjectPipelineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceDeleteProjectPipelineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service delete project pipeline default response has a 2xx status code
func (o *KafeidoServiceDeleteProjectPipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service delete project pipeline default response has a 3xx status code
func (o *KafeidoServiceDeleteProjectPipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service delete project pipeline default response has a 4xx status code
func (o *KafeidoServiceDeleteProjectPipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service delete project pipeline default response has a 5xx status code
func (o *KafeidoServiceDeleteProjectPipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service delete project pipeline default response a status code equal to that given
func (o *KafeidoServiceDeleteProjectPipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service delete project pipeline default response
func (o *KafeidoServiceDeleteProjectPipelineDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceDeleteProjectPipelineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] KafeidoService_DeleteProjectPipeline default %s", o._statusCode, payload)
}

func (o *KafeidoServiceDeleteProjectPipelineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] KafeidoService_DeleteProjectPipeline default %s", o._statusCode, payload)
}

func (o *KafeidoServiceDeleteProjectPipelineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceDeleteProjectPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
