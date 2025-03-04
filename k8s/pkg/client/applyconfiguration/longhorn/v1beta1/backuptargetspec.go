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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupTargetSpecApplyConfiguration represents a declarative configuration of the BackupTargetSpec type for use
// with apply.
type BackupTargetSpecApplyConfiguration struct {
	BackupTargetURL  *string      `json:"backupTargetURL,omitempty"`
	CredentialSecret *string      `json:"credentialSecret,omitempty"`
	PollInterval     *v1.Duration `json:"pollInterval,omitempty"`
	SyncRequestedAt  *v1.Time     `json:"syncRequestedAt,omitempty"`
}

// BackupTargetSpecApplyConfiguration constructs a declarative configuration of the BackupTargetSpec type for use with
// apply.
func BackupTargetSpec() *BackupTargetSpecApplyConfiguration {
	return &BackupTargetSpecApplyConfiguration{}
}

// WithBackupTargetURL sets the BackupTargetURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackupTargetURL field is set to the value of the last call.
func (b *BackupTargetSpecApplyConfiguration) WithBackupTargetURL(value string) *BackupTargetSpecApplyConfiguration {
	b.BackupTargetURL = &value
	return b
}

// WithCredentialSecret sets the CredentialSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CredentialSecret field is set to the value of the last call.
func (b *BackupTargetSpecApplyConfiguration) WithCredentialSecret(value string) *BackupTargetSpecApplyConfiguration {
	b.CredentialSecret = &value
	return b
}

// WithPollInterval sets the PollInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PollInterval field is set to the value of the last call.
func (b *BackupTargetSpecApplyConfiguration) WithPollInterval(value v1.Duration) *BackupTargetSpecApplyConfiguration {
	b.PollInterval = &value
	return b
}

// WithSyncRequestedAt sets the SyncRequestedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncRequestedAt field is set to the value of the last call.
func (b *BackupTargetSpecApplyConfiguration) WithSyncRequestedAt(value v1.Time) *BackupTargetSpecApplyConfiguration {
	b.SyncRequestedAt = &value
	return b
}
