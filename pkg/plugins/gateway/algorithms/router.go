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

package routingalgorithms

import (
	"context"

	v1 "k8s.io/api/core/v1"
)

type Algorithms string

// Router defines the interface for routing logic to select target pods.
type Router interface {
	// Route returns the target pod
	Route(ctx context.Context, pods map[string]*v1.Pod, model, message string) (string, error)
}

// Validate validates if user provided routing routers is supported by gateway
func Validate(algorithms Algorithms) bool {
	_, ok := routerStores[algorithms]
	return ok
}

// Select the user provided routing routers is supported by gateway
func Select(algorithms Algorithms) routerFunc {
	if !Validate(algorithms) {
		algorithms = RouterRandom
	}
	return routerRegistry[algorithms]
}

func Register(algorithms Algorithms, router routerFunc) {
	routerRegistry[algorithms] = router
	routerStores[algorithms] = struct{}{}
}

var routerRegistry = map[Algorithms]routerFunc{}
var routerStores = map[Algorithms]any{}

type routerFunc func() (Router, error)
