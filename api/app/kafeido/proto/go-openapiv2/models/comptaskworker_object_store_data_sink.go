// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComptaskworkerObjectStoreDataSink comptaskworker object store data sink
//
// swagger:model comptaskworkerObjectStoreDataSink
type ComptaskworkerObjectStoreDataSink struct {

	// bucket name
	BucketName string `json:"bucketName,omitempty"`

	// object prefix
	ObjectPrefix string `json:"objectPrefix,omitempty"`

	// object store info
	ObjectStoreInfo *ComptaskworkerObjectStoreInfo `json:"objectStoreInfo,omitempty"`
}

// Validate validates this comptaskworker object store data sink
func (m *ComptaskworkerObjectStoreDataSink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectStoreInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComptaskworkerObjectStoreDataSink) validateObjectStoreInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectStoreInfo) { // not required
		return nil
	}

	if m.ObjectStoreInfo != nil {
		if err := m.ObjectStoreInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectStoreInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectStoreInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this comptaskworker object store data sink based on the context it is used
func (m *ComptaskworkerObjectStoreDataSink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectStoreInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComptaskworkerObjectStoreDataSink) contextValidateObjectStoreInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectStoreInfo != nil {

		if swag.IsZero(m.ObjectStoreInfo) { // not required
			return nil
		}

		if err := m.ObjectStoreInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectStoreInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectStoreInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComptaskworkerObjectStoreDataSink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComptaskworkerObjectStoreDataSink) UnmarshalBinary(b []byte) error {
	var res ComptaskworkerObjectStoreDataSink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
