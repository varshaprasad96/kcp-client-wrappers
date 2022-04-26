package main

import (
	"context"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/kcp-client-wrappers/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	r, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: "admin"}).ClientConfig()
	if err != nil {
		klog.Fatal(err)
	}

	clusterClient, err := kubernetes.NewForConfig(r)
	if err != nil {
		klog.Fatal(err)
	}

	ctx := context.Background()
	clusterName := logicalcluster.New("admin")
	klog.Info("admin")
	list, err := clusterClient.Cluster(clusterName).RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	for _, cr := range list.Items {
		klog.InfoS("listed", "clusterName", cr.ClusterName, "name", cr.Name)
	}

	klog.Info("source")
	clusterName = logicalcluster.New("admin_source")
	list, err = clusterClient.Cluster(clusterName).RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	for _, cr := range list.Items {
		klog.InfoS("listed", "clusterName", cr.ClusterName, "name", cr.Name)
	}
}
