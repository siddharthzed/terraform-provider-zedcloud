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

// DocPolicy DocPolicy detail
//
// DocPolicy meta data
//
// swagger:model DocPolicy
type DocPolicy struct {

	// Policy doc fileURL
	FileURL string `json:"fileURL,omitempty"`

	// Unique system defined docpolicy ID
	// Read Only: true
	// Pattern: [0-9A-Za-z_=-]{28}
	ID string `json:"id,omitempty"`

	// Mark latest docpolicy check
	Latest bool `json:"latest,omitempty"`

	// User defined name of the docpolicy. Name cannot be changed once created
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	PolicyName *string `json:"policyName"`

	// System defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// Server Host
	ServerHost string `json:"serverHost,omitempty"`

	// Policy doc version
	Version string `json:"version,omitempty"`
}

// Validate validates this doc policy
func (m *DocPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocPolicy) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *DocPolicy) validatePolicyName(formats strfmt.Registry) error {

	if err := validate.Required("policyName", "body", m.PolicyName); err != nil {
		return err
	}

	if err := validate.MinLength("policyName", "body", *m.PolicyName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("policyName", "body", *m.PolicyName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("policyName", "body", *m.PolicyName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *DocPolicy) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this doc policy based on the context it is used
func (m *DocPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocPolicy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DocPolicy) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocPolicy) UnmarshalBinary(b []byte) error {
	var res DocPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
