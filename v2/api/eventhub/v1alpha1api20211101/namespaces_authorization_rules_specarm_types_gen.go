// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesAuthorizationRules_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Properties supplied to create or update AuthorizationRule
	Properties *NamespacesAuthorizationRules_SPEC_PropertiesARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesAuthorizationRules_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (specarm NamespacesAuthorizationRules_SPECARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (specarm NamespacesAuthorizationRules_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm NamespacesAuthorizationRules_SPECARM) GetType() string {
	return ""
}

type NamespacesAuthorizationRules_SPEC_PropertiesARM struct {
	//Rights: The rights associated with the rule.
	Rights []NamespacesAuthorizationRulesSPECPropertiesRights `json:"rights"`
}