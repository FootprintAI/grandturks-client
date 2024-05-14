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

// KafeidoListDataSourceReader is a Reader for the KafeidoListDataSource structure.
type KafeidoListDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoListDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoListDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoListDataSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoListDataSourceOK creates a KafeidoListDataSourceOK with default headers values
func NewKafeidoListDataSourceOK() *KafeidoListDataSourceOK {
	return &KafeidoListDataSourceOK{}
}

/*
KafeidoListDataSourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoListDataSourceOK struct {
	Payload *models.AppkafeidoListDataSourceResponse
}

// IsSuccess returns true when this kafeido list data source o k response has a 2xx status code
func (o *KafeidoListDataSourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido list data source o k response has a 3xx status code
func (o *KafeidoListDataSourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido list data source o k response has a 4xx status code
func (o *KafeidoListDataSourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido list data source o k response has a 5xx status code
func (o *KafeidoListDataSourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido list data source o k response a status code equal to that given
func (o *KafeidoListDataSourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido list data source o k response
func (o *KafeidoListDataSourceOK) Code() int {
	return 200
}

func (o *KafeidoListDataSourceOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource][%d] kafeidoListDataSourceOK  %+v", 200, o.Payload)
}

func (o *KafeidoListDataSourceOK) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource][%d] kafeidoListDataSourceOK  %+v", 200, o.Payload)
}

func (o *KafeidoListDataSourceOK) GetPayload() *models.AppkafeidoListDataSourceResponse {
	return o.Payload
}

func (o *KafeidoListDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppkafeidoListDataSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoListDataSourceDefault creates a KafeidoListDataSourceDefault with default headers values
func NewKafeidoListDataSourceDefault(code int) *KafeidoListDataSourceDefault {
	return &KafeidoListDataSourceDefault{
		_statusCode: code,
	}
}

/*
KafeidoListDataSourceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoListDataSourceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido list data source default response has a 2xx status code
func (o *KafeidoListDataSourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido list data source default response has a 3xx status code
func (o *KafeidoListDataSourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido list data source default response has a 4xx status code
func (o *KafeidoListDataSourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido list data source default response has a 5xx status code
func (o *KafeidoListDataSourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido list data source default response a status code equal to that given
func (o *KafeidoListDataSourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido list data source default response
func (o *KafeidoListDataSourceDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoListDataSourceDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource][%d] kafeido_ListDataSource default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoListDataSourceDefault) String() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/datasource][%d] kafeido_ListDataSource default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoListDataSourceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoListDataSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}