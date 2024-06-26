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

// KafeidoCreateUserRequest kafeido create user request
//
// swagger:model kafeidoCreateUserRequest
type KafeidoCreateUserRequest struct {

	// oauth2 token
	Oauth2Token *KafeidoOauth2Token `json:"oauth2Token,omitempty"`
}

// Validate validates this kafeido create user request
func (m *KafeidoCreateUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOauth2Token(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoCreateUserRequest) validateOauth2Token(formats strfmt.Registry) error {
	if swag.IsZero(m.Oauth2Token) { // not required
		return nil
	}

	if m.Oauth2Token != nil {
		if err := m.Oauth2Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2Token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2Token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido create user request based on the context it is used
func (m *KafeidoCreateUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOauth2Token(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoCreateUserRequest) contextValidateOauth2Token(ctx context.Context, formats strfmt.Registry) error {

	if m.Oauth2Token != nil {

		if swag.IsZero(m.Oauth2Token) { // not required
			return nil
		}

		if err := m.Oauth2Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2Token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2Token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoCreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoCreateUserRequest) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
