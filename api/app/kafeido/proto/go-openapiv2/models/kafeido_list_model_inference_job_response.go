// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KafeidoListModelInferenceJobResponse kafeido list model inference job response
//
// swagger:model kafeidoListModelInferenceJobResponse
type KafeidoListModelInferenceJobResponse struct {

	// model inference job ids
	ModelInferenceJobIds []string `json:"modelInferenceJobIds"`
}

// Validate validates this kafeido list model inference job response
func (m *KafeidoListModelInferenceJobResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido list model inference job response based on context it is used
func (m *KafeidoListModelInferenceJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoListModelInferenceJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoListModelInferenceJobResponse) UnmarshalBinary(b []byte) error {
	var res KafeidoListModelInferenceJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}