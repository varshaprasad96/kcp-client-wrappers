package kubernetes

import (
	"context"
	"fmt"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	rbacapiv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	rbacapplyv1 "k8s.io/client-go/applyconfigurations/rbac/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	flowcontrolv1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/rest"
)

func NewForConfig(config *rest.Config) (*ClusterClient, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP client: %w", err)
	}

	clusterRoundTripper := kcp.NewClusterRoundTripper(client.Transport)
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
	return &wrappedInterface{
		cluster:  cluster,
		delegate: c.delegate,
	}
}

type wrappedInterface struct {
	cluster  logicalcluster.LogicalCluster
	delegate kubernetes.Interface
}

func (w *wrappedInterface) Discovery() discovery.DiscoveryInterface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AppsV1() appsv1.AppsV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AutoscalingV2() autoscalingv2.AutoscalingV2Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) BatchV1() batchv1.BatchV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) CertificatesV1() certificatesv1.CertificatesV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) CoordinationV1() coordinationv1.CoordinationV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) CoreV1() corev1.CoreV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) EventsV1() eventsv1.EventsV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) NetworkingV1() networkingv1.NetworkingV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) NodeV1() nodev1.NodeV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) PolicyV1() policyv1.PolicyV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) RbacV1() rbacv1.RbacV1Interface {
	return &wrappedRbacV1{
		cluster:  w.cluster,
		delegate: w.delegate.RbacV1(),
	}
}

func (w *wrappedInterface) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) SchedulingV1() schedulingv1.SchedulingV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) StorageV1() storagev1.StorageV1Interface {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedInterface) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	panic("not implemented") // TODO: Implement
}

type wrappedRbacV1 struct {
	cluster  logicalcluster.LogicalCluster
	delegate rbacv1.RbacV1Interface
}

func (w *wrappedRbacV1) RESTClient() rest.Interface {
	return w.delegate.RESTClient()
}

func (w *wrappedRbacV1) ClusterRoles() rbacv1.ClusterRoleInterface {
	return &wrappedClusterRole{
		cluster:  w.cluster,
		delegate: w.delegate.ClusterRoles(),
	}
}

func (w *wrappedRbacV1) ClusterRoleBindings() rbacv1.ClusterRoleBindingInterface {
	panic("no")
}

func (w *wrappedRbacV1) Roles(namespace string) rbacv1.RoleInterface {
	panic("no")
}

func (w *wrappedRbacV1) RoleBindings(namespace string) rbacv1.RoleBindingInterface {
	panic("no")
}

type wrappedClusterRole struct {
	cluster  logicalcluster.LogicalCluster
	delegate rbacv1.ClusterRoleInterface
}

func (w *wrappedClusterRole) Create(ctx context.Context, clusterRole *rbacapiv1.ClusterRole, opts metav1.CreateOptions) (*rbacapiv1.ClusterRole, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		ctx = kcp.WithCluster(ctx, w.cluster)
	} else if ctxCluster != w.cluster {
		return nil, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return w.delegate.Create(ctx, clusterRole, opts)
}

func (w *wrappedClusterRole) Update(ctx context.Context, clusterRole *rbacapiv1.ClusterRole, opts metav1.UpdateOptions) (*rbacapiv1.ClusterRole, error) {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) Get(ctx context.Context, name string, opts metav1.GetOptions) (*rbacapiv1.ClusterRole, error) {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) List(ctx context.Context, opts metav1.ListOptions) (*rbacapiv1.ClusterRoleList, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		ctx = kcp.WithCluster(ctx, w.cluster)
	} else if ctxCluster != w.cluster {
		return nil, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedClusterRole) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *rbacapiv1.ClusterRole, err error) {
	panic("not implemented") // TODO: Implement
}

func (w *wrappedClusterRole) Apply(ctx context.Context, clusterRole *rbacapplyv1.ClusterRoleApplyConfiguration, opts metav1.ApplyOptions) (result *rbacapiv1.ClusterRole, err error) {
	panic("not implemented") // TODO: Implement
}
