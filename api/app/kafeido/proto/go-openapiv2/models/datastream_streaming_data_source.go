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

// DatastreamStreamingDataSource datastream streaming data source
//
// swagger:model datastreamStreamingDataSource
type DatastreamStreamingDataSource struct {

	// channel name
	ChannelName string `json:"channelName,omitempty"`

	// frames per second
	FramesPerSecond float32 `json:"framesPerSecond,omitempty"`

	// streaming info
	StreamingInfo *DatastreamStreamingInfo `json:"streamingInfo,omitempty"`
}

// Validate validates this datastream streaming data source
func (m *DatastreamStreamingDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStreamingInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastreamStreamingDataSource) validateStreamingInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.StreamingInfo) { // not required
		return nil
	}

	if m.StreamingInfo != nil {
		if err := m.StreamingInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamingInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datastream streaming data source based on the context it is used
func (m *DatastreamStreamingDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStreamingInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastreamStreamingDataSource) contextValidateStreamingInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.StreamingInfo != nil {

		if swag.IsZero(m.StreamingInfo) { // not required
			return nil
		}

		if err := m.StreamingInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamingInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamingInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatastreamStreamingDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastreamStreamingDataSource) UnmarshalBinary(b []byte) error {
	var res DatastreamStreamingDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
