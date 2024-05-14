// Code generated by go-swagger; DO NOT EDIT.

package kafeido

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoCreateProjectPipelineReader is a Reader for the KafeidoCreateProjectPipeline structure.
type KafeidoCreateProjectPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoCreateProjectPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoCreateProjectPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoCreateProjectPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoCreateProjectPipelineOK creates a KafeidoCreateProjectPipelineOK with default headers values
func NewKafeidoCreateProjectPipelineOK() *KafeidoCreateProjectPipelineOK {
	return &KafeidoCreateProjectPipelineOK{}
}

/*
KafeidoCreateProjectPipelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoCreateProjectPipelineOK struct {
	Payload *models.KafeidoCreateProjectPipelineResponse
}

// IsSuccess returns true when this kafeido create project pipeline o k response has a 2xx status code
func (o *KafeidoCreateProjectPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido create project pipeline o k response has a 3xx status code
func (o *KafeidoCreateProjectPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido create project pipeline o k response has a 4xx status code
func (o *KafeidoCreateProjectPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido create project pipeline o k response has a 5xx status code
func (o *KafeidoCreateProjectPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido create project pipeline o k response a status code equal to that given
func (o *KafeidoCreateProjectPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido create project pipeline o k response
func (o *KafeidoCreateProjectPipelineOK) Code() int {
	return 200
}

func (o *KafeidoCreateProjectPipelineOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/pipelines][%d] kafeidoCreateProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoCreateProjectPipelineOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/pipelines][%d] kafeidoCreateProjectPipelineOK  %+v", 200, o.Payload)
}

func (o *KafeidoCreateProjectPipelineOK) GetPayload() *models.KafeidoCreateProjectPipelineResponse {
	return o.Payload
}

func (o *KafeidoCreateProjectPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoCreateProjectPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoCreateProjectPipelineDefault creates a KafeidoCreateProjectPipelineDefault with default headers values
func NewKafeidoCreateProjectPipelineDefault(code int) *KafeidoCreateProjectPipelineDefault {
	return &KafeidoCreateProjectPipelineDefault{
		_statusCode: code,
	}
}

/*
KafeidoCreateProjectPipelineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoCreateProjectPipelineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido create project pipeline default response has a 2xx status code
func (o *KafeidoCreateProjectPipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido create project pipeline default response has a 3xx status code
func (o *KafeidoCreateProjectPipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido create project pipeline default response has a 4xx status code
func (o *KafeidoCreateProjectPipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido create project pipeline default response has a 5xx status code
func (o *KafeidoCreateProjectPipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido create project pipeline default response a status code equal to that given
func (o *KafeidoCreateProjectPipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido create project pipeline default response
func (o *KafeidoCreateProjectPipelineDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoCreateProjectPipelineDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/pipelines][%d] kafeido_CreateProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoCreateProjectPipelineDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/pipelines][%d] kafeido_CreateProjectPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoCreateProjectPipelineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoCreateProjectPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
KafeidoCreateProjectPipelineBody kafeido create project pipeline body
swagger:model KafeidoCreateProjectPipelineBody
*/
type KafeidoCreateProjectPipelineBody struct {

	// desc
	Desc string `json:"desc,omitempty"`

	// kf pipeline workflow
	// Format: byte
	KfPipelineWorkflow strfmt.Base64 `json:"kfPipelineWorkflow,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this kafeido create project pipeline body
func (o *KafeidoCreateProjectPipelineBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido create project pipeline body based on context it is used
func (o *KafeidoCreateProjectPipelineBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KafeidoCreateProjectPipelineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KafeidoCreateProjectPipelineBody) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateProjectPipelineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
