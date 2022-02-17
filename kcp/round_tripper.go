package kcp

import (
	"context"
	"fmt"
	"net/http"

	utilnet "k8s.io/apimachinery/pkg/util/net"
)

type ClusterRoundTripper struct {
	delegate http.RoundTripper
}

func NewClusterRoundTripper(delegate http.RoundTripper) *ClusterRoundTripper {
	return &ClusterRoundTripper{
		delegate: delegate,
	}
}

func (c *ClusterRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	cluster, ok := ClusterFromContext(req.Context())
	if !ok {
		return nil, fmt.Errorf("expected cluster in context")
	}

	// if !strings.HasPrefix(req.URL.Path, "/clusters/") {
	req = utilnet.CloneRequest(req)
	originalPath := req.URL.Path

	// start with /clusters/$name
	req.URL.Path = "/clusters/" + cluster

	// if the original path is relative, add a / separator
	if len(originalPath) > 0 && originalPath[0] != '/' {
		req.URL.Path += "/"
	}

	// finally append the original path
	req.URL.Path += originalPath
	// }

	return c.delegate.RoundTrip(req)
}

type key int

const (
	keyCluster key = iota
)

func WithCluster(ctx context.Context, cluster string) context.Context {
	return context.WithValue(ctx, keyCluster, cluster)
}

func ClusterFromContext(ctx context.Context) (string, bool) {
	s, ok := ctx.Value(keyCluster).(string)
	return s, ok
}
