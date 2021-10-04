// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EDConfigItem EDConfigItem payload detail
//
// ED Configuration Item request paylod
//
// swagger:model EDConfigItem
type EDConfigItem struct {

	// boolean value
	BoolValue bool `json:"boolValue,omitempty"`

	// float value
	FloatValue float32 `json:"floatValue,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// string value
	StringValue string `json:"stringValue,omitempty"`

	// uint32 value
	Uint32Value int64 `json:"uint32Value,omitempty"`

	// uint64 value in string format
	Uint64Value string `json:"uint64Value,omitempty"`

	// value type
	ValueType string `json:"valueType,omitempty"`
}

// Validate validates this e d config item
func (m *EDConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e d config item based on context it is used
func (m *EDConfigItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EDConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EDConfigItem) UnmarshalBinary(b []byte) error {
	var res EDConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
