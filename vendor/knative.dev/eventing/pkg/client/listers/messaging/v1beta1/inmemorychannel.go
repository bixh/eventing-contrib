/*
Copyright 2020 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing/pkg/apis/messaging/v1beta1"
)

// InMemoryChannelLister helps list InMemoryChannels.
type InMemoryChannelLister interface {
	// List lists all InMemoryChannels in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.InMemoryChannel, err error)
	// InMemoryChannels returns an object that can list and get InMemoryChannels.
	InMemoryChannels(namespace string) InMemoryChannelNamespaceLister
	InMemoryChannelListerExpansion
}

// inMemoryChannelLister implements the InMemoryChannelLister interface.
type inMemoryChannelLister struct {
	indexer cache.Indexer
}

// NewInMemoryChannelLister returns a new InMemoryChannelLister.
func NewInMemoryChannelLister(indexer cache.Indexer) InMemoryChannelLister {
	return &inMemoryChannelLister{indexer: indexer}
}

// List lists all InMemoryChannels in the indexer.
func (s *inMemoryChannelLister) List(selector labels.Selector) (ret []*v1beta1.InMemoryChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.InMemoryChannel))
	})
	return ret, err
}

// InMemoryChannels returns an object that can list and get InMemoryChannels.
func (s *inMemoryChannelLister) InMemoryChannels(namespace string) InMemoryChannelNamespaceLister {
	return inMemoryChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InMemoryChannelNamespaceLister helps list and get InMemoryChannels.
type InMemoryChannelNamespaceLister interface {
	// List lists all InMemoryChannels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.InMemoryChannel, err error)
	// Get retrieves the InMemoryChannel from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.InMemoryChannel, error)
	InMemoryChannelNamespaceListerExpansion
}

// inMemoryChannelNamespaceLister implements the InMemoryChannelNamespaceLister
// interface.
type inMemoryChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InMemoryChannels in the indexer for a given namespace.
func (s inMemoryChannelNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.InMemoryChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.InMemoryChannel))
	})
	return ret, err
}

// Get retrieves the InMemoryChannel from the indexer for a given namespace and name.
func (s inMemoryChannelNamespaceLister) Get(name string) (*v1beta1.InMemoryChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("inmemorychannel"), name)
	}
	return obj.(*v1beta1.InMemoryChannel), nil
}
