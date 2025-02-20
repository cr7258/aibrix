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

package rayclusterfleet

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	rayclusterv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	orchestrationv1alpha1 "github.com/vllm-project/aibrix/api/orchestration/v1alpha1"
)

var _ = Describe("RayClusterFleet Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		rayclusterfleet := &orchestrationv1alpha1.RayClusterFleet{}

		BeforeEach(func() {
			By("creating the custom resource for the Kind RayClusterFleet")
			err := k8sClient.Get(ctx, typeNamespacedName, rayclusterfleet)
			if err != nil && errors.IsNotFound(err) {
				resource := &orchestrationv1alpha1.RayClusterFleet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: orchestrationv1alpha1.RayClusterFleetSpec{

						Replicas: ptr.To(int32(1)),
						Strategy: appsv1.DeploymentStrategy{
							Type: appsv1.RollingUpdateDeploymentStrategyType,
						},
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"foot": "bar",
							},
						},
						Template: orchestrationv1alpha1.RayClusterTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{
									"foo": "bar",
								},
							},
							Spec: rayclusterv1.RayClusterSpec{
								RayVersion: "fake-ray-version",
								HeadGroupSpec: rayclusterv1.HeadGroupSpec{
									ServiceType: corev1.ServiceTypeClusterIP,
									Template: corev1.PodTemplateSpec{
										Spec: corev1.PodSpec{
											Containers: []corev1.Container{
												{
													Name:  "ray-head",
													Image: "rayproject/ray:2.10.0",
												},
											},
										},
									},
									RayStartParams: map[string]string{
										"dashboard-host": "0.0.0.0",
									},
								},
							},
						},
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &orchestrationv1alpha1.RayClusterFleet{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance RayClusterFleet")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})
		It("should successfully reconcile the resource", func() {
			// TODO: (@jiaxin.shan) Let's re-enable the test once the integration test is done.
			Skip("Skipping this test for now")
			By("Reconciling the created resource")
			controllerReconciler := &RayClusterFleetReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.
		})
	})
})
