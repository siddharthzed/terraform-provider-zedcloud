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
)

// EnrollmentDetail enrollment detail
//
// swagger:model EnrollmentDetail
type EnrollmentDetail struct {

	// allocation policy
	AllocationPolicy *AllocationPolicy `json:"allocationPolicy,omitempty"`

	// attached iot hubs name
	AttachedIotHubsName []string `json:"attachedIotHubsName"`

	// certificate enrollment
	CertificateEnrollment CertificateEnrollmentDetail `json:"certificateEnrollment,omitempty"`

	// enable iot edge device
	EnableIotEdgeDevice bool `json:"enableIotEdgeDevice,omitempty"`

	// mechanism
	Mechanism *EnrollmentMechanism `json:"mechanism,omitempty"`

	// symmetric key enrollment
	SymmetricKeyEnrollment *SymmetricKeyEnrollmentDetail `json:"symmetricKeyEnrollment,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// tpm enrollment
	TpmEnrollment *TPMEnrollmentDetail `json:"tpmEnrollment,omitempty"`
}

// Validate validates this enrollment detail
func (m *EnrollmentDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMechanism(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymmetricKeyEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTpmEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrollmentDetail) validateAllocationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AllocationPolicy) { // not required
		return nil
	}

	if m.AllocationPolicy != nil {
		if err := m.AllocationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) validateMechanism(formats strfmt.Registry) error {
	if swag.IsZero(m.Mechanism) { // not required
		return nil
	}

	if m.Mechanism != nil {
		if err := m.Mechanism.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mechanism")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) validateSymmetricKeyEnrollment(formats strfmt.Registry) error {
	if swag.IsZero(m.SymmetricKeyEnrollment) { // not required
		return nil
	}

	if m.SymmetricKeyEnrollment != nil {
		if err := m.SymmetricKeyEnrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("symmetricKeyEnrollment")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) validateTpmEnrollment(formats strfmt.Registry) error {
	if swag.IsZero(m.TpmEnrollment) { // not required
		return nil
	}

	if m.TpmEnrollment != nil {
		if err := m.TpmEnrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tpmEnrollment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enrollment detail based on the context it is used
func (m *EnrollmentDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMechanism(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSymmetricKeyEnrollment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTpmEnrollment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnrollmentDetail) contextValidateAllocationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AllocationPolicy != nil {
		if err := m.AllocationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) contextValidateMechanism(ctx context.Context, formats strfmt.Registry) error {

	if m.Mechanism != nil {
		if err := m.Mechanism.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mechanism")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) contextValidateSymmetricKeyEnrollment(ctx context.Context, formats strfmt.Registry) error {

	if m.SymmetricKeyEnrollment != nil {
		if err := m.SymmetricKeyEnrollment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("symmetricKeyEnrollment")
			}
			return err
		}
	}

	return nil
}

func (m *EnrollmentDetail) contextValidateTpmEnrollment(ctx context.Context, formats strfmt.Registry) error {

	if m.TpmEnrollment != nil {
		if err := m.TpmEnrollment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tpmEnrollment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnrollmentDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnrollmentDetail) UnmarshalBinary(b []byte) error {
	var res EnrollmentDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
