// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KafeidoCreateProjectPipelineResponse kafeido create project pipeline response
//
// swagger:model kafeidoCreateProjectPipelineResponse
type KafeidoCreateProjectPipelineResponse struct {

	// named pipeline Id
	NamedPipelineID string `json:"namedPipelineId,omitempty"`
}

// Validate validates this kafeido create project pipeline response
func (m *KafeidoCreateProjectPipelineResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido create project pipeline response based on context it is used
func (m *KafeidoCreateProjectPipelineResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoCreateProjectPipelineResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoCreateProjectPipelineResponse) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateProjectPipelineResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
