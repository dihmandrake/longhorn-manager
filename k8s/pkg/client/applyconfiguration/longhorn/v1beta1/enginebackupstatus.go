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

// EngineBackupStatusApplyConfiguration represents a declarative configuration of the EngineBackupStatus type for use
// with apply.
type EngineBackupStatusApplyConfiguration struct {
	Progress       *int    `json:"progress,omitempty"`
	BackupURL      *string `json:"backupURL,omitempty"`
	Error          *string `json:"error,omitempty"`
	SnapshotName   *string `json:"snapshotName,omitempty"`
	State          *string `json:"state,omitempty"`
	ReplicaAddress *string `json:"replicaAddress,omitempty"`
}

// EngineBackupStatusApplyConfiguration constructs a declarative configuration of the EngineBackupStatus type for use with
// apply.
func EngineBackupStatus() *EngineBackupStatusApplyConfiguration {
	return &EngineBackupStatusApplyConfiguration{}
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithProgress(value int) *EngineBackupStatusApplyConfiguration {
	b.Progress = &value
	return b
}

// WithBackupURL sets the BackupURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackupURL field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithBackupURL(value string) *EngineBackupStatusApplyConfiguration {
	b.BackupURL = &value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithError(value string) *EngineBackupStatusApplyConfiguration {
	b.Error = &value
	return b
}

// WithSnapshotName sets the SnapshotName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SnapshotName field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithSnapshotName(value string) *EngineBackupStatusApplyConfiguration {
	b.SnapshotName = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithState(value string) *EngineBackupStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithReplicaAddress sets the ReplicaAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReplicaAddress field is set to the value of the last call.
func (b *EngineBackupStatusApplyConfiguration) WithReplicaAddress(value string) *EngineBackupStatusApplyConfiguration {
	b.ReplicaAddress = &value
	return b
}
