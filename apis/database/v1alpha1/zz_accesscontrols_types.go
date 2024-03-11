// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AccessControlsInitParameters struct {

	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// +listType=set
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// The unique ID of the target database.
	// The ID of the database to manage the allow list for.
	DatabaseID *float64 `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The unique type of the target database. (mysql, mongodb, postgresql)
	// The type of the  database to manage the allow list for.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`
}

type AccessControlsObservation struct {

	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// +listType=set
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// The unique ID of the target database.
	// The ID of the database to manage the allow list for.
	DatabaseID *float64 `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The unique type of the target database. (mysql, mongodb, postgresql)
	// The type of the  database to manage the allow list for.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccessControlsParameters struct {

	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// The unique ID of the target database.
	// The ID of the database to manage the allow list for.
	// +kubebuilder:validation:Optional
	DatabaseID *float64 `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The unique type of the target database. (mysql, mongodb, postgresql)
	// The type of the  database to manage the allow list for.
	// +kubebuilder:validation:Optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`
}

// AccessControlsSpec defines the desired state of AccessControls
type AccessControlsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessControlsParameters `json:"forProvider"`
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
	InitProvider AccessControlsInitParameters `json:"initProvider,omitempty"`
}

// AccessControlsStatus defines the observed state of AccessControls.
type AccessControlsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessControlsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessControls is the Schema for the AccessControlss API. Manages the access controls for a Linode Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type AccessControls struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.allowList) || (has(self.initProvider) && has(self.initProvider.allowList))",message="spec.forProvider.allowList is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseId) || (has(self.initProvider) && has(self.initProvider.databaseId))",message="spec.forProvider.databaseId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseType) || (has(self.initProvider) && has(self.initProvider.databaseType))",message="spec.forProvider.databaseType is a required parameter"
	Spec   AccessControlsSpec   `json:"spec"`
	Status AccessControlsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessControlsList contains a list of AccessControlss
type AccessControlsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessControls `json:"items"`
}

// Repository type metadata.
var (
	AccessControls_Kind             = "AccessControls"
	AccessControls_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessControls_Kind}.String()
	AccessControls_KindAPIVersion   = AccessControls_Kind + "." + CRDGroupVersion.String()
	AccessControls_GroupVersionKind = CRDGroupVersion.WithKind(AccessControls_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessControls{}, &AccessControlsList{})
}
