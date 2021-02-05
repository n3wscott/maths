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
type Subtract struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the Subtract (from the client).
	// +optional
	Spec SubtractSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the Subtract (from the controller).
	// +optional
	Status SubtractStatus `json:"status,omitempty"`
}

var (
	// Check that Subtract can be validated and defaulted.
	_ apis.Validatable   = (*Subtract)(nil)
	_ apis.Defaultable   = (*Subtract)(nil)
	_ kmeta.OwnerRefable = (*Subtract)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*Subtract)(nil)
)

// SubtractSpec holds the desired state of the Subtract (from the client).
type SubtractSpec struct {
	Operands []Operand `json:"sub"`
}

const (
	// SubtractConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	SubtractConditionReady = apis.ConditionReady
)

// SubtractStatus communicates the observed state of the Subtract (from the controller).
type SubtractStatus struct {
	duckv1.Status `json:",inline"`

	Expression string `json:"expression,omitempty"`
	Result     int    `json:"result"`
}

// SubtractList is a list of Subtract resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SubtractList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Subtract `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (a *Subtract) GetStatus() *duckv1.Status {
	return &a.Status.Status
}
