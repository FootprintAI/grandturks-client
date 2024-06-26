// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KafeidoGetProjectPipelineResponse kafeido get project pipeline response
//
// swagger:model kafeidoGetProjectPipelineResponse
type KafeidoGetProjectPipelineResponse struct {

	// kf pipeline Id
	KfPipelineID string `json:"kfPipelineId,omitempty"`

	// kf pipeline version Id
	KfPipelineVersionID string `json:"kfPipelineVersionId,omitempty"`

	// kf pipeline workflow
	// Format: byte
	KfPipelineWorkflow strfmt.Base64 `json:"kfPipelineWorkflow,omitempty"`

	// named pipeline Id
	NamedPipelineID string `json:"namedPipelineId,omitempty"`

	// named pipeline name
	NamedPipelineName string `json:"namedPipelineName,omitempty"`

	// named pipeline type
	NamedPipelineType string `json:"namedPipelineType,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// run parameters
	RunParameters map[string]string `json:"runParameters,omitempty"`
}

// Validate validates this kafeido get project pipeline response
func (m *KafeidoGetProjectPipelineResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kafeido get project pipeline response based on context it is used
func (m *KafeidoGetProjectPipelineResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoGetProjectPipelineResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoGetProjectPipelineResponse) UnmarshalBinary(b []byte) error {
	var res KafeidoGetProjectPipelineResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
