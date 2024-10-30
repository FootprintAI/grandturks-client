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

// KafeidoServiceGetDataSourceReader is a Reader for the KafeidoServiceGetDataSource structure.
type KafeidoServiceGetDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServiceGetDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServiceGetDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServiceGetDataSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServiceGetDataSourceOK creates a KafeidoServiceGetDataSourceOK with default headers values
func NewKafeidoServiceGetDataSourceOK() *KafeidoServiceGetDataSourceOK {
	return &KafeidoServiceGetDataSourceOK{}
}

/*
KafeidoServiceGetDataSourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServiceGetDataSourceOK struct {
	Payload *models.AppkafeidoGetDataSourceResponse
}

// IsSuccess returns true when this kafeido service get data source o k response has a 2xx status code
func (o *KafeidoServiceGetDataSourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service get data source o k response has a 3xx status code
func (o *KafeidoServiceGetDataSourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service get data source o k response has a 4xx status code
func (o *KafeidoServiceGetDataSourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service get data source o k response has a 5xx status code
func (o *KafeidoServiceGetDataSourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service get data source o k response a status code equal to that given
func (o *KafeidoServiceGetDataSourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service get data source o k response
func (o *KafeidoServiceGetDataSourceOK) Code() int {
	return 200
}

func (o *KafeidoServiceGetDataSourceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource/{dataSourceId}][%d] kafeidoServiceGetDataSourceOK %s", 200, payload)
}

func (o *KafeidoServiceGetDataSourceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource/{dataSourceId}][%d] kafeidoServiceGetDataSourceOK %s", 200, payload)
}

func (o *KafeidoServiceGetDataSourceOK) GetPayload() *models.AppkafeidoGetDataSourceResponse {
	return o.Payload
}

func (o *KafeidoServiceGetDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppkafeidoGetDataSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServiceGetDataSourceDefault creates a KafeidoServiceGetDataSourceDefault with default headers values
func NewKafeidoServiceGetDataSourceDefault(code int) *KafeidoServiceGetDataSourceDefault {
	return &KafeidoServiceGetDataSourceDefault{
		_statusCode: code,
	}
}

/*
KafeidoServiceGetDataSourceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServiceGetDataSourceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service get data source default response has a 2xx status code
func (o *KafeidoServiceGetDataSourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service get data source default response has a 3xx status code
func (o *KafeidoServiceGetDataSourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service get data source default response has a 4xx status code
func (o *KafeidoServiceGetDataSourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service get data source default response has a 5xx status code
func (o *KafeidoServiceGetDataSourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service get data source default response a status code equal to that given
func (o *KafeidoServiceGetDataSourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service get data source default response
func (o *KafeidoServiceGetDataSourceDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServiceGetDataSourceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource/{dataSourceId}][%d] KafeidoService_GetDataSource default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetDataSourceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource/{dataSourceId}][%d] KafeidoService_GetDataSource default %s", o._statusCode, payload)
}

func (o *KafeidoServiceGetDataSourceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServiceGetDataSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}