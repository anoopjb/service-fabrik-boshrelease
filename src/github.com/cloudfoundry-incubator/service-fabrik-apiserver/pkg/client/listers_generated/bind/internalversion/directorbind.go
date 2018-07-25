//TODO copyright header

// This file was automatically generated by lister-gen

package internalversion

import (
	bind "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DirectorBindLister helps list DirectorBinds.
type DirectorBindLister interface {
	// List lists all DirectorBinds in the indexer.
	List(selector labels.Selector) (ret []*bind.DirectorBind, err error)
	// DirectorBinds returns an object that can list and get DirectorBinds.
	DirectorBinds(namespace string) DirectorBindNamespaceLister
	DirectorBindListerExpansion
}

// directorBindLister implements the DirectorBindLister interface.
type directorBindLister struct {
	indexer cache.Indexer
}

// NewDirectorBindLister returns a new DirectorBindLister.
func NewDirectorBindLister(indexer cache.Indexer) DirectorBindLister {
	return &directorBindLister{indexer: indexer}
}

// List lists all DirectorBinds in the indexer.
func (s *directorBindLister) List(selector labels.Selector) (ret []*bind.DirectorBind, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*bind.DirectorBind))
	})
	return ret, err
}

// DirectorBinds returns an object that can list and get DirectorBinds.
func (s *directorBindLister) DirectorBinds(namespace string) DirectorBindNamespaceLister {
	return directorBindNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectorBindNamespaceLister helps list and get DirectorBinds.
type DirectorBindNamespaceLister interface {
	// List lists all DirectorBinds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*bind.DirectorBind, err error)
	// Get retrieves the DirectorBind from the indexer for a given namespace and name.
	Get(name string) (*bind.DirectorBind, error)
	DirectorBindNamespaceListerExpansion
}

// directorBindNamespaceLister implements the DirectorBindNamespaceLister
// interface.
type directorBindNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectorBinds in the indexer for a given namespace.
func (s directorBindNamespaceLister) List(selector labels.Selector) (ret []*bind.DirectorBind, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*bind.DirectorBind))
	})
	return ret, err
}

// Get retrieves the DirectorBind from the indexer for a given namespace and name.
func (s directorBindNamespaceLister) Get(name string) (*bind.DirectorBind, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(bind.Resource("directorbind"), name)
	}
	return obj.(*bind.DirectorBind), nil
}