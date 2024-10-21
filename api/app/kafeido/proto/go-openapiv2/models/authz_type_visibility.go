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

// AuthzTypeVisibility - TYPE_VISIBILITY_PUBLIC: everyone can see it
//   - TYPE_VISIBILITY_TEAM: limited group of people can see it
//   - TYPE_VISIBILITY_ONLYOWNER: only owner can see it
//
// swagger:model authzTypeVisibility
type AuthzTypeVisibility string

func NewAuthzTypeVisibility(value AuthzTypeVisibility) *AuthzTypeVisibility {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AuthzTypeVisibility.
func (m AuthzTypeVisibility) Pointer() *AuthzTypeVisibility {
	return &m
}

const (

	// AuthzTypeVisibilityTYPEVISIBILITYUNSPECIFIED captures enum value "TYPE_VISIBILITY_UNSPECIFIED"
	AuthzTypeVisibilityTYPEVISIBILITYUNSPECIFIED AuthzTypeVisibility = "TYPE_VISIBILITY_UNSPECIFIED"

	// AuthzTypeVisibilityTYPEVISIBILITYPUBLIC captures enum value "TYPE_VISIBILITY_PUBLIC"
	AuthzTypeVisibilityTYPEVISIBILITYPUBLIC AuthzTypeVisibility = "TYPE_VISIBILITY_PUBLIC"

	// AuthzTypeVisibilityTYPEVISIBILITYTEAM captures enum value "TYPE_VISIBILITY_TEAM"
	AuthzTypeVisibilityTYPEVISIBILITYTEAM AuthzTypeVisibility = "TYPE_VISIBILITY_TEAM"

	// AuthzTypeVisibilityTYPEVISIBILITYONLYOWNER captures enum value "TYPE_VISIBILITY_ONLYOWNER"
	AuthzTypeVisibilityTYPEVISIBILITYONLYOWNER AuthzTypeVisibility = "TYPE_VISIBILITY_ONLYOWNER"
)

// for schema
var authzTypeVisibilityEnum []interface{}

func init() {
	var res []AuthzTypeVisibility
	if err := json.Unmarshal([]byte(`["TYPE_VISIBILITY_UNSPECIFIED","TYPE_VISIBILITY_PUBLIC","TYPE_VISIBILITY_TEAM","TYPE_VISIBILITY_ONLYOWNER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authzTypeVisibilityEnum = append(authzTypeVisibilityEnum, v)
	}
}

func (m AuthzTypeVisibility) validateAuthzTypeVisibilityEnum(path, location string, value AuthzTypeVisibility) error {
	if err := validate.EnumCase(path, location, value, authzTypeVisibilityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this authz type visibility
func (m AuthzTypeVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthzTypeVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this authz type visibility based on context it is used
func (m AuthzTypeVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
