/*
Copyright 2019 The Knative Authors

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
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Add struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the Add (from the client).
	// +optional
	Spec AddSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the Add (from the controller).
	// +optional
	Status AddStatus `json:"status,omitempty"`
}

var (
	// Check that Add can be validated and defaulted.
	_ apis.Validatable   = (*Add)(nil)
	_ apis.Defaultable   = (*Add)(nil)
	_ kmeta.OwnerRefable = (*Add)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*Add)(nil)
)

// AddSpec holds the desired state of the Add (from the client).
type AddSpec struct {
	// ServiceName holds the name of the Kubernetes Service to expose as an "addressable".
	Operands []Operand `json:"add"`
}

type Operand struct {
	Ref   duckv1.KReference `json:"ref"`
	Value int               `json:"value"`
}

const (
	// AddConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	AddConditionReady = apis.ConditionReady
)

// AddStatus communicates the observed state of the Add (from the controller).
type AddStatus struct {
	duckv1.Status `json:",inline"`

	Expression string `json:"expression,omitempty"`
	Result     int    `json:"result"`
}

// AddList is a list of Add resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AddList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Add `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (a *Add) GetStatus() *duckv1.Status {
	return &a.Status.Status
}
