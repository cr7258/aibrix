/*
Copyright 2024 The Aibrix Team.

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

package v1alpha1

import (
	v1alpha1 "github.com/vllm-project/aibrix/api/model/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelAdapterStatusApplyConfiguration represents an declarative configuration of the ModelAdapterStatus type for use
// with apply.
type ModelAdapterStatusApplyConfiguration struct {
	Phase      *v1alpha1.ModelAdapterPhase `json:"phase,omitempty"`
	Conditions []v1.Condition              `json:"conditions,omitempty"`
	Instances  []string                    `json:"instances,omitempty"`
}

// ModelAdapterStatusApplyConfiguration constructs an declarative configuration of the ModelAdapterStatus type for use with
// apply.
func ModelAdapterStatus() *ModelAdapterStatusApplyConfiguration {
	return &ModelAdapterStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *ModelAdapterStatusApplyConfiguration) WithPhase(value v1alpha1.ModelAdapterPhase) *ModelAdapterStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ModelAdapterStatusApplyConfiguration) WithConditions(values ...v1.Condition) *ModelAdapterStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithInstances adds the given value to the Instances field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Instances field.
func (b *ModelAdapterStatusApplyConfiguration) WithInstances(values ...string) *ModelAdapterStatusApplyConfiguration {
	for i := range values {
		b.Instances = append(b.Instances, values[i])
	}
	return b
}
