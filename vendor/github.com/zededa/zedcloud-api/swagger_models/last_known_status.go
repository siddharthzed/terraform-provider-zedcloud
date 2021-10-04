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

// LastKnownStatus last known status
//
// swagger:model LastKnownStatus
type LastKnownStatus struct {

	// last attempt state
	LastAttemptState int32 `json:"lastAttemptState,omitempty"`

	// last attempted at
	// Format: date-time
	LastAttemptedAt strfmt.DateTime `json:"lastAttemptedAt,omitempty"`

	// last known status line
	LastKnownStatusLine string `json:"lastKnownStatusLine,omitempty"`
}

// Validate validates this last known status
func (m *LastKnownStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastAttemptedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LastKnownStatus) validateLastAttemptedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastAttemptedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastAttemptedAt", "body", "date-time", m.LastAttemptedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this last known status based on context it is used
func (m *LastKnownStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LastKnownStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LastKnownStatus) UnmarshalBinary(b []byte) error {
	var res LastKnownStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
