/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// MongoDBRoute provides a way to route TCP requests. When combined with a Gateway
// listener, it can be used to forward connections on the port specified by the
// listener to a set of backends specified by the MongoDBRoute.
type MongoDBRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of MongoDBRoute.
	Spec MongoDBRouteSpec `json:"spec"`

	// Status defines the current state of MongoDBRoute.
	Status MongoDBRouteStatus `json:"status,omitempty"`
}

// MongoDBRouteSpec defines the desired state of MongoDBRoute
type MongoDBRouteSpec struct {
	gwv1beta1.CommonRouteSpec `json:",inline"`

	// Rules are a list of TCP matchers and actions.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	Rules []MongoDBRouteRule `json:"rules"`
}

// MongoDBRouteStatus defines the observed state of MongoDBRoute
type MongoDBRouteStatus struct {
	gwv1beta1.RouteStatus `json:",inline"`
}

// MongoDBRouteRule is the configuration for a given rule.
type MongoDBRouteRule struct {
	// BackendRefs defines the backend(s) where matching requests should be
	// sent. If unspecified or invalid (refers to a non-existent resource or a
	// Service with no endpoints), the underlying implementation MUST actively
	// reject connection attempts to this backend. Connection rejections must
	// respect weight; if an invalid backend is requested to have 80% of
	// connections, then 80% of connections must be rejected instead.
	//
	// Support: Core for Kubernetes Service
	//
	// Support: Extended for Kubernetes ServiceImport
	//
	// Support: Implementation-specific for any other resource
	//
	// Support for weight: Extended
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	BackendRefs []gwv1beta1.BackendRef `json:"backendRefs,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDBRouteList contains a list of MongoDBRoute
type MongoDBRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDBRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDBRoute{}, &MongoDBRouteList{})
}
