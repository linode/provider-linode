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

type ImageObservation struct {

	// The capabilities of this Image.
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Whether this image supports cloud-init.
	CloudInit *bool `json:"cloudInit,omitempty" tf:"cloud_init,omitempty"`

	// When this Image was created.
	// When this Image was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The name of the User who created this Image.
	// The name of the User who created this Image.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// Whether or not this Image is deprecated. Will only be True for deprecated public Images.
	// Whether or not this Image is deprecated. Will only be True for deprecated public Images.
	Deprecated *bool `json:"deprecated,omitempty" tf:"deprecated,omitempty"`

	// A detailed description of this Image.
	// A detailed description of this Image.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Linode Disk that this Image will be created from.
	// The ID of the Linode Disk that this Image will be created from.
	DiskID *float64 `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// Only Images created automatically (from a deleted Linode; type=automatic) will expire.
	// Only Images created automatically (from a deleted Linode; type=automatic) will expire.
	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The MD5 hash of the file to be uploaded. This is used to trigger file updates.
	// The MD5 hash of the image file.
	FileHash *string `json:"fileHash,omitempty" tf:"file_hash,omitempty"`

	// The path of the image file to be uploaded.
	// The name of the file to upload to this image.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// The unique ID of this Image.  The ID of private images begin with private/ followed by the numeric identifier of the private image, for example private/12345.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// True if the Image is public.
	// True if the Image is public.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// A short description of the Image. Labels cannot contain special characters.
	// A short description of the Image. Labels cannot contain special characters.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the Linode that this Image will be created from.
	// The ID of the Linode that this Image will be created from.
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// The region of the image. See all regions here.
	// The region to upload to.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The minimum size this Image needs to deploy. Size is in MB.
	// The minimum size this Image needs to deploy. Size is in MB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The current status of this Image.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// How the Image was created. 'Manual' Images can be created at any time. 'Automatic' images are created automatically from a deleted Linode.
	// How the Image was created. 'Manual' Images can be created at any time. 'Automatic' images are created automatically from a deleted Linode.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The upstream distribution vendor. Nil for private Images.
	// The upstream distribution vendor. Nil for private Images.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type ImageParameters struct {

	// Whether this image supports cloud-init.
	// +kubebuilder:validation:Optional
	CloudInit *bool `json:"cloudInit,omitempty" tf:"cloud_init,omitempty"`

	// A detailed description of this Image.
	// A detailed description of this Image.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Linode Disk that this Image will be created from.
	// The ID of the Linode Disk that this Image will be created from.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/instance/v1alpha1.Disk
	// +kubebuilder:validation:Optional
	DiskID *float64 `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// Reference to a Disk in instance to populate diskId.
	// +kubebuilder:validation:Optional
	DiskIDRef *v1.Reference `json:"diskIdRef,omitempty" tf:"-"`

	// Selector for a Disk in instance to populate diskId.
	// +kubebuilder:validation:Optional
	DiskIDSelector *v1.Selector `json:"diskIdSelector,omitempty" tf:"-"`

	// The MD5 hash of the file to be uploaded. This is used to trigger file updates.
	// The MD5 hash of the image file.
	// +kubebuilder:validation:Optional
	FileHash *string `json:"fileHash,omitempty" tf:"file_hash,omitempty"`

	// The path of the image file to be uploaded.
	// The name of the file to upload to this image.
	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// A short description of the Image. Labels cannot contain special characters.
	// A short description of the Image. Labels cannot contain special characters.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the Linode that this Image will be created from.
	// The ID of the Linode that this Image will be created from.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/instance/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	LinodeID *float64 `json:"linodeId,omitempty" tf:"linode_id,omitempty"`

	// Reference to a Instance in instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDRef *v1.Reference `json:"linodeIdRef,omitempty" tf:"-"`

	// Selector for a Instance in instance to populate linodeId.
	// +kubebuilder:validation:Optional
	LinodeIDSelector *v1.Selector `json:"linodeIdSelector,omitempty" tf:"-"`

	// The region of the image. See all regions here.
	// The region to upload to.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Manages a Linode Image.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.label)",message="label is a required parameter"
	Spec   ImageSpec   `json:"spec"`
	Status ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
