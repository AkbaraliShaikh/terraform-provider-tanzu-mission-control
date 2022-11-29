/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
Code generated by go-swagger; DO NOT EDIT.
*/

package policyrecipeimagemodel

import (
	"github.com/go-openapi/swag"

	policyrecipeimagecommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/image/common"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag is model for allowed-name-tag recipe version v1
//
// The input schema for image policy allowed-name-tag recipe.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag
type VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag struct {

	// Audit (dry-run)
	// Creates this policy for dry-run. Violations will be logged but not denied. Defaults to false (deny).
	Audit *bool `json:"audit,omitempty"`

	// This specifies a list of rules that defines allowed image patterns.
	// Required: true
	// Min Items: 1
	Rules []*VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules `json:"rules"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules Rules.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules
type VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules struct {

	// Allowed image names, wildcards are supported(for example: fooservice/*). Empty field is equivalent to *.
	ImageName string `json:"imageName,omitempty"`

	// Tag
	Tag *policyrecipeimagecommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecImageV1RulesTag `json:"tag,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecImageV1AllowedNameTagRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
