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

type ObjectInitParameters struct {

	// The canned ACL to apply. (private, public-read, authenticated-read, public-read-write, custom) (defaults to private).
	// The ACL config given to this object.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The REQUIRED access key to authenticate with. If it's not specified with the resource, you must provide its value by
	// The REQUIRED S3 access key with access to the target bucket. If not specified with the resource, you must provide its value by configuring the obj_access_key, or, opting-in generating it implicitly at apply-time using obj_use_temp_keys at provider-level.
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:refFieldName=AccessKeyRef
	// +crossplane:generate:reference:selectorFieldName=AccessKeySelector
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Reference to a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeyRef *v1.Reference `json:"accessKeyRef,omitempty" tf:"-"`

	// Selector for a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeySelector *v1.Selector `json:"accessKeySelector,omitempty" tf:"-"`

	// The name of the bucket to put the object in.
	// The target bucket to put this object in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	// This cache_control configuration of this object.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// (Deprecated) The cluster the bucket is in. Required if region is not configured. Deprecated in favor of region.
	// The target cluster that the bucket is in.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	// The contents of the Object to upload.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	// The base64 contents of the Object to upload.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	// The content disposition configuration of this object.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// The encoding of the content of this object.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	// The language metadata of this object.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	// The MIME type of the content.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used with the s3 client to make bucket changes and will be computed automatically if left blank, override for testing/debug purposes.
	// The endpoint for the bucket used for s3 connections.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	// The specific version of this object.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted regardless of any legal hold or object lock (defaults to false).
	// Whether the object should bypass deletion restrictions.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// They name of the object once it is in the bucket.
	// The name of the uploaded object.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A map of keys/values to provision metadata.
	// The metadata of this object
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The cluster the bucket is in. Required if cluster is not configured.
	// The target region that the bucket is in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The REQUIRED secret key to authenticate with. If it's not specified with the resource, you must provide its value by
	// The REQUIRED S3 secret key with access to the target bucket. If not specified with the resource, you must provide its value by configuring the obj_secret_key, or, opting-in generating it implicitly at apply-time using obj_use_temp_keys at provider-level.
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`

	// The path to a file that will be read and uploaded as raw bytes for the object content. The path must either be relative to the root module or absolute.
	// The source file to upload.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies a target URL for website redirect.
	// The website redirect location of this object.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectObservation struct {

	// The canned ACL to apply. (private, public-read, authenticated-read, public-read-write, custom) (defaults to private).
	// The ACL config given to this object.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The REQUIRED access key to authenticate with. If it's not specified with the resource, you must provide its value by
	// The REQUIRED S3 access key with access to the target bucket. If not specified with the resource, you must provide its value by configuring the obj_access_key, or, opting-in generating it implicitly at apply-time using obj_use_temp_keys at provider-level.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// The name of the bucket to put the object in.
	// The target bucket to put this object in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	// This cache_control configuration of this object.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// (Deprecated) The cluster the bucket is in. Required if region is not configured. Deprecated in favor of region.
	// The target cluster that the bucket is in.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	// The contents of the Object to upload.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	// The base64 contents of the Object to upload.
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	// The content disposition configuration of this object.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// The encoding of the content of this object.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	// The language metadata of this object.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	// The MIME type of the content.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used with the s3 client to make bucket changes and will be computed automatically if left blank, override for testing/debug purposes.
	// The endpoint for the bucket used for s3 connections.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	// The specific version of this object.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted regardless of any legal hold or object lock (defaults to false).
	// Whether the object should bypass deletion restrictions.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// They name of the object once it is in the bucket.
	// The name of the uploaded object.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A map of keys/values to provision metadata.
	// The metadata of this object
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The cluster the bucket is in. Required if cluster is not configured.
	// The target region that the bucket is in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The path to a file that will be read and uploaded as raw bytes for the object content. The path must either be relative to the root module or absolute.
	// The source file to upload.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// A unique version ID value for the object.
	// The version ID of this object.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`

	// Specifies a target URL for website redirect.
	// The website redirect location of this object.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectParameters struct {

	// The canned ACL to apply. (private, public-read, authenticated-read, public-read-write, custom) (defaults to private).
	// The ACL config given to this object.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The REQUIRED access key to authenticate with. If it's not specified with the resource, you must provide its value by
	// The REQUIRED S3 access key with access to the target bucket. If not specified with the resource, you must provide its value by configuring the obj_access_key, or, opting-in generating it implicitly at apply-time using obj_use_temp_keys at provider-level.
	// +crossplane:generate:reference:type=Key
	// +crossplane:generate:reference:refFieldName=AccessKeyRef
	// +crossplane:generate:reference:selectorFieldName=AccessKeySelector
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Reference to a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeyRef *v1.Reference `json:"accessKeyRef,omitempty" tf:"-"`

	// Selector for a Key to populate accessKey.
	// +kubebuilder:validation:Optional
	AccessKeySelector *v1.Selector `json:"accessKeySelector,omitempty" tf:"-"`

	// The name of the bucket to put the object in.
	// The target bucket to put this object in.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	// This cache_control configuration of this object.
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// (Deprecated) The cluster the bucket is in. Required if region is not configured. Deprecated in favor of region.
	// The target cluster that the bucket is in.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	// The contents of the Object to upload.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the gzipbase64 function with small text strings. For larger objects, use source to stream the content from a disk file.
	// The base64 contents of the Object to upload.
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	// The content disposition configuration of this object.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// The encoding of the content of this object.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	// The language metadata of this object.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	// The MIME type of the content.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used with the s3 client to make bucket changes and will be computed automatically if left blank, override for testing/debug purposes.
	// The endpoint for the bucket used for s3 connections.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Used to trigger updates.11.11.11 or earlier).
	// The specific version of this object.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Allow the object to be deleted regardless of any legal hold or object lock (defaults to false).
	// Whether the object should bypass deletion restrictions.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// They name of the object once it is in the bucket.
	// The name of the uploaded object.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A map of keys/values to provision metadata.
	// The metadata of this object
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The cluster the bucket is in. Required if cluster is not configured.
	// The target region that the bucket is in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The REQUIRED secret key to authenticate with. If it's not specified with the resource, you must provide its value by
	// The REQUIRED S3 secret key with access to the target bucket. If not specified with the resource, you must provide its value by configuring the obj_secret_key, or, opting-in generating it implicitly at apply-time using obj_use_temp_keys at provider-level.
	// +kubebuilder:validation:Optional
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`

	// The path to a file that will be read and uploaded as raw bytes for the object content. The path must either be relative to the root module or absolute.
	// The source file to upload.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies a target URL for website redirect.
	// The website redirect location of this object.
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Object is the Schema for the Objects API. Manages a Linode Object Storage Object.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}
