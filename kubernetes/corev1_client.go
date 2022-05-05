package kubernetes

import (
	"context"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1api "k8s.io/client-go/applyconfigurations/core/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type wrappedCoreV1 struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.CoreV1Interface
}

func (w *wrappedCoreV1) RESTClient() rest.Interface {
	return w.delegate.RESTClient()
}

func (w *wrappedCoreV1) Events(namespace string) corev1.EventInterface {
	return nil
}

func (w *wrappedCoreV1) LimitRanges(namespace string) corev1.LimitRangeInterface {
	return nil
}

func (w *wrappedCoreV1) Namespaces() corev1.NamespaceInterface {
	return nil
}

func (w *wrappedCoreV1) Nodes() corev1.NodeInterface {
	return nil
}

func (w *wrappedCoreV1) PersistentVolumeClaims(namespace string) corev1.PersistentVolumeClaimInterface {
	return nil
}

func (c *wrappedCoreV1) PersistentVolumes() corev1.PersistentVolumeInterface {
	return nil
}

func (c *wrappedCoreV1) Pods(namespace string) corev1.PodInterface {
	return nil
}

func (c *wrappedCoreV1) PodTemplates(namespace string) corev1.PodTemplateInterface {
	return nil
}

func (c *wrappedCoreV1) ReplicationControllers(namespace string) corev1.ReplicationControllerInterface {
	return nil
}

func (c *wrappedCoreV1) ResourceQuotas(namespace string) corev1.ResourceQuotaInterface {
	return nil
}

func (c *wrappedCoreV1) Services(namespace string) corev1.ServiceInterface {
	return nil
}

func (c *wrappedCoreV1) ServiceAccounts(namespace string) corev1.ServiceAccountInterface {
	return nil
}

func (w *wrappedCoreV1) Secrets(namespace string) corev1.SecretInterface {
	return &wrappedSecrets{
		cluster:  w.cluster,
		delegate: w.delegate.Secrets(namespace),
	}
}

type wrappedSecrets struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.SecretInterface
}

func (w *wrappedSecrets) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) (*v1.Secret, error) {
	panic("no")
}

func (w *wrappedSecrets) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) (*v1.Secret, error) {
	panic("no")
}

func (w *wrappedSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("no")
}

func (w *wrappedSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("no")
}

func (w *wrappedSecrets) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	panic("no")
}

func (w *wrappedSecrets) List(ctx context.Context, opts metav1.ListOptions) (*v1.SecretList, error) {
	panic("no")
}

func (w *wrappedSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("no")
}

func (w *wrappedSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Secret, err error) {
	panic("no")
}

func (w *wrappedSecrets) Apply(ctx context.Context, secret *corev1api.SecretApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Secret, err error) {
	panic("no")
}

type wrappedEvents struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.EventInterface
}

func (w *wrappedEvents) Create(ctx context.Context, event *v1.Event, opts metav1.CreateOptions) (*v1.Event, error) {
	panic("no")
}

func (w *wrappedEvents) CreateWithEventNamespace(*v1.Event) (*v1.Event, error) {
	panic("no")
}

func (w *wrappedEvents) Update(ctx context.Context, event *v1.Event, opts metav1.UpdateOptions) (*v1.Event, error) {
	panic("no")
}

func (w *wrappedEvents) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("no")
}

func (w *wrappedEvents) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("no")
}

func (w *wrappedEvents) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Event, error) {
	panic("no")
}

func (w *wrappedEvents) List(ctx context.Context, opts metav1.ListOptions) (*v1.EventList, error) {
	panic("no")
}

func (w *wrappedEvents) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("no")
}
func (w *wrappedEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Event, err error) {
	panic("no")
}

func (w *wrappedEvents) Apply(ctx context.Context, event *corev1api.EventApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Event, err error) {
	panic("no")
}

