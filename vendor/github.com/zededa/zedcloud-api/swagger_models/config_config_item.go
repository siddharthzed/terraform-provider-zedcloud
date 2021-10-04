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

// ConfigConfigItem Timers and other per-device policy which relates to the interaction
// with zedcloud. Note that the timers are randomized on the device
// to avoid synchronization with other devices. Random range is between
// between .5 and 1.5 of these nominal values. If not set (i.e. zero),
// it means the default value of 60 seconds.
//
// swagger:model configConfigItem
type ConfigConfigItem struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this config config item
func (m *ConfigConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config config item based on context it is used
func (m *ConfigConfigItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigConfigItem) UnmarshalBinary(b []byte) error {
	var res ConfigConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
