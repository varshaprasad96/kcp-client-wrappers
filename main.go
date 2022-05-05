package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/kcp-client-wrappers/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	// r, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
	// 	clientcmd.NewDefaultClientConfigLoadingRules(),
	// 	&clientcmd.ConfigOverrides{CurrentContext: "default"}).ClientConfig()
	// if err != nil {
	// 	klog.Fatal(err)
	// }

	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clusterClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	ctx := context.Background()
	clusterName := logicalcluster.New("root:default")
	klog.Info(clusterName)

	fmt.Println(clusterClient.Cluster(clusterName).RbacV1().RESTClient().APIVersion())

	list, err := clusterClient.Cluster(clusterName).RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	fmt.Println(len(list.Items))

	// role, err := clusterClient.Cluster(clusterName).RbacV1().ClusterRoles().Get(ctx, "secret-reader", metav1.GetOptions{})
	// fmt.Println(role.Name)
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	// for _, cr := range list.Items {
	// 	klog.InfoS("listed", "clusterName", cr.ClusterName, "name", cr.Name)
	// }

	// klog.Info("source")
	// clusterName = logicalcluster.New("admin_source")
	// list, err = clusterClient.Cluster(clusterName).RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	// for _, cr := range list.Items {
	// 	klog.InfoS("listed", "clusterName", cr.ClusterName, "name", cr.Name)
	// }
}
