// Code generated by go-swagger; DO NOT EDIT.

package kafeido

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoCreateModelInferenceJobReader is a Reader for the KafeidoCreateModelInferenceJob structure.
type KafeidoCreateModelInferenceJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoCreateModelInferenceJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoCreateModelInferenceJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoCreateModelInferenceJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoCreateModelInferenceJobOK creates a KafeidoCreateModelInferenceJobOK with default headers values
func NewKafeidoCreateModelInferenceJobOK() *KafeidoCreateModelInferenceJobOK {
	return &KafeidoCreateModelInferenceJobOK{}
}

/*
KafeidoCreateModelInferenceJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoCreateModelInferenceJobOK struct {
	Payload *models.KafeidoCreateModelInferenceJobResponse
}

// IsSuccess returns true when this kafeido create model inference job o k response has a 2xx status code
func (o *KafeidoCreateModelInferenceJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido create model inference job o k response has a 3xx status code
func (o *KafeidoCreateModelInferenceJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido create model inference job o k response has a 4xx status code
func (o *KafeidoCreateModelInferenceJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido create model inference job o k response has a 5xx status code
func (o *KafeidoCreateModelInferenceJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido create model inference job o k response a status code equal to that given
func (o *KafeidoCreateModelInferenceJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido create model inference job o k response
func (o *KafeidoCreateModelInferenceJobOK) Code() int {
	return 200
}

func (o *KafeidoCreateModelInferenceJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/job][%d] kafeidoCreateModelInferenceJobOK  %+v", 200, o.Payload)
}

func (o *KafeidoCreateModelInferenceJobOK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/job][%d] kafeidoCreateModelInferenceJobOK  %+v", 200, o.Payload)
}

func (o *KafeidoCreateModelInferenceJobOK) GetPayload() *models.KafeidoCreateModelInferenceJobResponse {
	return o.Payload
}

func (o *KafeidoCreateModelInferenceJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KafeidoCreateModelInferenceJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoCreateModelInferenceJobDefault creates a KafeidoCreateModelInferenceJobDefault with default headers values
func NewKafeidoCreateModelInferenceJobDefault(code int) *KafeidoCreateModelInferenceJobDefault {
	return &KafeidoCreateModelInferenceJobDefault{
		_statusCode: code,
	}
}

/*
KafeidoCreateModelInferenceJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoCreateModelInferenceJobDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido create model inference job default response has a 2xx status code
func (o *KafeidoCreateModelInferenceJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido create model inference job default response has a 3xx status code
func (o *KafeidoCreateModelInferenceJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido create model inference job default response has a 4xx status code
func (o *KafeidoCreateModelInferenceJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido create model inference job default response has a 5xx status code
func (o *KafeidoCreateModelInferenceJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido create model inference job default response a status code equal to that given
func (o *KafeidoCreateModelInferenceJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido create model inference job default response
func (o *KafeidoCreateModelInferenceJobDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoCreateModelInferenceJobDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/job][%d] kafeido_CreateModelInferenceJob default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoCreateModelInferenceJobDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects/{projectId}/job][%d] kafeido_CreateModelInferenceJob default  %+v", o._statusCode, o.Payload)
}

func (o *KafeidoCreateModelInferenceJobDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoCreateModelInferenceJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
KafeidoCreateModelInferenceJobBody kafeido create model inference job body
swagger:model KafeidoCreateModelInferenceJobBody
*/
type KafeidoCreateModelInferenceJobBody struct {

	// concurrent requests
	ConcurrentRequests int32 `json:"concurrentRequests,omitempty"`

	// data sink config
	DataSinkConfig *models.KafeidoDataSinkConfig `json:"dataSinkConfig,omitempty"`

	// data source Id
	DataSourceID string `json:"dataSourceId,omitempty"`

	// model inference Id
	ModelInferenceID string `json:"modelInferenceId,omitempty"`

	// output format
	OutputFormat *models.DatastreamOutputFormat `json:"outputFormat,omitempty"`

	// prometheus job label
	PrometheusJobLabel string `json:"prometheusJobLabel,omitempty"`
}

// Validate validates this kafeido create model inference job body
func (o *KafeidoCreateModelInferenceJobBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataSinkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOutputFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KafeidoCreateModelInferenceJobBody) validateDataSinkConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.DataSinkConfig) { // not required
		return nil
	}

	if o.DataSinkConfig != nil {
		if err := o.DataSinkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "dataSinkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "dataSinkConfig")
			}
			return err
		}
	}

	return nil
}

func (o *KafeidoCreateModelInferenceJobBody) validateOutputFormat(formats strfmt.Registry) error {
	if swag.IsZero(o.OutputFormat) { // not required
		return nil
	}

	if o.OutputFormat != nil {
		if err := o.OutputFormat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "outputFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "outputFormat")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido create model inference job body based on the context it is used
func (o *KafeidoCreateModelInferenceJobBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDataSinkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOutputFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KafeidoCreateModelInferenceJobBody) contextValidateDataSinkConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.DataSinkConfig != nil {

		if swag.IsZero(o.DataSinkConfig) { // not required
			return nil
		}

		if err := o.DataSinkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "dataSinkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "dataSinkConfig")
			}
			return err
		}
	}

	return nil
}

func (o *KafeidoCreateModelInferenceJobBody) contextValidateOutputFormat(ctx context.Context, formats strfmt.Registry) error {

	if o.OutputFormat != nil {

		if swag.IsZero(o.OutputFormat) { // not required
			return nil
		}

		if err := o.OutputFormat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "outputFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "outputFormat")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KafeidoCreateModelInferenceJobBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KafeidoCreateModelInferenceJobBody) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateModelInferenceJobBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
