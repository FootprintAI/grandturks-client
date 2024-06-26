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

// KafeidoCreateUserResponse kafeido create user response
//
// swagger:model kafeidoCreateUserResponse
type KafeidoCreateUserResponse struct {

	// user info
	UserInfo *UserstoreUserInfo `json:"userInfo,omitempty"`
}

// Validate validates this kafeido create user response
func (m *KafeidoCreateUserResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoCreateUserResponse) validateUserInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {
		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido create user response based on the context it is used
func (m *KafeidoCreateUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoCreateUserResponse) contextValidateUserInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserInfo != nil {

		if swag.IsZero(m.UserInfo) { // not required
			return nil
		}

		if err := m.UserInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoCreateUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoCreateUserResponse) UnmarshalBinary(b []byte) error {
	var res KafeidoCreateUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
