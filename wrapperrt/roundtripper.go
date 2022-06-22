package wrapperrt

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/kcp-dev/logicalcluster"
)

type ClusterRt struct {
	Rt          http.RoundTripper
	ClusterName string
}

func (rt *ClusterRt) RoundTrip(req *http.Request) (*http.Response, error) {
	cluster, ok := ClusterFromContext(req.Context())
	if !ok {
		return nil, fmt.Errorf("expected cluster in context")
	}
	req = req.Clone(req.Context())
	req.URL.Path = generatePath(req.URL.Path, cluster)
	req.URL.RawPath = generatePath(req.URL.RawPath, cluster)

	return rt.Rt.RoundTrip(req)
}

// apiRegex matches any string that has /api/ or /apis/ in it.
var apiRegex = regexp.MustCompile(`(/api/|/apis/)`)

// generatePath formats the request path to target the specified cluster
func generatePath(originalPath string, cluster logicalcluster.Name) string {
	// If the originalPath has /api/ or /apis/ in it, it might be anywhere in the path, so we use a regex to find and
	// replaces /api/ or /apis/ with $cluster/api/ or $cluster/apis/
	if apiRegex.MatchString(originalPath) {
		return apiRegex.ReplaceAllString(originalPath, fmt.Sprintf("%s$1", cluster.Path()))
	}

	// Otherwise, we're just prepending /clusters/$name
	path := cluster.Path()

	// if the original path is relative, add a / separator
	if len(originalPath) > 0 && originalPath[0] != '/' {
		path += "/"
	}

	// finally append the original path
	path += originalPath

	return path
}

type key int

const (
	keyCluster key = iota
)

// WithCluster injects a cluster name into a context
func WithCluster(ctx context.Context, cluster logicalcluster.Name) context.Context {
	return context.WithValue(ctx, keyCluster, cluster)
}

// ClusterFromContext extracts a cluster name from the context
func ClusterFromContext(ctx context.Context) (logicalcluster.Name, bool) {
	s, ok := ctx.Value(keyCluster).(logicalcluster.Name)
	return s, ok
}
