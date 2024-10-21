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

// UserstoreApplicationID userstore application ID
//
// swagger:model userstoreApplicationID
type UserstoreApplicationID string

func NewUserstoreApplicationID(value UserstoreApplicationID) *UserstoreApplicationID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserstoreApplicationID.
func (m UserstoreApplicationID) Pointer() *UserstoreApplicationID {
	return &m
}

const (

	// UserstoreApplicationIDAPPLICATIONIDUNSPECIFIED captures enum value "APPLICATION_ID_UNSPECIFIED"
	UserstoreApplicationIDAPPLICATIONIDUNSPECIFIED UserstoreApplicationID = "APPLICATION_ID_UNSPECIFIED"

	// UserstoreApplicationIDAPPLICATIONIDKAFEIDO captures enum value "APPLICATION_ID_KAFEIDO"
	UserstoreApplicationIDAPPLICATIONIDKAFEIDO UserstoreApplicationID = "APPLICATION_ID_KAFEIDO"
)

// for schema
var userstoreApplicationIdEnum []interface{}

func init() {
	var res []UserstoreApplicationID
	if err := json.Unmarshal([]byte(`["APPLICATION_ID_UNSPECIFIED","APPLICATION_ID_KAFEIDO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userstoreApplicationIdEnum = append(userstoreApplicationIdEnum, v)
	}
}

func (m UserstoreApplicationID) validateUserstoreApplicationIDEnum(path, location string, value UserstoreApplicationID) error {
	if err := validate.EnumCase(path, location, value, userstoreApplicationIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this userstore application ID
func (m UserstoreApplicationID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserstoreApplicationIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this userstore application ID based on context it is used
func (m UserstoreApplicationID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
