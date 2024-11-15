// Code generated by go-swagger; DO NOT EDIT.

package kafeido_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoServicePutProjectPipelineReader is a Reader for the KafeidoServicePutProjectPipeline structure.
type KafeidoServicePutProjectPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServicePutProjectPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServicePutProjectPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServicePutProjectPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServicePutProjectPipelineOK creates a KafeidoServicePutProjectPipelineOK with default headers values
func NewKafeidoServicePutProjectPipelineOK() *KafeidoServicePutProjectPipelineOK {
	return &KafeidoServicePutProjectPipelineOK{}
}

/*
KafeidoServicePutProjectPipelineOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServicePutProjectPipelineOK struct {
	Payload *models.KafeidoPutProjectPipelineResponse
}

// IsSuccess returns true when this kafeido service put project pipeline o k response has a 2xx status code
func (o *KafeidoServicePutProjectPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service put project pipeline o k response has a 3xx status code
func (o *KafeidoServicePutProjectPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service put project pipeline o k response has a 4xx status code
func (o *KafeidoServicePutProjectPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service put project pipeline o k response has a 5xx status code
func (o *KafeidoServicePutProjectPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service put project pipeline o k response a status code equal to that given
func (o *KafeidoServicePutProjectPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service put project pipeline o k response
func (o *KafeidoServicePutProjectPipelineOK) Code() int {
	return 200
}

func (o *KafeidoServicePutProjectPipelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoServicePutProjectPipelineOK %s", 200, payload)
}

func (o *KafeidoServicePutProjectPipelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] kafeidoServicePutProjectPipelineOK %s", 200, payload)
}

func (o *KafeidoServicePutProjectPipelineOK) GetPayload() *models.KafeidoPutProjectPipelineResponse {
	return o.Payload
}

func (o *KafeidoServicePutProjectPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoPutProjectPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServicePutProjectPipelineDefault creates a KafeidoServicePutProjectPipelineDefault with default headers values
func NewKafeidoServicePutProjectPipelineDefault(code int) *KafeidoServicePutProjectPipelineDefault {
	return &KafeidoServicePutProjectPipelineDefault{
		_statusCode: code,
	}
}

/*
KafeidoServicePutProjectPipelineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServicePutProjectPipelineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service put project pipeline default response has a 2xx status code
func (o *KafeidoServicePutProjectPipelineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service put project pipeline default response has a 3xx status code
func (o *KafeidoServicePutProjectPipelineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service put project pipeline default response has a 4xx status code
func (o *KafeidoServicePutProjectPipelineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service put project pipeline default response has a 5xx status code
func (o *KafeidoServicePutProjectPipelineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service put project pipeline default response a status code equal to that given
func (o *KafeidoServicePutProjectPipelineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service put project pipeline default response
func (o *KafeidoServicePutProjectPipelineDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServicePutProjectPipelineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] KafeidoService_PutProjectPipeline default %s", o._statusCode, payload)
}

func (o *KafeidoServicePutProjectPipelineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}/pipelines/{namedPipelineId}][%d] KafeidoService_PutProjectPipeline default %s", o._statusCode, payload)
}

func (o *KafeidoServicePutProjectPipelineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServicePutProjectPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
KafeidoServicePutProjectPipelineBody kafeido service put project pipeline body
swagger:model KafeidoServicePutProjectPipelineBody
*/
type KafeidoServicePutProjectPipelineBody struct {

	// raw yaml string
	// Format: byte
	PipelineWorkflowInRawYaml strfmt.Base64 `json:"pipelineWorkflowInRawYaml,omitempty"`
}

// Validate validates this kafeido service put project pipeline body
func (o *KafeidoServicePutProjectPipelineBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido service put project pipeline body based on context it is used
func (o *KafeidoServicePutProjectPipelineBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KafeidoServicePutProjectPipelineBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KafeidoServicePutProjectPipelineBody) UnmarshalBinary(b []byte) error {
	var res KafeidoServicePutProjectPipelineBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
