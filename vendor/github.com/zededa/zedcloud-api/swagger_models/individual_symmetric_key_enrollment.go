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

// IndividualSymmetricKeyEnrollment individual symmetric key enrollment
//
// swagger:model IndividualSymmetricKeyEnrollment
type IndividualSymmetricKeyEnrollment struct {

	// registration Id
	RegistrationID string `json:"registrationId,omitempty"`
}

// Validate validates this individual symmetric key enrollment
func (m *IndividualSymmetricKeyEnrollment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this individual symmetric key enrollment based on context it is used
func (m *IndividualSymmetricKeyEnrollment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IndividualSymmetricKeyEnrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndividualSymmetricKeyEnrollment) UnmarshalBinary(b []byte) error {
	var res IndividualSymmetricKeyEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
