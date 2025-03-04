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

// AttachmentTicketStatusApplyConfiguration represents a declarative configuration of the AttachmentTicketStatus type for use
// with apply.
type AttachmentTicketStatusApplyConfiguration struct {
	ID         *string                       `json:"id,omitempty"`
	Satisfied  *bool                         `json:"satisfied,omitempty"`
	Conditions []ConditionApplyConfiguration `json:"conditions,omitempty"`
	Generation *int64                        `json:"generation,omitempty"`
}

// AttachmentTicketStatusApplyConfiguration constructs a declarative configuration of the AttachmentTicketStatus type for use with
// apply.
func AttachmentTicketStatus() *AttachmentTicketStatusApplyConfiguration {
	return &AttachmentTicketStatusApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *AttachmentTicketStatusApplyConfiguration) WithID(value string) *AttachmentTicketStatusApplyConfiguration {
	b.ID = &value
	return b
}

// WithSatisfied sets the Satisfied field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Satisfied field is set to the value of the last call.
func (b *AttachmentTicketStatusApplyConfiguration) WithSatisfied(value bool) *AttachmentTicketStatusApplyConfiguration {
	b.Satisfied = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *AttachmentTicketStatusApplyConfiguration) WithConditions(values ...*ConditionApplyConfiguration) *AttachmentTicketStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *AttachmentTicketStatusApplyConfiguration) WithGeneration(value int64) *AttachmentTicketStatusApplyConfiguration {
	b.Generation = &value
	return b
}
