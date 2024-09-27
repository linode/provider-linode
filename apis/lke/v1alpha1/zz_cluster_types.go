/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ACLInitParameters struct {

	// A list of ip addresses to allow.
	Addresses []AddressesInitParameters `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ACLObservation struct {

	// A list of ip addresses to allow.
	Addresses []AddressesObservation `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ACLParameters struct {

	// A list of ip addresses to allow.
	// +kubebuilder:validation:Optional
	Addresses []AddressesParameters `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	// Defines default policy. A value of true results in a default policy of DENY. A value of false results in default policy of ALLOW, and has the same effect as delete the ACL configuration.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AddressesInitParameters struct {

	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// +listType=set
	IPv4 []*string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// +listType=set
	IPv6 []*string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
}

type AddressesObservation struct {

	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// +listType=set
	IPv4 []*string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// +listType=set
	IPv6 []*string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
}

type AddressesParameters struct {

	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// A set of individual ipv4 addresses or CIDRs to ALLOW.
	// +kubebuilder:validation:Optional
	// +listType=set
	IPv4 []*string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// A set of individual ipv6 addresses or CIDRs to ALLOW.
	// +kubebuilder:validation:Optional
	// +listType=set
	IPv6 []*string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`
}

type AutoscalerInitParameters struct {

	// The maximum number of nodes to autoscale to.
	// The maximum number of nodes to autoscale to.
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// The minimum number of nodes to autoscale to.
	// The minimum number of nodes to autoscale to.
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AutoscalerObservation struct {

	// The maximum number of nodes to autoscale to.
	// The maximum number of nodes to autoscale to.
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// The minimum number of nodes to autoscale to.
	// The minimum number of nodes to autoscale to.
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`
}

type AutoscalerParameters struct {

	// The maximum number of nodes to autoscale to.
	// The maximum number of nodes to autoscale to.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum number of nodes to autoscale to.
	// The minimum number of nodes to autoscale to.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type ClusterInitParameters struct {

	// Defines settings for the Kubernetes Control Plane.
	ControlPlane []ControlPlaneInitParameters `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// A set of node pool tags to ignore when planning and applying this cluster. This prevents externally managed node pools from being deleted or unintentionally updated on subsequent applies. See Externally Managed Node Pools for more details.
	// An array of tags indicating that node pools having those tags are defined with a separate nodepool resource, rather than inside the current cluster resource.
	// +listType=set
	ExternalPoolTags []*string `json:"externalPoolTags,omitempty" tf:"external_pool_tags,omitempty"`

	// The desired Kubernetes version for this Kubernetes cluster in the format of major.minor (e.g. 1.21), and the latest supported patch version will be deployed.
	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// This Kubernetes cluster's unique label.
	// The unique label for the cluster.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Additional nested attributes:
	// A node pool in the cluster.
	Pool []PoolInitParameters `json:"pool,omitempty" tf:"pool,omitempty"`

	// This Kubernetes cluster's location.
	// This cluster's location.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterObservation struct {

	// The endpoints for the Kubernetes API server.
	// The API endpoints for the cluster.
	APIEndpoints []*string `json:"apiEndpoints,omitempty" tf:"api_endpoints,omitempty"`

	// Defines settings for the Kubernetes Control Plane.
	ControlPlane []ControlPlaneObservation `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// The Kubernetes Dashboard access URL for this cluster.
	// The dashboard URL of the cluster.
	DashboardURL *string `json:"dashboardUrl,omitempty" tf:"dashboard_url,omitempty"`

	// A set of node pool tags to ignore when planning and applying this cluster. This prevents externally managed node pools from being deleted or unintentionally updated on subsequent applies. See Externally Managed Node Pools for more details.
	// An array of tags indicating that node pools having those tags are defined with a separate nodepool resource, rather than inside the current cluster resource.
	// +listType=set
	ExternalPoolTags []*string `json:"externalPoolTags,omitempty" tf:"external_pool_tags,omitempty"`

	// The ID of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The desired Kubernetes version for this Kubernetes cluster in the format of major.minor (e.g. 1.21), and the latest supported patch version will be deployed.
	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// This Kubernetes cluster's unique label.
	// The unique label for the cluster.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Additional nested attributes:
	// A node pool in the cluster.
	Pool []PoolObservation `json:"pool,omitempty" tf:"pool,omitempty"`

	// This Kubernetes cluster's location.
	// This cluster's location.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The status of the cluster.
	// The status of the cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterParameters struct {

	// Defines settings for the Kubernetes Control Plane.
	// +kubebuilder:validation:Optional
	ControlPlane []ControlPlaneParameters `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// A set of node pool tags to ignore when planning and applying this cluster. This prevents externally managed node pools from being deleted or unintentionally updated on subsequent applies. See Externally Managed Node Pools for more details.
	// An array of tags indicating that node pools having those tags are defined with a separate nodepool resource, rather than inside the current cluster resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExternalPoolTags []*string `json:"externalPoolTags,omitempty" tf:"external_pool_tags,omitempty"`

	// The desired Kubernetes version for this Kubernetes cluster in the format of major.minor (e.g. 1.21), and the latest supported patch version will be deployed.
	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	// +kubebuilder:validation:Optional
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// This Kubernetes cluster's unique label.
	// The unique label for the cluster.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Additional nested attributes:
	// A node pool in the cluster.
	// +kubebuilder:validation:Optional
	Pool []PoolParameters `json:"pool,omitempty" tf:"pool,omitempty"`

	// This Kubernetes cluster's location.
	// This cluster's location.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ControlPlaneInitParameters struct {

	// Defines the ACL configuration for an LKE cluster's control plane.
	ACL []ACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Defines whether High Availability is enabled for the cluster Control Plane. This is an irreversible change.
	// Defines whether High Availability is enabled for the Control Plane Components of the cluster.
	HighAvailability *bool `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`
}

type ControlPlaneObservation struct {

	// Defines the ACL configuration for an LKE cluster's control plane.
	ACL []ACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// Defines whether High Availability is enabled for the cluster Control Plane. This is an irreversible change.
	// Defines whether High Availability is enabled for the Control Plane Components of the cluster.
	HighAvailability *bool `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`
}

type ControlPlaneParameters struct {

	// Defines the ACL configuration for an LKE cluster's control plane.
	// +kubebuilder:validation:Optional
	ACL []ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// Defines whether High Availability is enabled for the cluster Control Plane. This is an irreversible change.
	// Defines whether High Availability is enabled for the Control Plane Components of the cluster.
	// +kubebuilder:validation:Optional
	HighAvailability *bool `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`
}

type NodesInitParameters struct {
}

type NodesObservation struct {

	// The ID of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the underlying Linode instance.
	InstanceID *float64 `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The status of the cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NodesParameters struct {
}

type PoolInitParameters struct {

	// When specified, the number of nodes autoscales within the defined minimum and maximum values.
	Autoscaler []AutoscalerInitParameters `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// The number of nodes in the Node Pool. If undefined with an autoscaler the initial node count will equal the autoscaler minimum.
	// The number of nodes in the Node Pool.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// A set of tags applied to this node pool.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A Linode Type for all of the nodes in the Node Pool. See all node types here.
	// A Linode Type for all of the nodes in the Node Pool.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PoolObservation struct {

	// When specified, the number of nodes autoscales within the defined minimum and maximum values.
	Autoscaler []AutoscalerObservation `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// The number of nodes in the Node Pool. If undefined with an autoscaler the initial node count will equal the autoscaler minimum.
	// The number of nodes in the Node Pool.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The disk encryption policy for nodes in this pool.
	// The disk encryption policy for the nodes in this pool. NOTE: Disk encryption may not currently be available to all users.
	DiskEncryption *string `json:"diskEncryption,omitempty" tf:"disk_encryption,omitempty"`

	// The ID of the cluster.
	// The ID of the Node Pool.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// The nodes in the node pool.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// A set of tags applied to this node pool.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A Linode Type for all of the nodes in the Node Pool. See all node types here.
	// A Linode Type for all of the nodes in the Node Pool.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PoolParameters struct {

	// When specified, the number of nodes autoscales within the defined minimum and maximum values.
	// +kubebuilder:validation:Optional
	Autoscaler []AutoscalerParameters `json:"autoscaler,omitempty" tf:"autoscaler,omitempty"`

	// The number of nodes in the Node Pool. If undefined with an autoscaler the initial node count will equal the autoscaler minimum.
	// The number of nodes in the Node Pool.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// An array of tags applied to the Kubernetes cluster. Tags are case-insensitive and are for organizational purposes only.
	// A set of tags applied to this node pool.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A Linode Type for all of the nodes in the Node Pool. See all node types here.
	// A Linode Type for all of the nodes in the Node Pool.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cluster is the Schema for the Clusters API. Manages a Linode instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.k8sVersion) || (has(self.initProvider) && has(self.initProvider.k8sVersion))",message="spec.forProvider.k8sVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pool) || (has(self.initProvider) && has(self.initProvider.pool))",message="spec.forProvider.pool is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
