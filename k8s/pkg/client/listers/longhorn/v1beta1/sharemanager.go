/*
Copyright The Longhorn Authors.

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
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ShareManagerLister helps list ShareManagers.
// All objects returned here must be treated as read-only.
type ShareManagerLister interface {
	// List lists all ShareManagers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.ShareManager, err error)
	// ShareManagers returns an object that can list and get ShareManagers.
	ShareManagers(namespace string) ShareManagerNamespaceLister
	ShareManagerListerExpansion
}

// shareManagerLister implements the ShareManagerLister interface.
type shareManagerLister struct {
	listers.ResourceIndexer[*longhornv1beta1.ShareManager]
}

// NewShareManagerLister returns a new ShareManagerLister.
func NewShareManagerLister(indexer cache.Indexer) ShareManagerLister {
	return &shareManagerLister{listers.New[*longhornv1beta1.ShareManager](indexer, longhornv1beta1.Resource("sharemanager"))}
}

// ShareManagers returns an object that can list and get ShareManagers.
func (s *shareManagerLister) ShareManagers(namespace string) ShareManagerNamespaceLister {
	return shareManagerNamespaceLister{listers.NewNamespaced[*longhornv1beta1.ShareManager](s.ResourceIndexer, namespace)}
}

// ShareManagerNamespaceLister helps list and get ShareManagers.
// All objects returned here must be treated as read-only.
type ShareManagerNamespaceLister interface {
	// List lists all ShareManagers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.ShareManager, err error)
	// Get retrieves the ShareManager from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*longhornv1beta1.ShareManager, error)
	ShareManagerNamespaceListerExpansion
}

// shareManagerNamespaceLister implements the ShareManagerNamespaceLister
// interface.
type shareManagerNamespaceLister struct {
	listers.ResourceIndexer[*longhornv1beta1.ShareManager]
}
