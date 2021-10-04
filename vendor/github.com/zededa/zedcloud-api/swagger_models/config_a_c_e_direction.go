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

// ConfigACEDirection config a c e direction
//
// swagger:model configACEDirection
type ConfigACEDirection string

func NewConfigACEDirection(value ConfigACEDirection) *ConfigACEDirection {
	v := value
	return &v
}

const (

	// ConfigACEDirectionBOTH captures enum value "BOTH"
	ConfigACEDirectionBOTH ConfigACEDirection = "BOTH"

	// ConfigACEDirectionINGRESS captures enum value "INGRESS"
	ConfigACEDirectionINGRESS ConfigACEDirection = "INGRESS"

	// ConfigACEDirectionEGRESS captures enum value "EGRESS"
	ConfigACEDirectionEGRESS ConfigACEDirection = "EGRESS"
)

// for schema
var configACEDirectionEnum []interface{}

func init() {
	var res []ConfigACEDirection
	if err := json.Unmarshal([]byte(`["BOTH","INGRESS","EGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configACEDirectionEnum = append(configACEDirectionEnum, v)
	}
}

func (m ConfigACEDirection) validateConfigACEDirectionEnum(path, location string, value ConfigACEDirection) error {
	if err := validate.EnumCase(path, location, value, configACEDirectionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config a c e direction
func (m ConfigACEDirection) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigACEDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config a c e direction based on context it is used
func (m ConfigACEDirection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
