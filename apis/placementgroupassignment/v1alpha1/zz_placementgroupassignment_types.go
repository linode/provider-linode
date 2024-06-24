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

type PlacementGroupAssignmentInitParameters struct {
	CompliantOnly *bool `json:"compliantOnly,omitempty" tf:"compliant_only,omitempty"`

	// The unique ID of the Linode to assign.
	// A set of Linode IDs to assign to the Placement Group.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The unique ID of the target Placement Group.
	// The ID of the Placement Group for this assignment.
	PlacementGroupID *float64 `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`
}

type PlacementGroupAssignmentObservation struct {
	CompliantOnly *bool `json:"compliantOnly,omitempty" tf:"compliant_only,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique ID of the Linode to assign.
	// A set of Linode IDs to assign to the Placement Group.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The unique ID of the target Placement Group.
	// The ID of the Placement Group for this assignment.
	PlacementGroupID *float64 `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`
}

type PlacementGroupAssignmentParameters struct {

	// +kubebuilder:validation:Optional
	CompliantOnly *bool `json:"compliantOnly,omitempty" tf:"compliant_only,omitempty"`

	// The unique ID of the Linode to assign.
	// A set of Linode IDs to assign to the Placement Group.
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The unique ID of the target Placement Group.
	// The ID of the Placement Group for this assignment.
	// +kubebuilder:validation:Optional
	PlacementGroupID *float64 `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`
}

// PlacementGroupAssignmentSpec defines the desired state of PlacementGroupAssignment
type PlacementGroupAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlacementGroupAssignmentParameters `json:"forProvider"`
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
	InitProvider PlacementGroupAssignmentInitParameters `json:"initProvider,omitempty"`
}

// PlacementGroupAssignmentStatus defines the observed state of PlacementGroupAssignment.
type PlacementGroupAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlacementGroupAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PlacementGroupAssignment is the Schema for the PlacementGroupAssignments API. Manages a single assignment between a Linode and a Placement Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type PlacementGroupAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.linodeId) || (has(self.initProvider) && has(self.initProvider.linodeId))",message="spec.forProvider.linodeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.placementGroupId) || (has(self.initProvider) && has(self.initProvider.placementGroupId))",message="spec.forProvider.placementGroupId is a required parameter"
	Spec   PlacementGroupAssignmentSpec   `json:"spec"`
	Status PlacementGroupAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroupAssignmentList contains a list of PlacementGroupAssignments
type PlacementGroupAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementGroupAssignment `json:"items"`
}

// Repository type metadata.
var (
	PlacementGroupAssignment_Kind             = "PlacementGroupAssignment"
	PlacementGroupAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlacementGroupAssignment_Kind}.String()
	PlacementGroupAssignment_KindAPIVersion   = PlacementGroupAssignment_Kind + "." + CRDGroupVersion.String()
	PlacementGroupAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PlacementGroupAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PlacementGroupAssignment{}, &PlacementGroupAssignmentList{})
}
