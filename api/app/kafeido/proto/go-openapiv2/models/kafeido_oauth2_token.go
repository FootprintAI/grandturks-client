// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KafeidoOauth2Token kafeido oauth2 token
//
// swagger:model kafeidoOauth2Token
type KafeidoOauth2Token struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`
}

// Validate validates this kafeido oauth2 token
func (m *KafeidoOauth2Token) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido oauth2 token based on context it is used
func (m *KafeidoOauth2Token) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoOauth2Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoOauth2Token) UnmarshalBinary(b []byte) error {
	var res KafeidoOauth2Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
