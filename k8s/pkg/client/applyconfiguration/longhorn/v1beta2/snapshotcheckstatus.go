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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SnapshotCheckStatusApplyConfiguration represents a declarative configuration of the SnapshotCheckStatus type for use
// with apply.
type SnapshotCheckStatusApplyConfiguration struct {
	LastPeriodicCheckedAt *v1.Time `json:"lastPeriodicCheckedAt,omitempty"`
}

// SnapshotCheckStatusApplyConfiguration constructs a declarative configuration of the SnapshotCheckStatus type for use with
// apply.
func SnapshotCheckStatus() *SnapshotCheckStatusApplyConfiguration {
	return &SnapshotCheckStatusApplyConfiguration{}
}

// WithLastPeriodicCheckedAt sets the LastPeriodicCheckedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastPeriodicCheckedAt field is set to the value of the last call.
func (b *SnapshotCheckStatusApplyConfiguration) WithLastPeriodicCheckedAt(value v1.Time) *SnapshotCheckStatusApplyConfiguration {
	b.LastPeriodicCheckedAt = &value
	return b
}
