// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatastreamImageURLDataSource datastream image Url data source
//
// swagger:model datastreamImageUrlDataSource
type DatastreamImageURLDataSource struct {

	// streaming duration
	StreamingDuration string `json:"streamingDuration,omitempty"`

	// https://example.com/1.jpg
	URL string `json:"url,omitempty"`
}

// Validate validates this datastream image Url data source
func (m *DatastreamImageURLDataSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datastream image Url data source based on context it is used
func (m *DatastreamImageURLDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatastreamImageURLDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastreamImageURLDataSource) UnmarshalBinary(b []byte) error {
	var res DatastreamImageURLDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
