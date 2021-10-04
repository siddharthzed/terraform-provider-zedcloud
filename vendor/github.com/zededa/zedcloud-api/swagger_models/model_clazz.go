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

// ModelClazz Global identifier of clazz types
//
// swagger:model ModelClazz
type ModelClazz string

func NewModelClazz(value ModelClazz) *ModelClazz {
	v := value
	return &v
}

const (

	// ModelClazzModelClazzUnspecified captures enum value "ModelClazzUnspecified"
	ModelClazzModelClazzUnspecified ModelClazz = "ModelClazzUnspecified"

	// ModelClazzModelClazzDetailedUser captures enum value "ModelClazzDetailedUser"
	ModelClazzModelClazzDetailedUser ModelClazz = "ModelClazzDetailedUser"

	// ModelClazzModelClazzSimpleUser captures enum value "ModelClazzSimpleUser"
	ModelClazzModelClazzSimpleUser ModelClazz = "ModelClazzSimpleUser"

	// ModelClazzModelClazzEnterprise captures enum value "ModelClazzEnterprise"
	ModelClazzModelClazzEnterprise ModelClazz = "ModelClazzEnterprise"

	// ModelClazzModelClazzRealm captures enum value "ModelClazzRealm"
	ModelClazzModelClazzRealm ModelClazz = "ModelClazzRealm"

	// ModelClazzModelClazzCredential captures enum value "ModelClazzCredential"
	ModelClazzModelClazzCredential ModelClazz = "ModelClazzCredential"

	// ModelClazzModelClazzPolicy captures enum value "ModelClazzPolicy"
	ModelClazzModelClazzPolicy ModelClazz = "ModelClazzPolicy"

	// ModelClazzModelClazzRole captures enum value "ModelClazzRole"
	ModelClazzModelClazzRole ModelClazz = "ModelClazzRole"

	// ModelClazzModelClazzProfile captures enum value "ModelClazzProfile"
	ModelClazzModelClazzProfile ModelClazz = "ModelClazzProfile"

	// ModelClazzModelClazzDocPolicy captures enum value "ModelClazzDocPolicy"
	ModelClazzModelClazzDocPolicy ModelClazz = "ModelClazzDocPolicy"
)

// for schema
var modelClazzEnum []interface{}

func init() {
	var res []ModelClazz
	if err := json.Unmarshal([]byte(`["ModelClazzUnspecified","ModelClazzDetailedUser","ModelClazzSimpleUser","ModelClazzEnterprise","ModelClazzRealm","ModelClazzCredential","ModelClazzPolicy","ModelClazzRole","ModelClazzProfile","ModelClazzDocPolicy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelClazzEnum = append(modelClazzEnum, v)
	}
}

func (m ModelClazz) validateModelClazzEnum(path, location string, value ModelClazz) error {
	if err := validate.EnumCase(path, location, value, modelClazzEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model clazz
func (m ModelClazz) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelClazzEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model clazz based on context it is used
func (m ModelClazz) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
