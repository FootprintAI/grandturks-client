// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaskworkerRestColDataSink taskworker rest col data sink
//
// swagger:model taskworkerRestColDataSink
type TaskworkerRestColDataSink struct {

	// collection Id
	CollectionID string `json:"collectionId,omitempty"`
}

// Validate validates this taskworker rest col data sink
func (m *TaskworkerRestColDataSink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this taskworker rest col data sink based on context it is used
func (m *TaskworkerRestColDataSink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskworkerRestColDataSink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskworkerRestColDataSink) UnmarshalBinary(b []byte) error {
	var res TaskworkerRestColDataSink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
