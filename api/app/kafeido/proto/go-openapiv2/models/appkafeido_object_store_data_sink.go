// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppkafeidoObjectStoreDataSink appkafeido object store data sink
//
// swagger:model appkafeidoObjectStoreDataSink
type AppkafeidoObjectStoreDataSink struct {

	// bucket name
	BucketName string `json:"BucketName,omitempty"`

	// object prefix
	ObjectPrefix string `json:"ObjectPrefix,omitempty"`
}

// Validate validates this appkafeido object store data sink
func (m *AppkafeidoObjectStoreDataSink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this appkafeido object store data sink based on context it is used
func (m *AppkafeidoObjectStoreDataSink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppkafeidoObjectStoreDataSink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppkafeidoObjectStoreDataSink) UnmarshalBinary(b []byte) error {
	var res AppkafeidoObjectStoreDataSink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}