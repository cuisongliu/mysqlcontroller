/*
Copyright The Kubernetes Authors.

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
	mysqlcontrollerv1 "github.com/cuisongliu/mysqlcontroller/pkg/apis/mysqlcontroller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodGroups implements PodGroupInterface
type FakePodGroups struct {
	Fake *FakeMysqlcontrollerV1
	ns   string
}

var podgroupsResource = schema.GroupVersionResource{Group: "mysqlcontroller", Version: "v1", Resource: "podgroups"}

var podgroupsKind = schema.GroupVersionKind{Group: "mysqlcontroller", Version: "v1", Kind: "PodGroup"}

// Get takes name of the podGroup, and returns the corresponding podGroup object, and an error if there is any.
func (c *FakePodGroups) Get(name string, options v1.GetOptions) (result *mysqlcontrollerv1.PodGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podgroupsResource, c.ns, name), &mysqlcontrollerv1.PodGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlcontrollerv1.PodGroup), err
}

// List takes label and field selectors, and returns the list of PodGroups that match those selectors.
func (c *FakePodGroups) List(opts v1.ListOptions) (result *mysqlcontrollerv1.PodGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podgroupsResource, podgroupsKind, c.ns, opts), &mysqlcontrollerv1.PodGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mysqlcontrollerv1.PodGroupList{ListMeta: obj.(*mysqlcontrollerv1.PodGroupList).ListMeta}
	for _, item := range obj.(*mysqlcontrollerv1.PodGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podGroups.
func (c *FakePodGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podgroupsResource, c.ns, opts))

}

// Create takes the representation of a podGroup and creates it.  Returns the server's representation of the podGroup, and an error, if there is any.
func (c *FakePodGroups) Create(podGroup *mysqlcontrollerv1.PodGroup) (result *mysqlcontrollerv1.PodGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podgroupsResource, c.ns, podGroup), &mysqlcontrollerv1.PodGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlcontrollerv1.PodGroup), err
}

// Update takes the representation of a podGroup and updates it. Returns the server's representation of the podGroup, and an error, if there is any.
func (c *FakePodGroups) Update(podGroup *mysqlcontrollerv1.PodGroup) (result *mysqlcontrollerv1.PodGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podgroupsResource, c.ns, podGroup), &mysqlcontrollerv1.PodGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlcontrollerv1.PodGroup), err
}

// Delete takes name of the podGroup and deletes it. Returns an error if one occurs.
func (c *FakePodGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podgroupsResource, c.ns, name), &mysqlcontrollerv1.PodGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mysqlcontrollerv1.PodGroupList{})
	return err
}

// Patch applies the patch and returns the patched podGroup.
func (c *FakePodGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mysqlcontrollerv1.PodGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podgroupsResource, c.ns, name, pt, data, subresources...), &mysqlcontrollerv1.PodGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mysqlcontrollerv1.PodGroup), err
}