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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/networkservicemesh/networkservicemesh/k8s/pkg/apis/networkservice/v1alpha1"
	scheme "github.com/networkservicemesh/networkservicemesh/k8s/pkg/networkservice/clientset/versioned/scheme"
)

// NetworkServiceManagersGetter has a method to return a NetworkServiceManagerInterface.
// A group's client should implement this interface.
type NetworkServiceManagersGetter interface {
	NetworkServiceManagers(namespace string) NetworkServiceManagerInterface
}

// NetworkServiceManagerInterface has methods to work with NetworkServiceManager resources.
type NetworkServiceManagerInterface interface {
	Create(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.CreateOptions) (*v1alpha1.NetworkServiceManager, error)
	Update(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.UpdateOptions) (*v1alpha1.NetworkServiceManager, error)
	UpdateStatus(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.UpdateOptions) (*v1alpha1.NetworkServiceManager, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkServiceManager, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NetworkServiceManagerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkServiceManager, err error)
	NetworkServiceManagerExpansion
}

// networkServiceManagers implements NetworkServiceManagerInterface
type networkServiceManagers struct {
	client rest.Interface
	ns     string
}

// newNetworkServiceManagers returns a NetworkServiceManagers
func newNetworkServiceManagers(c *NetworkserviceV1alpha1Client, namespace string) *networkServiceManagers {
	return &networkServiceManagers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkServiceManager, and returns the corresponding networkServiceManager object, and an error if there is any.
func (c *networkServiceManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkServiceManager, err error) {
	result = &v1alpha1.NetworkServiceManager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkServiceManagers that match those selectors.
func (c *networkServiceManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkServiceManagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkServiceManagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkServiceManagers.
func (c *networkServiceManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkServiceManager and creates it.  Returns the server's representation of the networkServiceManager, and an error, if there is any.
func (c *networkServiceManagers) Create(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.CreateOptions) (result *v1alpha1.NetworkServiceManager, err error) {
	result = &v1alpha1.NetworkServiceManager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkServiceManager).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkServiceManager and updates it. Returns the server's representation of the networkServiceManager, and an error, if there is any.
func (c *networkServiceManagers) Update(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.UpdateOptions) (result *v1alpha1.NetworkServiceManager, err error) {
	result = &v1alpha1.NetworkServiceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		Name(networkServiceManager.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkServiceManager).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkServiceManagers) UpdateStatus(ctx context.Context, networkServiceManager *v1alpha1.NetworkServiceManager, opts v1.UpdateOptions) (result *v1alpha1.NetworkServiceManager, err error) {
	result = &v1alpha1.NetworkServiceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		Name(networkServiceManager.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkServiceManager).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkServiceManager and deletes it. Returns an error if one occurs.
func (c *networkServiceManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkServiceManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkservicemanagers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkServiceManager.
func (c *networkServiceManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkServiceManager, err error) {
	result = &v1alpha1.NetworkServiceManager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkservicemanagers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
