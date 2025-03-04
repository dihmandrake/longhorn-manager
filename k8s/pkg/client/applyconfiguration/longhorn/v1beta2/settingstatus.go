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

// SettingStatusApplyConfiguration represents a declarative configuration of the SettingStatus type for use
// with apply.
type SettingStatusApplyConfiguration struct {
	Applied *bool `json:"applied,omitempty"`
}

// SettingStatusApplyConfiguration constructs a declarative configuration of the SettingStatus type for use with
// apply.
func SettingStatus() *SettingStatusApplyConfiguration {
	return &SettingStatusApplyConfiguration{}
}

// WithApplied sets the Applied field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Applied field is set to the value of the last call.
func (b *SettingStatusApplyConfiguration) WithApplied(value bool) *SettingStatusApplyConfiguration {
	b.Applied = &value
	return b
}
