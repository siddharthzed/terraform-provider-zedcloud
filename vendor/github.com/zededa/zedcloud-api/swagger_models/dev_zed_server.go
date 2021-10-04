// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DevZedServer DevZedServer payload detail
//
// DevZedServer request paylod
//
// swagger:model DevZedServer
type DevZedServer struct {

	// EID
	// Required: true
	EID []string `json:"EID"`

	// Hostname for dev zed server
	// Required: true
	HostName *string `json:"hostName"`
}

// Validate validates this dev zed server
func (m *DevZedServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevZedServer) validateEID(formats strfmt.Registry) error {

	if err := validate.Required("EID", "body", m.EID); err != nil {
		return err
	}

	return nil
}

func (m *DevZedServer) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dev zed server based on context it is used
func (m *DevZedServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DevZedServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevZedServer) UnmarshalBinary(b []byte) error {
	var res DevZedServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
