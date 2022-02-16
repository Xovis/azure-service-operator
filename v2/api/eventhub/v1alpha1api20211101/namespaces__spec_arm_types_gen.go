// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Namespaces_SpecARM struct {
	//Identity: Properties to configure Identity for Bring your Own Keys
	Identity *IdentityARM `json:"identity,omitempty"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Namespace properties supplied for create namespace operation.
	Properties Namespaces_Spec_PropertiesARM `json:"properties"`

	//Sku: SKU parameters supplied to the create namespace operation
	Sku *SkuARM `json:"sku,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespaces Namespaces_SpecARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (namespaces Namespaces_SpecARM) GetName() string {
	return namespaces.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces"
func (namespaces Namespaces_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces"
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity
type IdentityARM struct {
	//Type: Type of managed service identity.
	Type *IdentityType `json:"type,omitempty"`
}

type Namespaces_Spec_PropertiesARM struct {
	//AlternateName: Alternate name specified when alias and namespace names are same.
	AlternateName *string `json:"alternateName,omitempty"`
	ClusterArmId  *string `json:"clusterArmId,omitempty"`

	//DisableLocalAuth: This property disables SAS authentication for the Event Hubs
	//namespace.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	//Encryption: Properties to configure Encryption
	Encryption *EncryptionARM `json:"encryption,omitempty"`

	//IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for
	//eventhub namespace.
	IsAutoInflateEnabled *bool `json:"isAutoInflateEnabled,omitempty"`

	//KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub
	//namespace.
	KafkaEnabled *bool `json:"kafkaEnabled,omitempty"`

	//MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is
	//enabled, value should be within 0 to 20 throughput units. ( '0' if
	//AutoInflateEnabled = true)
	MaximumThroughputUnits *int `json:"maximumThroughputUnits,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []Namespaces_Spec_Properties_PrivateEndpointConnectionsARM `json:"privateEndpointConnections,omitempty"`

	//ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in
	//regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku
type SkuARM struct {
	//Capacity: The Event Hubs throughput units for Basic or Standard tiers, where
	//value should be 0 to 20 throughput units. The Event Hubs premium units for
	//Premium tier, where value should be 0 to 10 premium units.
	Capacity *int `json:"capacity,omitempty"`

	//Name: Name of this SKU.
	Name SkuName `json:"name"`

	//Tier: The billing tier of this particular SKU.
	Tier *SkuTier `json:"tier,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption
type EncryptionARM struct {
	//KeySource: Enumerates the possible value of keySource for Encryption.
	KeySource *EncryptionKeySource `json:"keySource,omitempty"`

	//KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultPropertiesARM `json:"keyVaultProperties,omitempty"`

	//RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double
	//Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type IdentityType string

const (
	IdentityTypeNone                       = IdentityType("None")
	IdentityTypeSystemAssigned             = IdentityType("SystemAssigned")
	IdentityTypeSystemAssignedUserAssigned = IdentityType("SystemAssigned, UserAssigned")
	IdentityTypeUserAssigned               = IdentityType("UserAssigned")
)

type Namespaces_Spec_Properties_PrivateEndpointConnectionsARM struct {
	//Properties: Properties of the private endpoint connection resource.
	Properties *PrivateEndpointConnectionPropertiesARM `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNamePremium  = SkuName("Premium")
	SkuNameStandard = SkuName("Standard")
)

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type SkuTier string

const (
	SkuTierBasic    = SkuTier("Basic")
	SkuTierPremium  = SkuTier("Premium")
	SkuTierStandard = SkuTier("Standard")
)

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties
type KeyVaultPropertiesARM struct {
	Identity *UserAssignedIdentityPropertiesARM `json:"identity,omitempty"`

	//KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	//KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	//KeyVersion: Key Version
	KeyVersion *string `json:"keyVersion,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpointConnectionProperties
type PrivateEndpointConnectionPropertiesARM struct {
	//PrivateEndpoint: PrivateEndpoint information.
	PrivateEndpoint *PrivateEndpointARM `json:"privateEndpoint,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint
type PrivateEndpointARM struct {
	Id *string `json:"id,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties
type UserAssignedIdentityPropertiesARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}