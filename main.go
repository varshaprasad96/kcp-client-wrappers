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

	sec, err := clusterClient.Cluster(clusterName).CoreV1().Secrets("default").Get(ctx, "mysecret", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Println(sec.Name)
}
