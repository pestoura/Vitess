/*
Copyright 2020 The Vitess Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1beta1 "vitess.io/vitess/go/vt/topo/k8stopo/apis/topo/v1beta1"
	scheme "vitess.io/vitess/go/vt/topo/k8stopo/client/clientset/versioned/scheme"
)

// VitessTopoNodesGetter has a method to return a VitessTopoNodeInterface.
// A group's client should implement this interface.
type VitessTopoNodesGetter interface {
	VitessTopoNodes(namespace string) VitessTopoNodeInterface
}

// VitessTopoNodeInterface has methods to work with VitessTopoNode resources.
type VitessTopoNodeInterface interface {
	Create(ctx context.Context, vitessTopoNode *v1beta1.VitessTopoNode, opts v1.CreateOptions) (*v1beta1.VitessTopoNode, error)
	Update(ctx context.Context, vitessTopoNode *v1beta1.VitessTopoNode, opts v1.UpdateOptions) (*v1beta1.VitessTopoNode, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VitessTopoNode, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VitessTopoNodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VitessTopoNode, err error)
	VitessTopoNodeExpansion
}

// vitessTopoNodes implements VitessTopoNodeInterface
type vitessTopoNodes struct {
	client rest.Interface
	ns     string
}

// newVitessTopoNodes returns a VitessTopoNodes
func newVitessTopoNodes(c *TopoV1beta1Client, namespace string) *vitessTopoNodes {
	return &vitessTopoNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vitessTopoNode, and returns the corresponding vitessTopoNode object, and an error if there is any.
func (c *vitessTopoNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VitessTopoNode, err error) {
	result = &v1beta1.VitessTopoNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VitessTopoNodes that match those selectors.
func (c *vitessTopoNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VitessTopoNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VitessTopoNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vitessTopoNodes.
func (c *vitessTopoNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vitessTopoNode and creates it.  Returns the server's representation of the vitessTopoNode, and an error, if there is any.
func (c *vitessTopoNodes) Create(ctx context.Context, vitessTopoNode *v1beta1.VitessTopoNode, opts v1.CreateOptions) (result *v1beta1.VitessTopoNode, err error) {
	result = &v1beta1.VitessTopoNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vitessTopoNode).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vitessTopoNode and updates it. Returns the server's representation of the vitessTopoNode, and an error, if there is any.
func (c *vitessTopoNodes) Update(ctx context.Context, vitessTopoNode *v1beta1.VitessTopoNode, opts v1.UpdateOptions) (result *v1beta1.VitessTopoNode, err error) {
	result = &v1beta1.VitessTopoNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		Name(vitessTopoNode.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vitessTopoNode).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vitessTopoNode and deletes it. Returns an error if one occurs.
func (c *vitessTopoNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vitessTopoNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vitesstoponodes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vitessTopoNode.
func (c *vitessTopoNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VitessTopoNode, err error) {
	result = &v1beta1.VitessTopoNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vitesstoponodes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
