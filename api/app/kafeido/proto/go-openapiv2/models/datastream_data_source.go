// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DatastreamDataSource datastream data source
//
// swagger:model datastreamDataSource
type DatastreamDataSource string

func NewDatastreamDataSource(value DatastreamDataSource) *DatastreamDataSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DatastreamDataSource.
func (m DatastreamDataSource) Pointer() *DatastreamDataSource {
	return &m
}

const (

	// DatastreamDataSourceDATASOURCENONE captures enum value "DATASOURCE_NONE"
	DatastreamDataSourceDATASOURCENONE DatastreamDataSource = "DATASOURCE_NONE"

	// DatastreamDataSourceDATASOURCEPHOTO captures enum value "DATASOURCE_PHOTO"
	DatastreamDataSourceDATASOURCEPHOTO DatastreamDataSource = "DATASOURCE_PHOTO"

	// DatastreamDataSourceDATASOURCEVIDEO captures enum value "DATASOURCE_VIDEO"
	DatastreamDataSourceDATASOURCEVIDEO DatastreamDataSource = "DATASOURCE_VIDEO"

	// DatastreamDataSourceDATASOURCESTREAMING captures enum value "DATASOURCE_STREAMING"
	DatastreamDataSourceDATASOURCESTREAMING DatastreamDataSource = "DATASOURCE_STREAMING"

	// DatastreamDataSourceDATASOURCEYOUTUBE captures enum value "DATASOURCE_YOUTUBE"
	DatastreamDataSourceDATASOURCEYOUTUBE DatastreamDataSource = "DATASOURCE_YOUTUBE"

	// DatastreamDataSourceDATASOURCEIMAGEURL captures enum value "DATASOURCE_IMAGEURL"
	DatastreamDataSourceDATASOURCEIMAGEURL DatastreamDataSource = "DATASOURCE_IMAGEURL"
)

// for schema
var datastreamDataSourceEnum []interface{}

func init() {
	var res []DatastreamDataSource
	if err := json.Unmarshal([]byte(`["DATASOURCE_NONE","DATASOURCE_PHOTO","DATASOURCE_VIDEO","DATASOURCE_STREAMING","DATASOURCE_YOUTUBE","DATASOURCE_IMAGEURL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datastreamDataSourceEnum = append(datastreamDataSourceEnum, v)
	}
}

func (m DatastreamDataSource) validateDatastreamDataSourceEnum(path, location string, value DatastreamDataSource) error {
	if err := validate.EnumCase(path, location, value, datastreamDataSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datastream data source
func (m DatastreamDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatastreamDataSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this datastream data source based on context it is used
func (m DatastreamDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}