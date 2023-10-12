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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "tableflip.dev/maths/pkg/apis/maths/v1alpha1"
)

// FakeAdds implements AddInterface
type FakeAdds struct {
	Fake *FakeMathsV1alpha1
	ns   string
}

var addsResource = v1alpha1.SchemeGroupVersion.WithResource("adds")

var addsKind = v1alpha1.SchemeGroupVersion.WithKind("Add")

// Get takes name of the add, and returns the corresponding add object, and an error if there is any.
func (c *FakeAdds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Add, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addsResource, c.ns, name), &v1alpha1.Add{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Add), err
}

// List takes label and field selectors, and returns the list of Adds that match those selectors.
func (c *FakeAdds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addsResource, addsKind, c.ns, opts), &v1alpha1.AddList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddList{ListMeta: obj.(*v1alpha1.AddList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested adds.
func (c *FakeAdds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addsResource, c.ns, opts))

}

// Create takes the representation of a add and creates it.  Returns the server's representation of the add, and an error, if there is any.
func (c *FakeAdds) Create(ctx context.Context, add *v1alpha1.Add, opts v1.CreateOptions) (result *v1alpha1.Add, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addsResource, c.ns, add), &v1alpha1.Add{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Add), err
}

// Update takes the representation of a add and updates it. Returns the server's representation of the add, and an error, if there is any.
func (c *FakeAdds) Update(ctx context.Context, add *v1alpha1.Add, opts v1.UpdateOptions) (result *v1alpha1.Add, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addsResource, c.ns, add), &v1alpha1.Add{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Add), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdds) UpdateStatus(ctx context.Context, add *v1alpha1.Add, opts v1.UpdateOptions) (*v1alpha1.Add, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(addsResource, "status", c.ns, add), &v1alpha1.Add{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Add), err
}

// Delete takes name of the add and deletes it. Returns an error if one occurs.
func (c *FakeAdds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(addsResource, c.ns, name, opts), &v1alpha1.Add{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddList{})
	return err
}

// Patch applies the patch and returns the patched add.
func (c *FakeAdds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Add, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Add{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Add), err
}
