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

// ConfigInstanceOpsCmd config instance ops cmd
//
// swagger:model configInstanceOpsCmd
type ConfigInstanceOpsCmd struct {

	// counter
	Counter int64 `json:"counter,omitempty"`

	// ops time
	OpsTime string `json:"opsTime,omitempty"`
}

// Validate validates this config instance ops cmd
func (m *ConfigInstanceOpsCmd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config instance ops cmd based on context it is used
func (m *ConfigInstanceOpsCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigInstanceOpsCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigInstanceOpsCmd) UnmarshalBinary(b []byte) error {
	var res ConfigInstanceOpsCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
