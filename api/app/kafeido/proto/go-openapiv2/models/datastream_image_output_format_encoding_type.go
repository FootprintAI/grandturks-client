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

// DatastreamImageOutputFormatEncodingType datastream image output format encoding type
//
// swagger:model datastreamImageOutputFormatEncodingType
type DatastreamImageOutputFormatEncodingType string

func NewDatastreamImageOutputFormatEncodingType(value DatastreamImageOutputFormatEncodingType) *DatastreamImageOutputFormatEncodingType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DatastreamImageOutputFormatEncodingType.
func (m DatastreamImageOutputFormatEncodingType) Pointer() *DatastreamImageOutputFormatEncodingType {
	return &m
}

const (

	// DatastreamImageOutputFormatEncodingTypeNone captures enum value "None"
	DatastreamImageOutputFormatEncodingTypeNone DatastreamImageOutputFormatEncodingType = "None"

	// DatastreamImageOutputFormatEncodingTypeBASE64 captures enum value "BASE64"
	DatastreamImageOutputFormatEncodingTypeBASE64 DatastreamImageOutputFormatEncodingType = "BASE64"
)

// for schema
var datastreamImageOutputFormatEncodingTypeEnum []interface{}

func init() {
	var res []DatastreamImageOutputFormatEncodingType
	if err := json.Unmarshal([]byte(`["None","BASE64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datastreamImageOutputFormatEncodingTypeEnum = append(datastreamImageOutputFormatEncodingTypeEnum, v)
	}
}

func (m DatastreamImageOutputFormatEncodingType) validateDatastreamImageOutputFormatEncodingTypeEnum(path, location string, value DatastreamImageOutputFormatEncodingType) error {
	if err := validate.EnumCase(path, location, value, datastreamImageOutputFormatEncodingTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datastream image output format encoding type
func (m DatastreamImageOutputFormatEncodingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatastreamImageOutputFormatEncodingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this datastream image output format encoding type based on context it is used
func (m DatastreamImageOutputFormatEncodingType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}