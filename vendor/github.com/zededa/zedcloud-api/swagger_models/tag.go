// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tag Resource group detail
//
// Resource group for edge gateway Base OS or for eedge applications.
// Example: {"attr":{},"description":"My test resource group for Edge computing","id":"d1125b0f-633d-459c-99c6-f47e7467cebc","name":"my-test-project","numdevices":208,"revision":{"createdAt":{"seconds":1592068270},"createdBy":"admin@my-company.com","curr":"1","updatedAt":{"seconds":1592068271},"updatedBy":"admin@my-company.com"},"title":"My Test resource group","type":"TAG_TYPE_PROJECT"}
//
// swagger:model Tag
type Tag struct {

	// Resource group wide policy for edge applications to be deployed on all edge nodes on this resource group
	// Read Only: true
	AppPolicy *PolicyConfig `json:"appPolicy,omitempty"`

	// Attestation policy to enforce on all devices of this project
	AttestationPolicy *PolicyConfig `json:"attestationPolicy,omitempty"`

	// attr
	Attr map[string]string `json:"attr,omitempty"`

	// Resource group wide policy for Azure IoTEdge configuration to be applied to all edge applications
	// Read Only: true
	CloudPolicy *PolicyConfig `json:"cloudPolicy,omitempty"`

	// Detailed description of the resource group.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// System defined universally unique Id of the resource group.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// Resource group wide policy for Azure module configuration to be applied to all edge applications
	// Read Only: true
	ModulePolicy []*PolicyConfig `json:"modulePolicy"`

	// User defined name of the resource group, unique across the enterprise. Once resource group is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// Network policy to enforce on all devices of this project
	NetworkPolicy *PolicyConfig `json:"networkPolicy,omitempty"`

	// Number of edge nodes in this resource group
	// Read Only: true
	Numdevices int64 `json:"numdevices,omitempty"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// User defined title of the resource group. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// Resource group type
	// Required: true
	Type *TagType `json:"type"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttestationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModulePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateAppPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AppPolicy) { // not required
		return nil
	}

	if m.AppPolicy != nil {
		if err := m.AppPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateAttestationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AttestationPolicy) { // not required
		return nil
	}

	if m.AttestationPolicy != nil {
		if err := m.AttestationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateCloudPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudPolicy) { // not required
		return nil
	}

	if m.CloudPolicy != nil {
		if err := m.CloudPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateModulePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModulePolicy) { // not required
		return nil
	}

	for i := 0; i < len(m.ModulePolicy); i++ {
		if swag.IsZero(m.ModulePolicy[i]) { // not required
			continue
		}

		if m.ModulePolicy[i] != nil {
			if err := m.ModulePolicy[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modulePolicy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tag) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateNetworkPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPolicy) { // not required
		return nil
	}

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) validateRevision(formats strfmt.Registry) error {
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

func (m *Tag) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tag based on the context it is used
func (m *Tag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttestationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModulePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumdevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) contextValidateAppPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AppPolicy != nil {
		if err := m.AppPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) contextValidateAttestationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AttestationPolicy != nil {
		if err := m.AttestationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) contextValidateCloudPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudPolicy != nil {
		if err := m.CloudPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Tag) contextValidateModulePolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modulePolicy", "body", []*PolicyConfig(m.ModulePolicy)); err != nil {
		return err
	}

	for i := 0; i < len(m.ModulePolicy); i++ {

		if m.ModulePolicy[i] != nil {
			if err := m.ModulePolicy[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modulePolicy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tag) contextValidateNetworkPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Tag) contextValidateNumdevices(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numdevices", "body", int64(m.Numdevices)); err != nil {
		return err
	}

	return nil
}

func (m *Tag) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Tag) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
