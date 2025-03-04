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

package v1beta1

import (
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
)

// BackingImageDataSourceStatusApplyConfiguration represents a declarative configuration of the BackingImageDataSourceStatus type for use
// with apply.
type BackingImageDataSourceStatusApplyConfiguration struct {
	OwnerID           *string                            `json:"ownerID,omitempty"`
	RunningParameters map[string]string                  `json:"runningParameters,omitempty"`
	CurrentState      *longhornv1beta1.BackingImageState `json:"currentState,omitempty"`
	Size              *int64                             `json:"size,omitempty"`
	Progress          *int                               `json:"progress,omitempty"`
	Checksum          *string                            `json:"checksum,omitempty"`
	Message           *string                            `json:"message,omitempty"`
}

// BackingImageDataSourceStatusApplyConfiguration constructs a declarative configuration of the BackingImageDataSourceStatus type for use with
// apply.
func BackingImageDataSourceStatus() *BackingImageDataSourceStatusApplyConfiguration {
	return &BackingImageDataSourceStatusApplyConfiguration{}
}

// WithOwnerID sets the OwnerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OwnerID field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithOwnerID(value string) *BackingImageDataSourceStatusApplyConfiguration {
	b.OwnerID = &value
	return b
}

// WithRunningParameters puts the entries into the RunningParameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the RunningParameters field,
// overwriting an existing map entries in RunningParameters field with the same key.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithRunningParameters(entries map[string]string) *BackingImageDataSourceStatusApplyConfiguration {
	if b.RunningParameters == nil && len(entries) > 0 {
		b.RunningParameters = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.RunningParameters[k] = v
	}
	return b
}

// WithCurrentState sets the CurrentState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentState field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithCurrentState(value longhornv1beta1.BackingImageState) *BackingImageDataSourceStatusApplyConfiguration {
	b.CurrentState = &value
	return b
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithSize(value int64) *BackingImageDataSourceStatusApplyConfiguration {
	b.Size = &value
	return b
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithProgress(value int) *BackingImageDataSourceStatusApplyConfiguration {
	b.Progress = &value
	return b
}

// WithChecksum sets the Checksum field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Checksum field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithChecksum(value string) *BackingImageDataSourceStatusApplyConfiguration {
	b.Checksum = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *BackingImageDataSourceStatusApplyConfiguration) WithMessage(value string) *BackingImageDataSourceStatusApplyConfiguration {
	b.Message = &value
	return b
}
