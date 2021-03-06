/*
Copyright 2021 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "tableflip.dev/maths/pkg/apis/maths/v1alpha1"
)

// SubtractLister helps list Subtracts.
// All objects returned here must be treated as read-only.
type SubtractLister interface {
	// List lists all Subtracts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Subtract, err error)
	// Subtracts returns an object that can list and get Subtracts.
	Subtracts(namespace string) SubtractNamespaceLister
	SubtractListerExpansion
}

// subtractLister implements the SubtractLister interface.
type subtractLister struct {
	indexer cache.Indexer
}

// NewSubtractLister returns a new SubtractLister.
func NewSubtractLister(indexer cache.Indexer) SubtractLister {
	return &subtractLister{indexer: indexer}
}

// List lists all Subtracts in the indexer.
func (s *subtractLister) List(selector labels.Selector) (ret []*v1alpha1.Subtract, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subtract))
	})
	return ret, err
}

// Subtracts returns an object that can list and get Subtracts.
func (s *subtractLister) Subtracts(namespace string) SubtractNamespaceLister {
	return subtractNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubtractNamespaceLister helps list and get Subtracts.
// All objects returned here must be treated as read-only.
type SubtractNamespaceLister interface {
	// List lists all Subtracts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Subtract, err error)
	// Get retrieves the Subtract from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Subtract, error)
	SubtractNamespaceListerExpansion
}

// subtractNamespaceLister implements the SubtractNamespaceLister
// interface.
type subtractNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Subtracts in the indexer for a given namespace.
func (s subtractNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Subtract, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subtract))
	})
	return ret, err
}

// Get retrieves the Subtract from the indexer for a given namespace and name.
func (s subtractNamespaceLister) Get(name string) (*v1alpha1.Subtract, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subtract"), name)
	}
	return obj.(*v1alpha1.Subtract), nil
}
