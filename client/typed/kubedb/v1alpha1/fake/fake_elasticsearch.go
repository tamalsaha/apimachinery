/*
Copyright 2018 The KubeDB Authors.

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

package fake

import (
	v1alpha1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElasticsearches implements ElasticsearchInterface
type FakeElasticsearches struct {
	Fake *FakeKubedbV1alpha1
	ns   string
}

var elasticsearchesResource = schema.GroupVersionResource{Group: "kubedb.com", Version: "v1alpha1", Resource: "elasticsearches"}

var elasticsearchesKind = schema.GroupVersionKind{Group: "kubedb.com", Version: "v1alpha1", Kind: "Elasticsearch"}

// Get takes name of the elasticsearch, and returns the corresponding elasticsearch object, and an error if there is any.
func (c *FakeElasticsearches) Get(name string, options v1.GetOptions) (result *v1alpha1.Elasticsearch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticsearchesResource, c.ns, name), &v1alpha1.Elasticsearch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elasticsearch), err
}

// List takes label and field selectors, and returns the list of Elasticsearches that match those selectors.
func (c *FakeElasticsearches) List(opts v1.ListOptions) (result *v1alpha1.ElasticsearchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticsearchesResource, elasticsearchesKind, c.ns, opts), &v1alpha1.ElasticsearchList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticsearchList{}
	for _, item := range obj.(*v1alpha1.ElasticsearchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticsearches.
func (c *FakeElasticsearches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticsearchesResource, c.ns, opts))

}

// Create takes the representation of a elasticsearch and creates it.  Returns the server's representation of the elasticsearch, and an error, if there is any.
func (c *FakeElasticsearches) Create(elasticsearch *v1alpha1.Elasticsearch) (result *v1alpha1.Elasticsearch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticsearchesResource, c.ns, elasticsearch), &v1alpha1.Elasticsearch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elasticsearch), err
}

// Update takes the representation of a elasticsearch and updates it. Returns the server's representation of the elasticsearch, and an error, if there is any.
func (c *FakeElasticsearches) Update(elasticsearch *v1alpha1.Elasticsearch) (result *v1alpha1.Elasticsearch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticsearchesResource, c.ns, elasticsearch), &v1alpha1.Elasticsearch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elasticsearch), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticsearches) UpdateStatus(elasticsearch *v1alpha1.Elasticsearch) (*v1alpha1.Elasticsearch, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(elasticsearchesResource, "status", c.ns, elasticsearch), &v1alpha1.Elasticsearch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elasticsearch), err
}

// Delete takes name of the elasticsearch and deletes it. Returns an error if one occurs.
func (c *FakeElasticsearches) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elasticsearchesResource, c.ns, name), &v1alpha1.Elasticsearch{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticsearches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticsearchesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticsearchList{})
	return err
}

// Patch applies the patch and returns the patched elasticsearch.
func (c *FakeElasticsearches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elasticsearch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticsearchesResource, c.ns, name, data, subresources...), &v1alpha1.Elasticsearch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Elasticsearch), err
}
