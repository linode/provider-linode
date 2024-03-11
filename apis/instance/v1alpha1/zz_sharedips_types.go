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

type SharedIPsInitParameters struct {

	// The set of IPs to share with the Linode.
	// A set of IP addresses to share to the Linode
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The ID of the Linode to share the IPs to.
	// The ID of the Linode to share these IP addresses with.
	// +crossplane:generate:reference:type=Instance
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// Reference to a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDRef *v1.Reference `json:"linodeIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDSelector *v1.Selector `json:"linodeIdSelector,omitempty" tf:"-"`
}

type SharedIPsObservation struct {

	// The set of IPs to share with the Linode.
	// A set of IP addresses to share to the Linode
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Linode to share the IPs to.
	// The ID of the Linode to share these IP addresses with.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`
}

type SharedIPsParameters struct {

	// The set of IPs to share with the Linode.
	// A set of IP addresses to share to the Linode
	// +kubebuilder:validation:Optional
	// +listType=set
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// The ID of the Linode to share the IPs to.
	// The ID of the Linode to share these IP addresses with.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// Reference to a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDRef *v1.Reference `json:"linodeIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDSelector *v1.Selector `json:"linodeIdSelector,omitempty" tf:"-"`
}

// SharedIPsSpec defines the desired state of SharedIPs
type SharedIPsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedIPsParameters `json:"forProvider"`
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
	InitProvider SharedIPsInitParameters `json:"initProvider,omitempty"`
}

// SharedIPsStatus defines the observed state of SharedIPs.
type SharedIPsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedIPsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SharedIPs is the Schema for the SharedIPss API. Manages IP addresses shared to a Linode.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type SharedIPs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addresses) || (has(self.initProvider) && has(self.initProvider.addresses))",message="spec.forProvider.addresses is a required parameter"
	Spec   SharedIPsSpec   `json:"spec"`
	Status SharedIPsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedIPsList contains a list of SharedIPss
type SharedIPsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedIPs `json:"items"`
}

// Repository type metadata.
var (
	SharedIPs_Kind             = "SharedIPs"
	SharedIPs_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedIPs_Kind}.String()
	SharedIPs_KindAPIVersion   = SharedIPs_Kind + "." + CRDGroupVersion.String()
	SharedIPs_GroupVersionKind = CRDGroupVersion.WithKind(SharedIPs_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedIPs{}, &SharedIPsList{})
}
