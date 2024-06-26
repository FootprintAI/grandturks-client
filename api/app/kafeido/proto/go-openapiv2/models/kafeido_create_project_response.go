// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KafeidoCreateProjectResponse kafeido create project response
//
// swagger:model kafeidoCreateProjectResponse
type KafeidoCreateProjectResponse struct {

	// given when a project has been created
	ProjectID string `json:"projectId,omitempty"`
}

// Validate validates this kafeido create project response
func (m *KafeidoCreateProjectResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido create project response based on context it is used
func (m *KafeidoCreateProjectResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoCreateProjectResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoCreateProjectResponse) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateProjectResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
