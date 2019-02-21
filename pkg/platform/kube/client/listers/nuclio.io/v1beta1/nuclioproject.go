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

package v1beta1

import (
	v1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/apis/nuclio.io/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NuclioProjectLister helps list NuclioProjects.
type NuclioProjectLister interface {
	// List lists all NuclioProjects in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.NuclioProject, err error)
	// NuclioProjects returns an object that can list and get NuclioProjects.
	NuclioProjects(namespace string) NuclioProjectNamespaceLister
	NuclioProjectListerExpansion
}

// nuclioProjectLister implements the NuclioProjectLister interface.
type nuclioProjectLister struct {
	indexer cache.Indexer
}

// NewNuclioProjectLister returns a new NuclioProjectLister.
func NewNuclioProjectLister(indexer cache.Indexer) NuclioProjectLister {
	return &nuclioProjectLister{indexer: indexer}
}

// List lists all NuclioProjects in the indexer.
func (s *nuclioProjectLister) List(selector labels.Selector) (ret []*v1beta1.NuclioProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.NuclioProject))
	})
	return ret, err
}

// NuclioProjects returns an object that can list and get NuclioProjects.
func (s *nuclioProjectLister) NuclioProjects(namespace string) NuclioProjectNamespaceLister {
	return nuclioProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NuclioProjectNamespaceLister helps list and get NuclioProjects.
type NuclioProjectNamespaceLister interface {
	// List lists all NuclioProjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.NuclioProject, err error)
	// Get retrieves the NuclioProject from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.NuclioProject, error)
	NuclioProjectNamespaceListerExpansion
}

// nuclioProjectNamespaceLister implements the NuclioProjectNamespaceLister
// interface.
type nuclioProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NuclioProjects in the indexer for a given namespace.
func (s nuclioProjectNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.NuclioProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.NuclioProject))
	})
	return ret, err
}

// Get retrieves the NuclioProject from the indexer for a given namespace and name.
func (s nuclioProjectNamespaceLister) Get(name string) (*v1beta1.NuclioProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("nuclioproject"), name)
	}
	return obj.(*v1beta1.NuclioProject), nil
}
