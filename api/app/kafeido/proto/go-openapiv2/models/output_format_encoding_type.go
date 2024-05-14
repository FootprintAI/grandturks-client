// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OutputFormatEncodingType output format encoding type
//
// swagger:model OutputFormatEncodingType
type OutputFormatEncodingType string

func NewOutputFormatEncodingType(value OutputFormatEncodingType) *OutputFormatEncodingType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OutputFormatEncodingType.
func (m OutputFormatEncodingType) Pointer() *OutputFormatEncodingType {
	return &m
}

const (

	// OutputFormatEncodingTypeNone captures enum value "None"
	OutputFormatEncodingTypeNone OutputFormatEncodingType = "None"

	// OutputFormatEncodingTypeBASE64 captures enum value "BASE64"
	OutputFormatEncodingTypeBASE64 OutputFormatEncodingType = "BASE64"
)

// for schema
var outputFormatEncodingTypeEnum []interface{}

func init() {
	var res []OutputFormatEncodingType
	if err := json.Unmarshal([]byte(`["None","BASE64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outputFormatEncodingTypeEnum = append(outputFormatEncodingTypeEnum, v)
	}
}

func (m OutputFormatEncodingType) validateOutputFormatEncodingTypeEnum(path, location string, value OutputFormatEncodingType) error {
	if err := validate.EnumCase(path, location, value, outputFormatEncodingTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this output format encoding type
func (m OutputFormatEncodingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOutputFormatEncodingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this output format encoding type based on context it is used
func (m OutputFormatEncodingType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
