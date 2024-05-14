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

// KafeidoPredictAudioRequestBody kafeido predict audio request body
//
// swagger:model kafeidoPredictAudioRequestBody
type KafeidoPredictAudioRequestBody struct {

	// audio bytes
	AudioBytes *KafeidoRequestInBytes `json:"audioBytes,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// lang
	Lang []string `json:"lang"`
}

// Validate validates this kafeido predict audio request body
func (m *KafeidoPredictAudioRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoPredictAudioRequestBody) validateAudioBytes(formats strfmt.Registry) error {
	if swag.IsZero(m.AudioBytes) { // not required
		return nil
	}

	if m.AudioBytes != nil {
		if err := m.AudioBytes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioBytes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioBytes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido predict audio request body based on the context it is used
func (m *KafeidoPredictAudioRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudioBytes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoPredictAudioRequestBody) contextValidateAudioBytes(ctx context.Context, formats strfmt.Registry) error {

	if m.AudioBytes != nil {

		if swag.IsZero(m.AudioBytes) { // not required
			return nil
		}

		if err := m.AudioBytes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioBytes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioBytes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoPredictAudioRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoPredictAudioRequestBody) UnmarshalBinary(b []byte) error {
	var res KafeidoPredictAudioRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
