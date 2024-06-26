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

// KafeidoDataSourceInfo kafeido data source info
//
// swagger:model kafeidoDataSourceInfo
type KafeidoDataSourceInfo struct {

	// audio files data source
	AudioFilesDataSource *DatastreamAudioFilesDataSource `json:"audioFilesDataSource,omitempty"`

	// delay in duration for next frame
	DelayInDurationForNextFrame string `json:"delayInDurationForNextFrame,omitempty"`

	// image Url data source
	ImageURLDataSource *DatastreamImageURLDataSource `json:"imageUrlDataSource,omitempty"`

	// photo data source
	PhotoDataSource *DatastreamPhotoDataSource `json:"photoDataSource,omitempty"`

	// streaming data source
	StreamingDataSource *DatastreamStreamingDataSource `json:"streamingDataSource,omitempty"`

	// type
	Type *DatastreamDataSource `json:"type,omitempty"`

	// video data source
	VideoDataSource *DatastreamVideoDataSource `json:"videoDataSource,omitempty"`

	// youtube data source
	YoutubeDataSource *DatastreamYoutubeDataSource `json:"youtubeDataSource,omitempty"`
}

// Validate validates this kafeido data source info
func (m *KafeidoDataSourceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioFilesDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageURLDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreamingDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVideoDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYoutubeDataSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoDataSourceInfo) validateAudioFilesDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.AudioFilesDataSource) { // not required
		return nil
	}

	if m.AudioFilesDataSource != nil {
		if err := m.AudioFilesDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioFilesDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioFilesDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validateImageURLDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageURLDataSource) { // not required
		return nil
	}

	if m.ImageURLDataSource != nil {
		if err := m.ImageURLDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageUrlDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageUrlDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validatePhotoDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.PhotoDataSource) { // not required
		return nil
	}

	if m.PhotoDataSource != nil {
		if err := m.PhotoDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photoDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("photoDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validateStreamingDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.StreamingDataSource) { // not required
		return nil
	}

	if m.StreamingDataSource != nil {
		if err := m.StreamingDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamingDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamingDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validateVideoDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.VideoDataSource) { // not required
		return nil
	}

	if m.VideoDataSource != nil {
		if err := m.VideoDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("videoDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("videoDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) validateYoutubeDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.YoutubeDataSource) { // not required
		return nil
	}

	if m.YoutubeDataSource != nil {
		if err := m.YoutubeDataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("youtubeDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("youtubeDataSource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kafeido data source info based on the context it is used
func (m *KafeidoDataSourceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudioFilesDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageURLDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhotoDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreamingDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVideoDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYoutubeDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateAudioFilesDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.AudioFilesDataSource != nil {

		if swag.IsZero(m.AudioFilesDataSource) { // not required
			return nil
		}

		if err := m.AudioFilesDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioFilesDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioFilesDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateImageURLDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageURLDataSource != nil {

		if swag.IsZero(m.ImageURLDataSource) { // not required
			return nil
		}

		if err := m.ImageURLDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageUrlDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageUrlDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidatePhotoDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.PhotoDataSource != nil {

		if swag.IsZero(m.PhotoDataSource) { // not required
			return nil
		}

		if err := m.PhotoDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photoDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("photoDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateStreamingDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.StreamingDataSource != nil {

		if swag.IsZero(m.StreamingDataSource) { // not required
			return nil
		}

		if err := m.StreamingDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamingDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamingDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateVideoDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.VideoDataSource != nil {

		if swag.IsZero(m.VideoDataSource) { // not required
			return nil
		}

		if err := m.VideoDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("videoDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("videoDataSource")
			}
			return err
		}
	}

	return nil
}

func (m *KafeidoDataSourceInfo) contextValidateYoutubeDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.YoutubeDataSource != nil {

		if swag.IsZero(m.YoutubeDataSource) { // not required
			return nil
		}

		if err := m.YoutubeDataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("youtubeDataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("youtubeDataSource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KafeidoDataSourceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KafeidoDataSourceInfo) UnmarshalBinary(b []byte) error {
	var res KafeidoDataSourceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
