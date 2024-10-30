// Code generated by go-swagger; DO NOT EDIT.

package kafeido_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// KafeidoServicePutProjectReader is a Reader for the KafeidoServicePutProject structure.
type KafeidoServicePutProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KafeidoServicePutProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKafeidoServicePutProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKafeidoServicePutProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKafeidoServicePutProjectOK creates a KafeidoServicePutProjectOK with default headers values
func NewKafeidoServicePutProjectOK() *KafeidoServicePutProjectOK {
	return &KafeidoServicePutProjectOK{}
}

/*
KafeidoServicePutProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type KafeidoServicePutProjectOK struct {
	Payload models.KafeidoPutProjectResponse
}

// IsSuccess returns true when this kafeido service put project o k response has a 2xx status code
func (o *KafeidoServicePutProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kafeido service put project o k response has a 3xx status code
func (o *KafeidoServicePutProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kafeido service put project o k response has a 4xx status code
func (o *KafeidoServicePutProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kafeido service put project o k response has a 5xx status code
func (o *KafeidoServicePutProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kafeido service put project o k response a status code equal to that given
func (o *KafeidoServicePutProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kafeido service put project o k response
func (o *KafeidoServicePutProjectOK) Code() int {
	return 200
}

func (o *KafeidoServicePutProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}][%d] kafeidoServicePutProjectOK %s", 200, payload)
}

func (o *KafeidoServicePutProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}][%d] kafeidoServicePutProjectOK %s", 200, payload)
}

func (o *KafeidoServicePutProjectOK) GetPayload() models.KafeidoPutProjectResponse {
	return o.Payload
}

func (o *KafeidoServicePutProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKafeidoServicePutProjectDefault creates a KafeidoServicePutProjectDefault with default headers values
func NewKafeidoServicePutProjectDefault(code int) *KafeidoServicePutProjectDefault {
	return &KafeidoServicePutProjectDefault{
		_statusCode: code,
	}
}

/*
KafeidoServicePutProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type KafeidoServicePutProjectDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this kafeido service put project default response has a 2xx status code
func (o *KafeidoServicePutProjectDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this kafeido service put project default response has a 3xx status code
func (o *KafeidoServicePutProjectDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this kafeido service put project default response has a 4xx status code
func (o *KafeidoServicePutProjectDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this kafeido service put project default response has a 5xx status code
func (o *KafeidoServicePutProjectDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this kafeido service put project default response a status code equal to that given
func (o *KafeidoServicePutProjectDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the kafeido service put project default response
func (o *KafeidoServicePutProjectDefault) Code() int {
	return o._statusCode
}

func (o *KafeidoServicePutProjectDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}][%d] KafeidoService_PutProject default %s", o._statusCode, payload)
}

func (o *KafeidoServicePutProjectDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/projects/{projectId}][%d] KafeidoService_PutProject default %s", o._statusCode, payload)
}

func (o *KafeidoServicePutProjectDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *KafeidoServicePutProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
KafeidoServicePutProjectBody kafeido service put project body
swagger:model KafeidoServicePutProjectBody
*/
type KafeidoServicePutProjectBody struct {

	// default public
	Visibility *models.AuthzTypeVisibility `json:"visibility,omitempty"`
}

// Validate validates this kafeido service put project body
func (o *KafeidoServicePutProjectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KafeidoServicePutProjectBody) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(o.Visibility) { // not required
		return nil
	}

	if o.Visibility != nil {
		if err := o.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "visibility")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido service put project body based on the context it is used
func (o *KafeidoServicePutProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVisibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KafeidoServicePutProjectBody) contextValidateVisibility(ctx context.Context, formats strfmt.Registry) error {

	if o.Visibility != nil {

		if swag.IsZero(o.Visibility) { // not required
			return nil
		}

		if err := o.Visibility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KafeidoServicePutProjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KafeidoServicePutProjectBody) UnmarshalBinary(b []byte) error {
	var res KafeidoServicePutProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}