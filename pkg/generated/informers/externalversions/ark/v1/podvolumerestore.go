/*
Copyright the Heptio Ark contributors.

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

package v1

import (
	time "time"

	ark_v1 "github.com/heptio/velero/pkg/apis/ark/v1"
	versioned "github.com/heptio/velero/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/heptio/velero/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/heptio/velero/pkg/generated/listers/ark/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodVolumeRestoreInformer provides access to a shared informer and lister for
// PodVolumeRestores.
type PodVolumeRestoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PodVolumeRestoreLister
}

type podVolumeRestoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodVolumeRestoreInformer constructs a new informer for PodVolumeRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodVolumeRestoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodVolumeRestoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodVolumeRestoreInformer constructs a new informer for PodVolumeRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodVolumeRestoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArkV1().PodVolumeRestores(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArkV1().PodVolumeRestores(namespace).Watch(options)
			},
		},
		&ark_v1.PodVolumeRestore{},
		resyncPeriod,
		indexers,
	)
}

func (f *podVolumeRestoreInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodVolumeRestoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podVolumeRestoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ark_v1.PodVolumeRestore{}, f.defaultInformer)
}

func (f *podVolumeRestoreInformer) Lister() v1.PodVolumeRestoreLister {
	return v1.NewPodVolumeRestoreLister(f.Informer().GetIndexer())
}
