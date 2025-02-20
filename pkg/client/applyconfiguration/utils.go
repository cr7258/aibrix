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

package applyconfiguration

import (
	v1alpha1 "github.com/vllm-project/aibrix/api/autoscaling/v1alpha1"
	modelv1alpha1 "github.com/vllm-project/aibrix/api/model/v1alpha1"
	orchestrationv1alpha1 "github.com/vllm-project/aibrix/api/orchestration/v1alpha1"
	autoscalingv1alpha1 "github.com/vllm-project/aibrix/pkg/client/applyconfiguration/autoscaling/v1alpha1"
	applyconfigurationmodelv1alpha1 "github.com/vllm-project/aibrix/pkg/client/applyconfiguration/model/v1alpha1"
	applyconfigurationorchestrationv1alpha1 "github.com/vllm-project/aibrix/pkg/client/applyconfiguration/orchestration/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=autoscaling, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("MetricSource"):
		return &autoscalingv1alpha1.MetricSourceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PodAutoscaler"):
		return &autoscalingv1alpha1.PodAutoscalerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PodAutoscalerSpec"):
		return &autoscalingv1alpha1.PodAutoscalerSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PodAutoscalerStatus"):
		return &autoscalingv1alpha1.PodAutoscalerStatusApplyConfiguration{}

		// Group=model, Version=v1alpha1
	case modelv1alpha1.SchemeGroupVersion.WithKind("ModelAdapter"):
		return &applyconfigurationmodelv1alpha1.ModelAdapterApplyConfiguration{}
	case modelv1alpha1.SchemeGroupVersion.WithKind("ModelAdapterSpec"):
		return &applyconfigurationmodelv1alpha1.ModelAdapterSpecApplyConfiguration{}
	case modelv1alpha1.SchemeGroupVersion.WithKind("ModelAdapterStatus"):
		return &applyconfigurationmodelv1alpha1.ModelAdapterStatusApplyConfiguration{}

		// Group=orchestration, Version=v1alpha1
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterFleet"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterFleetApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterFleetCondition"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterFleetConditionApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterFleetSpec"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterFleetSpecApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterFleetStatus"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterFleetStatusApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterReplicaSet"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterReplicaSetApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterReplicaSetSpec"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterReplicaSetSpecApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterReplicaSetStatus"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterReplicaSetStatusApplyConfiguration{}
	case orchestrationv1alpha1.SchemeGroupVersion.WithKind("RayClusterTemplateSpec"):
		return &applyconfigurationorchestrationv1alpha1.RayClusterTemplateSpecApplyConfiguration{}

	}
	return nil
}
