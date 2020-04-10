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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	networkservicev1alpha1 "github.com/networkservicemesh/networkservicemesh/k8s/pkg/apis/networkservice/v1alpha1"
	versioned "github.com/networkservicemesh/networkservicemesh/k8s/pkg/networkservice/clientset/versioned"
	internalinterfaces "github.com/networkservicemesh/networkservicemesh/k8s/pkg/networkservice/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/networkservicemesh/networkservicemesh/k8s/pkg/networkservice/listers/networkservice/v1alpha1"
)

// NetworkServiceManagerInformer provides access to a shared informer and lister for
// NetworkServiceManagers.
type NetworkServiceManagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NetworkServiceManagerLister
}

type networkServiceManagerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkServiceManagerInformer constructs a new informer for NetworkServiceManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkServiceManagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkServiceManagerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkServiceManagerInformer constructs a new informer for NetworkServiceManager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkServiceManagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkserviceV1alpha1().NetworkServiceManagers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkserviceV1alpha1().NetworkServiceManagers(namespace).Watch(context.TODO(), options)
			},
		},
		&networkservicev1alpha1.NetworkServiceManager{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkServiceManagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkServiceManagerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkServiceManagerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkservicev1alpha1.NetworkServiceManager{}, f.defaultInformer)
}

func (f *networkServiceManagerInformer) Lister() v1alpha1.NetworkServiceManagerLister {
	return v1alpha1.NewNetworkServiceManagerLister(f.Informer().GetIndexer())
}
