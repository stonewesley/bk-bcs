/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
	scheme "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/bcs-hook/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HookTemplatesGetter has a method to return a HookTemplateInterface.
// A group's client should implement this interface.
type HookTemplatesGetter interface {
	HookTemplates(namespace string) HookTemplateInterface
}

// HookTemplateInterface has methods to work with HookTemplate resources.
type HookTemplateInterface interface {
	Create(*v1alpha1.HookTemplate) (*v1alpha1.HookTemplate, error)
	Update(*v1alpha1.HookTemplate) (*v1alpha1.HookTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HookTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.HookTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HookTemplate, err error)
	HookTemplateExpansion
}

// hookTemplates implements HookTemplateInterface
type hookTemplates struct {
	client rest.Interface
	ns     string
}

// newHookTemplates returns a HookTemplates
func newHookTemplates(c *TkexV1alpha1Client, namespace string) *hookTemplates {
	return &hookTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hookTemplate, and returns the corresponding hookTemplate object, and an error if there is any.
func (c *hookTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.HookTemplate, err error) {
	result = &v1alpha1.HookTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hooktemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HookTemplates that match those selectors.
func (c *hookTemplates) List(opts v1.ListOptions) (result *v1alpha1.HookTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HookTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hooktemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hookTemplates.
func (c *hookTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hooktemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a hookTemplate and creates it.  Returns the server's representation of the hookTemplate, and an error, if there is any.
func (c *hookTemplates) Create(hookTemplate *v1alpha1.HookTemplate) (result *v1alpha1.HookTemplate, err error) {
	result = &v1alpha1.HookTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hooktemplates").
		Body(hookTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hookTemplate and updates it. Returns the server's representation of the hookTemplate, and an error, if there is any.
func (c *hookTemplates) Update(hookTemplate *v1alpha1.HookTemplate) (result *v1alpha1.HookTemplate, err error) {
	result = &v1alpha1.HookTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hooktemplates").
		Name(hookTemplate.Name).
		Body(hookTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the hookTemplate and deletes it. Returns an error if one occurs.
func (c *hookTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hooktemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hookTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hooktemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hookTemplate.
func (c *hookTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HookTemplate, err error) {
	result = &v1alpha1.HookTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hooktemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
