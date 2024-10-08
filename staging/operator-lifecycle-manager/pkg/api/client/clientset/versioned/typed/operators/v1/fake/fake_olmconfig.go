/*
Copyright Red Hat, Inc.

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

	v1 "github.com/operator-framework/api/pkg/operators/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOLMConfigs implements OLMConfigInterface
type FakeOLMConfigs struct {
	Fake *FakeOperatorsV1
}

var olmconfigsResource = v1.SchemeGroupVersion.WithResource("olmconfigs")

var olmconfigsKind = v1.SchemeGroupVersion.WithKind("OLMConfig")

// Get takes name of the oLMConfig, and returns the corresponding oLMConfig object, and an error if there is any.
func (c *FakeOLMConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OLMConfig, err error) {
	emptyResult := &v1.OLMConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(olmconfigsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OLMConfig), err
}

// List takes label and field selectors, and returns the list of OLMConfigs that match those selectors.
func (c *FakeOLMConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OLMConfigList, err error) {
	emptyResult := &v1.OLMConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(olmconfigsResource, olmconfigsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OLMConfigList{ListMeta: obj.(*v1.OLMConfigList).ListMeta}
	for _, item := range obj.(*v1.OLMConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oLMConfigs.
func (c *FakeOLMConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(olmconfigsResource, opts))
}

// Create takes the representation of a oLMConfig and creates it.  Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *FakeOLMConfigs) Create(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.CreateOptions) (result *v1.OLMConfig, err error) {
	emptyResult := &v1.OLMConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(olmconfigsResource, oLMConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OLMConfig), err
}

// Update takes the representation of a oLMConfig and updates it. Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *FakeOLMConfigs) Update(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (result *v1.OLMConfig, err error) {
	emptyResult := &v1.OLMConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(olmconfigsResource, oLMConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OLMConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOLMConfigs) UpdateStatus(ctx context.Context, oLMConfig *v1.OLMConfig, opts metav1.UpdateOptions) (result *v1.OLMConfig, err error) {
	emptyResult := &v1.OLMConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(olmconfigsResource, "status", oLMConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OLMConfig), err
}

// Delete takes name of the oLMConfig and deletes it. Returns an error if one occurs.
func (c *FakeOLMConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(olmconfigsResource, name, opts), &v1.OLMConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOLMConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(olmconfigsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OLMConfigList{})
	return err
}

// Patch applies the patch and returns the patched oLMConfig.
func (c *FakeOLMConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OLMConfig, err error) {
	emptyResult := &v1.OLMConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(olmconfigsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OLMConfig), err
}
