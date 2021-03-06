/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
)

// InitializerConfigurationLister helps list InitializerConfigurations.
type InitializerConfigurationLister interface {
	// List lists all InitializerConfigurations in the indexer.
	List(selector labels.Selector) (ret []*admissionregistration.InitializerConfiguration, err error)
	// InitializerConfigurations returns an object that can list and get InitializerConfigurations.
	InitializerConfigurations(namespace string) InitializerConfigurationNamespaceLister
	InitializerConfigurationListerExpansion
}

// initializerConfigurationLister implements the InitializerConfigurationLister interface.
type initializerConfigurationLister struct {
	indexer cache.Indexer
}

// NewInitializerConfigurationLister returns a new InitializerConfigurationLister.
func NewInitializerConfigurationLister(indexer cache.Indexer) InitializerConfigurationLister {
	return &initializerConfigurationLister{indexer: indexer}
}

// List lists all InitializerConfigurations in the indexer.
func (s *initializerConfigurationLister) List(selector labels.Selector) (ret []*admissionregistration.InitializerConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistration.InitializerConfiguration))
	})
	return ret, err
}

// InitializerConfigurations returns an object that can list and get InitializerConfigurations.
func (s *initializerConfigurationLister) InitializerConfigurations(namespace string) InitializerConfigurationNamespaceLister {
	return initializerConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InitializerConfigurationNamespaceLister helps list and get InitializerConfigurations.
type InitializerConfigurationNamespaceLister interface {
	// List lists all InitializerConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*admissionregistration.InitializerConfiguration, err error)
	// Get retrieves the InitializerConfiguration from the indexer for a given namespace and name.
	Get(name string) (*admissionregistration.InitializerConfiguration, error)
	InitializerConfigurationNamespaceListerExpansion
}

// initializerConfigurationNamespaceLister implements the InitializerConfigurationNamespaceLister
// interface.
type initializerConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InitializerConfigurations in the indexer for a given namespace.
func (s initializerConfigurationNamespaceLister) List(selector labels.Selector) (ret []*admissionregistration.InitializerConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*admissionregistration.InitializerConfiguration))
	})
	return ret, err
}

// Get retrieves the InitializerConfiguration from the indexer for a given namespace and name.
func (s initializerConfigurationNamespaceLister) Get(name string) (*admissionregistration.InitializerConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(admissionregistration.Resource("initializerconfiguration"), name)
	}
	return obj.(*admissionregistration.InitializerConfiguration), nil
}
