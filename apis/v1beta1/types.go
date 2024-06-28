/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// Credentials required to authenticate to this provider.
	Credentials ProviderCredentials `json:"credentials"`

	// +kubebuilder:validation:Optional
	Configuration ProviderConfiguration `json:"config"`
}

// ProviderConfiguration for configuring the terraform provider
// see https://registry.terraform.io/providers/linode/linode/latest/docs#configuration-reference
type ProviderConfiguration struct {
	// +kubebuilder:validation:Optional
	UserAgentPrefix string `json:"ua_prefix"`
	// +kubebuilder:validation:Optional
	SkipInstanceReadyPoll bool `json:"skip_instance_ready_poll"`
	// +kubebuilder:validation:Optional
	SkipInstanceDeletePoll bool `json:"skip_instance_delete_poll"`
	// +kubebuilder:validation:Optional
	SkipImplicitReboots bool `json:"skip_implicit_reboots"`
	// +kubebuilder:validation:Optional
	DisableInternalCache bool `json:"disable_internal_cache"`
	// +kubebuilder:validation:Optional
	MinRetryDelayms int `json:"min_retry_delay_ms"`
	// +kubebuilder:validation:Optional
	MaxRetryDelayms int `json:"max_retry_delay_ms"`
	// +kubebuilder:validation:Optional
	EventPollms int `json:"event_poll_ms"`
	// +kubebuilder:validation:Optional
	LKEEventPollms int `json:"lke_event_poll_ms"`
	// +kubebuilder:validation:Optional
	LKENodeReadyPollms int `json:"lke_node_ready_poll_ms"`
	// +kubebuilder:validation:Optional
	ObjAccessKey string `json:"obj_access_key"`
	// +kubebuilder:validation:Optional
	ObjSecretKey string `json:"obj_secret_key"`
	// +kubebuilder:validation:Optional
	ObjUseTempKeys bool `json:"obj_use_temp_keys"`
	// +kubebuilder:validation:Optional
	ObjForceDelete bool `json:"obj_bucket_force_delete"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Linode provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,linode}
// +kubebuilder:storageversion
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,providerconfig,linode}
// +kubebuilder:storageversion
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
