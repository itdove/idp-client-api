// Copyright Red Hat

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/identitatem/idp-client-api/api/client/clientset/versioned/typed/identitatem/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIdentityconfigV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIdentityconfigV1alpha1) AuthRealms(namespace string) v1alpha1.AuthRealmInterface {
	return &FakeAuthRealms{c, namespace}
}

func (c *FakeIdentityconfigV1alpha1) Strategies(namespace string) v1alpha1.StrategyInterface {
	return &FakeStrategies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIdentityconfigV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
