//TODO copyright header

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DockerLister helps list Dockers.
type DockerLister interface {
	// List lists all Dockers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Docker, err error)
	// Dockers returns an object that can list and get Dockers.
	Dockers(namespace string) DockerNamespaceLister
	DockerListerExpansion
}

// dockerLister implements the DockerLister interface.
type dockerLister struct {
	indexer cache.Indexer
}

// NewDockerLister returns a new DockerLister.
func NewDockerLister(indexer cache.Indexer) DockerLister {
	return &dockerLister{indexer: indexer}
}

// List lists all Dockers in the indexer.
func (s *dockerLister) List(selector labels.Selector) (ret []*v1alpha1.Docker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Docker))
	})
	return ret, err
}

// Dockers returns an object that can list and get Dockers.
func (s *dockerLister) Dockers(namespace string) DockerNamespaceLister {
	return dockerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DockerNamespaceLister helps list and get Dockers.
type DockerNamespaceLister interface {
	// List lists all Dockers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Docker, err error)
	// Get retrieves the Docker from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Docker, error)
	DockerNamespaceListerExpansion
}

// dockerNamespaceLister implements the DockerNamespaceLister
// interface.
type dockerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Dockers in the indexer for a given namespace.
func (s dockerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Docker, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Docker))
	})
	return ret, err
}

// Get retrieves the Docker from the indexer for a given namespace and name.
func (s dockerNamespaceLister) Get(name string) (*v1alpha1.Docker, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("docker"), name)
	}
	return obj.(*v1alpha1.Docker), nil
}