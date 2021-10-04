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

// IotHubServiceDetail iot hub service detail
//
// swagger:model IotHubServiceDetail
type IotHubServiceDetail struct {

	// service detail
	ServiceDetail *AzureResourceAndServiceDetail `json:"serviceDetail,omitempty"`
}

// Validate validates this iot hub service detail
func (m *IotHubServiceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IotHubServiceDetail) validateServiceDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceDetail) { // not required
		return nil
	}

	if m.ServiceDetail != nil {
		if err := m.ServiceDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iot hub service detail based on the context it is used
func (m *IotHubServiceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IotHubServiceDetail) contextValidateServiceDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceDetail != nil {
		if err := m.ServiceDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IotHubServiceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IotHubServiceDetail) UnmarshalBinary(b []byte) error {
	var res IotHubServiceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
