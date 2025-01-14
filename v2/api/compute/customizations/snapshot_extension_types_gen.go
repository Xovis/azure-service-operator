// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20200930"
	"github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20200930storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SnapshotExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SnapshotExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&compute.Snapshot{},
		&v1alpha1api20200930storage.Snapshot{}}
}
