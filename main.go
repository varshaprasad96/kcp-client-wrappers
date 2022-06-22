package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/kcp-dev/kcp-client-wrappers/wrapperrt"
	"github.com/kcp-dev/logicalcluster"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

	config.Wrap(func(rt http.RoundTripper) http.RoundTripper {
		return &wrapperrt.ClusterRt{
			Rt:          rt,
			ClusterName: "root:default",
		}
	})

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	ctx := context.Background()
	scopedContext := wrapperrt.WithCluster(ctx, logicalcluster.New("root:default"))

	sec, err := client.CoreV1().Secrets("default").Get(scopedContext, "mysecret", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	fmt.Println(sec.Name)
	fmt.Println(sec.ClusterName)
}
