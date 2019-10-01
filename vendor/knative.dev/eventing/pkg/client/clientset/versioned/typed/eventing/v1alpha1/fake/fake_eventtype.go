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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
)

// FakeEventTypes implements EventTypeInterface
type FakeEventTypes struct {
	Fake *FakeEventingV1alpha1
	ns   string
}

var eventtypesResource = schema.GroupVersionResource{Group: "eventing.knative.dev", Version: "v1alpha1", Resource: "eventtypes"}

var eventtypesKind = schema.GroupVersionKind{Group: "eventing.knative.dev", Version: "v1alpha1", Kind: "EventType"}

// Get takes name of the eventType, and returns the corresponding eventType object, and an error if there is any.
func (c *FakeEventTypes) Get(name string, options v1.GetOptions) (result *v1alpha1.EventType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventtypesResource, c.ns, name), &v1alpha1.EventType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventType), err
}

// List takes label and field selectors, and returns the list of EventTypes that match those selectors.
func (c *FakeEventTypes) List(opts v1.ListOptions) (result *v1alpha1.EventTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventtypesResource, eventtypesKind, c.ns, opts), &v1alpha1.EventTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventTypeList{ListMeta: obj.(*v1alpha1.EventTypeList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventTypes.
func (c *FakeEventTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventtypesResource, c.ns, opts))

}

// Create takes the representation of a eventType and creates it.  Returns the server's representation of the eventType, and an error, if there is any.
func (c *FakeEventTypes) Create(eventType *v1alpha1.EventType) (result *v1alpha1.EventType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventtypesResource, c.ns, eventType), &v1alpha1.EventType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventType), err
}

// Update takes the representation of a eventType and updates it. Returns the server's representation of the eventType, and an error, if there is any.
func (c *FakeEventTypes) Update(eventType *v1alpha1.EventType) (result *v1alpha1.EventType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventtypesResource, c.ns, eventType), &v1alpha1.EventType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventTypes) UpdateStatus(eventType *v1alpha1.EventType) (*v1alpha1.EventType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventtypesResource, "status", c.ns, eventType), &v1alpha1.EventType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventType), err
}

// Delete takes name of the eventType and deletes it. Returns an error if one occurs.
func (c *FakeEventTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventtypesResource, c.ns, name), &v1alpha1.EventType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventtypesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventTypeList{})
	return err
}

// Patch applies the patch and returns the patched eventType.
func (c *FakeEventTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventtypesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventType), err
}
