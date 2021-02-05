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

package v1

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck/ducktypes"
	"knative.dev/pkg/kmeta"
)

// +genduck

type Results struct {
	Expression string `json:"expression,omitempty"`
	Result     int    `json:"result"`
}

var (
	// Results is a Convertible type.
	_ apis.Convertible = (*Results)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ResultsType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ResultsStatus `json:"status"`
}

type ResultsStatus struct {
	Results `json:",inline"`
}

// Verify AddressableType resources meet duck contracts.
var (
	_ apis.Listable         = (*ResultsType)(nil)
	_ ducktypes.Populatable = (*ResultsType)(nil)
	_ kmeta.OwnerRefable    = (*ResultsType)(nil)
)

// GetFullType implements duck.Implementable
func (*Results) GetFullType() ducktypes.Populatable {
	return &ResultsType{}
}

// ConvertTo implements apis.Convertible
func (a *Results) ConvertTo(ctx context.Context, to apis.Convertible) error {
	return fmt.Errorf("v1 is the highest known version, got: %T", to)
}

// ConvertFrom implements apis.Convertible
func (a *Results) ConvertFrom(ctx context.Context, from apis.Convertible) error {
	return fmt.Errorf("v1 is the highest known version, got: %T", from)
}

// Populate implements duck.Populatable
func (t *ResultsType) Populate() {
	t.Status = ResultsStatus{
		Results: Results{
			Expression: "expression",
			Result:     42,
		},
	}
}

// GetGroupVersionKind implements kmeta.OwnerRefable
func (t *ResultsType) GetGroupVersionKind() schema.GroupVersionKind {
	return t.GroupVersionKind()
}

// GetListType implements apis.Listable
func (*ResultsType) GetListType() runtime.Object {
	return &ResultsTypeList{}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResultsTypeList is a list of ResultsType resources
type ResultsTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ResultsType `json:"items"`
}
