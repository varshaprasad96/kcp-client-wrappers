package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"github.com/kcp-dev/kcp-client-wrappers/client"
	"github.com/kcp-dev/kcp-client-wrappers/roundtripper"
	lcluster "github.com/kcp-dev/logicalcluster"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	clusterName := logicalcluster.New("root:default")
	scopedContext := roundtripper.WithCluster(ctx, lcluster.New("root:default"))
	klog.Info(clusterName)

	clusterClient, err := client.NewForConfig(config, lcluster.New("root:default"))
	if err != nil {
		klog.Fatal(err)
	}

	sec, err := clusterClient.Cluster(clusterName).CoreV1().Secrets("default").Get(scopedContext, "mysecret", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	fmt.Println(sec.Name)
	fmt.Println(sec.ClusterName)
}
