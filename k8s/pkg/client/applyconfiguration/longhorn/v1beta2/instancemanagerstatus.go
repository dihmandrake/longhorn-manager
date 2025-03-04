/*
Copyright The Longhorn Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
)

// InstanceManagerStatusApplyConfiguration represents a declarative configuration of the InstanceManagerStatus type for use
// with apply.
type InstanceManagerStatusApplyConfiguration struct {
	OwnerID            *string                                             `json:"ownerID,omitempty"`
	CurrentState       *longhornv1beta2.InstanceManagerState               `json:"currentState,omitempty"`
	InstanceEngines    map[string]InstanceProcessApplyConfiguration        `json:"instanceEngines,omitempty"`
	InstanceReplicas   map[string]InstanceProcessApplyConfiguration        `json:"instanceReplicas,omitempty"`
	BackingImages      map[string]BackingImageV2CopyInfoApplyConfiguration `json:"backingImages,omitempty"`
	IP                 *string                                             `json:"ip,omitempty"`
	APIMinVersion      *int                                                `json:"apiMinVersion,omitempty"`
	APIVersion         *int                                                `json:"apiVersion,omitempty"`
	ProxyAPIMinVersion *int                                                `json:"proxyApiMinVersion,omitempty"`
	ProxyAPIVersion    *int                                                `json:"proxyApiVersion,omitempty"`
	DataEngineStatus   *DataEngineStatusApplyConfiguration                 `json:"dataEngineStatus,omitempty"`
	Instances          map[string]InstanceProcessApplyConfiguration        `json:"instances,omitempty"`
}

// InstanceManagerStatusApplyConfiguration constructs a declarative configuration of the InstanceManagerStatus type for use with
// apply.
func InstanceManagerStatus() *InstanceManagerStatusApplyConfiguration {
	return &InstanceManagerStatusApplyConfiguration{}
}

// WithOwnerID sets the OwnerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OwnerID field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithOwnerID(value string) *InstanceManagerStatusApplyConfiguration {
	b.OwnerID = &value
	return b
}

// WithCurrentState sets the CurrentState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentState field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithCurrentState(value longhornv1beta2.InstanceManagerState) *InstanceManagerStatusApplyConfiguration {
	b.CurrentState = &value
	return b
}

// WithInstanceEngines puts the entries into the InstanceEngines field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the InstanceEngines field,
// overwriting an existing map entries in InstanceEngines field with the same key.
func (b *InstanceManagerStatusApplyConfiguration) WithInstanceEngines(entries map[string]InstanceProcessApplyConfiguration) *InstanceManagerStatusApplyConfiguration {
	if b.InstanceEngines == nil && len(entries) > 0 {
		b.InstanceEngines = make(map[string]InstanceProcessApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.InstanceEngines[k] = v
	}
	return b
}

// WithInstanceReplicas puts the entries into the InstanceReplicas field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the InstanceReplicas field,
// overwriting an existing map entries in InstanceReplicas field with the same key.
func (b *InstanceManagerStatusApplyConfiguration) WithInstanceReplicas(entries map[string]InstanceProcessApplyConfiguration) *InstanceManagerStatusApplyConfiguration {
	if b.InstanceReplicas == nil && len(entries) > 0 {
		b.InstanceReplicas = make(map[string]InstanceProcessApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.InstanceReplicas[k] = v
	}
	return b
}

// WithBackingImages puts the entries into the BackingImages field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the BackingImages field,
// overwriting an existing map entries in BackingImages field with the same key.
func (b *InstanceManagerStatusApplyConfiguration) WithBackingImages(entries map[string]BackingImageV2CopyInfoApplyConfiguration) *InstanceManagerStatusApplyConfiguration {
	if b.BackingImages == nil && len(entries) > 0 {
		b.BackingImages = make(map[string]BackingImageV2CopyInfoApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.BackingImages[k] = v
	}
	return b
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithIP(value string) *InstanceManagerStatusApplyConfiguration {
	b.IP = &value
	return b
}

// WithAPIMinVersion sets the APIMinVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIMinVersion field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithAPIMinVersion(value int) *InstanceManagerStatusApplyConfiguration {
	b.APIMinVersion = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithAPIVersion(value int) *InstanceManagerStatusApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithProxyAPIMinVersion sets the ProxyAPIMinVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyAPIMinVersion field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithProxyAPIMinVersion(value int) *InstanceManagerStatusApplyConfiguration {
	b.ProxyAPIMinVersion = &value
	return b
}

// WithProxyAPIVersion sets the ProxyAPIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyAPIVersion field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithProxyAPIVersion(value int) *InstanceManagerStatusApplyConfiguration {
	b.ProxyAPIVersion = &value
	return b
}

// WithDataEngineStatus sets the DataEngineStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataEngineStatus field is set to the value of the last call.
func (b *InstanceManagerStatusApplyConfiguration) WithDataEngineStatus(value *DataEngineStatusApplyConfiguration) *InstanceManagerStatusApplyConfiguration {
	b.DataEngineStatus = value
	return b
}

// WithInstances puts the entries into the Instances field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Instances field,
// overwriting an existing map entries in Instances field with the same key.
func (b *InstanceManagerStatusApplyConfiguration) WithInstances(entries map[string]InstanceProcessApplyConfiguration) *InstanceManagerStatusApplyConfiguration {
	if b.Instances == nil && len(entries) > 0 {
		b.Instances = make(map[string]InstanceProcessApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Instances[k] = v
	}
	return b
}
