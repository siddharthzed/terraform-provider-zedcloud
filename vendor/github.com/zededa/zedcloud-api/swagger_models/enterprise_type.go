// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EnterpriseType enterprise type
//
// swagger:model EnterpriseType
type EnterpriseType string

func NewEnterpriseType(value EnterpriseType) *EnterpriseType {
	v := value
	return &v
}

const (

	// EnterpriseTypeENTERPRISETYPEUNSPECIFIED captures enum value "ENTERPRISE_TYPE_UNSPECIFIED"
	EnterpriseTypeENTERPRISETYPEUNSPECIFIED EnterpriseType = "ENTERPRISE_TYPE_UNSPECIFIED"

	// EnterpriseTypeENTERPRISETYPESELFSIGNUP captures enum value "ENTERPRISE_TYPE_SELFSIGNUP"
	EnterpriseTypeENTERPRISETYPESELFSIGNUP EnterpriseType = "ENTERPRISE_TYPE_SELFSIGNUP"
)

// for schema
var enterpriseTypeEnum []interface{}

func init() {
	var res []EnterpriseType
	if err := json.Unmarshal([]byte(`["ENTERPRISE_TYPE_UNSPECIFIED","ENTERPRISE_TYPE_SELFSIGNUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enterpriseTypeEnum = append(enterpriseTypeEnum, v)
	}
}

func (m EnterpriseType) validateEnterpriseTypeEnum(path, location string, value EnterpriseType) error {
	if err := validate.EnumCase(path, location, value, enterpriseTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this enterprise type
func (m EnterpriseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnterpriseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enterprise type based on context it is used
func (m EnterpriseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
