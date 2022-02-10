// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RoleAssignments_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Role assignment properties.
	Properties RoleAssignmentProperties_SpecARM `json:"properties"`
}

var _ genruntime.ARMResourceSpec = &RoleAssignments_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01"
func (specarm RoleAssignments_SPECARM) GetAPIVersion() string {
	return "2020-08-01"
}

// GetName returns the Name of the resource
func (specarm RoleAssignments_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm RoleAssignments_SPECARM) GetType() string {
	return ""
}

type RoleAssignmentProperties_SpecARM struct {
	//Condition: The conditions on the role assignment. This limits the resources it
	//can be assigned to. e.g.:
	//@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName]
	//StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	//ConditionVersion: Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	//DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	//Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	//PrincipalId: The principal ID.
	PrincipalId string `json:"principalId"`

	//PrincipalType: The principal type of the assigned principal ID.
	PrincipalType *RoleAssignmentPropertiesSpecPrincipalType `json:"principalType,omitempty"`

	//RoleDefinitionId: The role definition ID.
	RoleDefinitionId string `json:"roleDefinitionId"`
}