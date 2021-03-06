/*
   Copyright 2018 Amadeus SAS

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

// Generated file, do not modify manually!

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	cronworkflow_v1 "github.com/amadeusitgroup/workflow-controller/pkg/api/cronworkflow/v1"
	versioned "github.com/amadeusitgroup/workflow-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/amadeusitgroup/workflow-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/amadeusitgroup/workflow-controller/pkg/client/listers/cronworkflow/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CronWorkflowInformer provides access to a shared informer and lister for
// CronWorkflows.
type CronWorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CronWorkflowLister
}

type cronWorkflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCronWorkflowInformer constructs a new informer for CronWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCronWorkflowInformer constructs a new informer for CronWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CronworkflowV1().CronWorkflows(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CronworkflowV1().CronWorkflows(namespace).Watch(options)
			},
		},
		&cronworkflow_v1.CronWorkflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronWorkflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronWorkflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cronworkflow_v1.CronWorkflow{}, f.defaultInformer)
}

func (f *cronWorkflowInformer) Lister() v1.CronWorkflowLister {
	return v1.NewCronWorkflowLister(f.Informer().GetIndexer())
}
