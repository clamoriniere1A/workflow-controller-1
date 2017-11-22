// Generated file, do not modify manually!
package v1

import (
	v1 "github.com/amadeusitgroup/workflow-controller/pkg/api/workflow/v1"
	"github.com/amadeusitgroup/workflow-controller/pkg/client/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type WorkflowV1Interface interface {
	RESTClient() rest.Interface
	WorkflowsGetter
}

// WorkflowV1Client is used to interact with features provided by the workflow group.
type WorkflowV1Client struct {
	restClient rest.Interface
}

func (c *WorkflowV1Client) Workflows(namespace string) WorkflowInterface {
	return newWorkflows(c, namespace)
}

// NewForConfig creates a new WorkflowV1Client for the given config.
func NewForConfig(c *rest.Config) (*WorkflowV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &WorkflowV1Client{client}, nil
}

// NewForConfigOrDie creates a new WorkflowV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WorkflowV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WorkflowV1Client for the given RESTClient.
func New(c rest.Interface) *WorkflowV1Client {
	return &WorkflowV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *WorkflowV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
