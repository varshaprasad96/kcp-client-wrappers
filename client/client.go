package client

import (
	"fmt"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/kcp-client-wrappers/roundtripper"
	lcluster "github.com/kcp-dev/logicalcluster"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewForConfig(config *rest.Config, cluster lcluster.Name) (*ClusterClient, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}

	clusterRoundTripper := roundtripper.NewClusterRoundTripper(client.Transport, cluster)
	client.Transport = clusterRoundTripper

	delegate, err := kubernetes.NewForConfigAndClient(config, client)
	if err != nil {
		return nil, fmt.Errorf("error creating delegate clientset: %w", err)
	}

	return &ClusterClient{
		delegate: delegate,
	}, nil
}

type ClusterClient struct {
	delegate kubernetes.Interface
}

func (c *ClusterClient) Cluster(cluster logicalcluster.LogicalCluster) kubernetes.Interface {
	return c.delegate
}
