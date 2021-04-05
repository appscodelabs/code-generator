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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/code-generator/examples/keyword/apis/default/v1"
)

// InterfaceLister helps list Interfaces.
// All objects returned here must be treated as read-only.
type InterfaceLister interface {
	// List lists all Interfaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Interface, err error)
	// Interfaces returns an object that can list and get Interfaces.
	Interfaces(namespace string) InterfaceNamespaceLister
	InterfaceListerExpansion
}

// interfaceLister implements the InterfaceLister interface.
type interfaceLister struct {
	indexer cache.Indexer
}

// NewInterfaceLister returns a new InterfaceLister.
func NewInterfaceLister(indexer cache.Indexer) InterfaceLister {
	return &interfaceLister{indexer: indexer}
}

// List lists all Interfaces in the indexer.
func (s *interfaceLister) List(selector labels.Selector) (ret []*v1.Interface, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Interface))
	})
	return ret, err
}

// Interfaces returns an object that can list and get Interfaces.
func (s *interfaceLister) Interfaces(namespace string) InterfaceNamespaceLister {
	return interfaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InterfaceNamespaceLister helps list and get Interfaces.
// All objects returned here must be treated as read-only.
type InterfaceNamespaceLister interface {
	// List lists all Interfaces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Interface, err error)
	// Get retrieves the Interface from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Interface, error)
	InterfaceNamespaceListerExpansion
}

// interfaceNamespaceLister implements the InterfaceNamespaceLister
// interface.
type interfaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Interfaces in the indexer for a given namespace.
func (s interfaceNamespaceLister) List(selector labels.Selector) (ret []*v1.Interface, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Interface))
	})
	return ret, err
}

// Get retrieves the Interface from the indexer for a given namespace and name.
func (s interfaceNamespaceLister) Get(name string) (*v1.Interface, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("interface"), name)
	}
	return obj.(*v1.Interface), nil
}
