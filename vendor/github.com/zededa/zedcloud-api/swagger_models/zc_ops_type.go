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

// ZcOpsType ZedCloud internal operation type
//
// swagger:model ZcOpsType
type ZcOpsType string

func NewZcOpsType(value ZcOpsType) *ZcOpsType {
	v := value
	return &v
}

const (

	// ZcOpsTypeOPSTYPEUNSPECIFIED captures enum value "OPS_TYPE_UNSPECIFIED"
	ZcOpsTypeOPSTYPEUNSPECIFIED ZcOpsType = "OPS_TYPE_UNSPECIFIED"

	// ZcOpsTypeOPSTYPEREAD captures enum value "OPS_TYPE_READ"
	ZcOpsTypeOPSTYPEREAD ZcOpsType = "OPS_TYPE_READ"

	// ZcOpsTypeOPSTYPEDELETE captures enum value "OPS_TYPE_DELETE"
	ZcOpsTypeOPSTYPEDELETE ZcOpsType = "OPS_TYPE_DELETE"

	// ZcOpsTypeOPSTYPECREATE captures enum value "OPS_TYPE_CREATE"
	ZcOpsTypeOPSTYPECREATE ZcOpsType = "OPS_TYPE_CREATE"

	// ZcOpsTypeOPSTYPEUPDATE captures enum value "OPS_TYPE_UPDATE"
	ZcOpsTypeOPSTYPEUPDATE ZcOpsType = "OPS_TYPE_UPDATE"

	// ZcOpsTypeOPSTYPELIST captures enum value "OPS_TYPE_LIST"
	ZcOpsTypeOPSTYPELIST ZcOpsType = "OPS_TYPE_LIST"
)

// for schema
var zcOpsTypeEnum []interface{}

func init() {
	var res []ZcOpsType
	if err := json.Unmarshal([]byte(`["OPS_TYPE_UNSPECIFIED","OPS_TYPE_READ","OPS_TYPE_DELETE","OPS_TYPE_CREATE","OPS_TYPE_UPDATE","OPS_TYPE_LIST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zcOpsTypeEnum = append(zcOpsTypeEnum, v)
	}
}

func (m ZcOpsType) validateZcOpsTypeEnum(path, location string, value ZcOpsType) error {
	if err := validate.EnumCase(path, location, value, zcOpsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this zc ops type
func (m ZcOpsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateZcOpsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this zc ops type based on context it is used
func (m ZcOpsType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