func (w *wrappedCoreV1) Endpoints(namespace string) corev1.EndpointsInterface {
	return &wrappedEndpoints{
		cluster:  w.cluster,
		delegate: w.delegate.Endpoints(namespace),
	}
}

type wrappedEndpoints struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.EndpointsInterface
}

func (w *wrappedEndpoints) Create(ctx context.Context, endpoints *v1.Endpoints, opts metav1.CreateOptions) (*v1.Endpoints, error) {
	panic("no")
}

func (w *wrappedEndpoints) Update(ctx context.Context, endpoints *v1.Endpoints, opts metav1.UpdateOptions) (*v1.Endpoints, error) {
	panic("no")
}

func (w *wrappedEndpoints) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("no")
}

func (w *wrappedEndpoints) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("no")
}

func (w *wrappedEndpoints) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Endpoints, error) {
	panic("no")
}

func (w *wrappedEndpoints) List(ctx context.Context, opts metav1.ListOptions) (*v1.EndpointsList, error) {
	panic("no")
}

func (w *wrappedEndpoints) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("no")
}

func (w *wrappedEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoints, err error) {
	panic("no")
}

func (w *wrappedEndpoints) Apply(ctx context.Context, endpoints *corev1api.EndpointsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Endpoints, err error) {
	panic("no")
}

func (w *wrappedCoreV1) ComponentStatuses() corev1.ComponentStatusInterface {
	return &wrappedComponentStatuses{
		cluster:  w.cluster,
		delegate: w.delegate.ComponentStatuses(),
	}
}

type wrappedComponentStatuses struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.ComponentStatusInterface
}

func (w *wrappedComponentStatuses) Create(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.CreateOptions) (*v1.ComponentStatus, error) {
	panic("no")
}

func (w *wrappedComponentStatuses) Update(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.UpdateOptions) (*v1.ComponentStatus, error) {
	panic("no")
}

func (w *wrappedComponentStatuses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("no")
}

func (w *wrappedComponentStatuses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("no")
}

func (w *wrappedComponentStatuses) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error) {
	panic("no")
}

func (w *wrappedComponentStatuses) List(ctx context.Context, opts metav1.ListOptions) (*v1.ComponentStatusList, error) {
	panic("no")
}

func (w *wrappedComponentStatuses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("no")
}

func (w *wrappedComponentStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentStatus, err error) {
	panic("no")
}

func (w *wrappedComponentStatuses) Apply(ctx context.Context, componentStatus *corev1api.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ComponentStatus, err error) {
	panic("no")
}

func (w *wrappedCoreV1) ConfigMaps(namespace string) corev1.ConfigMapInterface {
	return &wrappedConfigMap{
		cluster:  w.cluster,
		delegate: w.delegate.ConfigMaps(namespace),
	}
}

type wrappedConfigMap struct {
	cluster  logicalcluster.LogicalCluster
	delegate corev1.ConfigMapInterface
}

func (w *wrappedConfigMap) Create(ctx context.Context, configMap *v1.ConfigMap, opts metav1.CreateOptions) (*v1.ConfigMap, error) {
	panic("no")
}

func (w *wrappedConfigMap) Update(ctx context.Context, configMap *v1.ConfigMap, opts metav1.UpdateOptions) (*v1.ConfigMap, error) {
	panic("no")
}

func (w *wrappedConfigMap) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	panic("no")
}

func (w *wrappedConfigMap) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	panic("no")
}
func (w *wrappedConfigMap) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ConfigMap, error) {
	panic("no")
}
func (w *wrappedConfigMap) List(ctx context.Context, opts metav1.ListOptions) (*v1.ConfigMapList, error) {
	panic("no")
}

func (w *wrappedConfigMap) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	panic("no")
}

func (w *wrappedConfigMap) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConfigMap, err error) {
	panic("no")
}
func (w *wrappedConfigMap) Apply(ctx context.Context, configMap *corev1api.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConfigMap, err error) {
	panic("no")
}
