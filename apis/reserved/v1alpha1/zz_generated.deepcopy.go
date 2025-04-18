//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignment) DeepCopyInto(out *IPAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignment.
func (in *IPAssignment) DeepCopy() *IPAssignment {
	if in == nil {
		return nil
	}
	out := new(IPAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentInitParameters) DeepCopyInto(out *IPAssignmentInitParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.LinodeIDRef != nil {
		in, out := &in.LinodeIDRef, &out.LinodeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LinodeIDSelector != nil {
		in, out := &in.LinodeIDSelector, &out.LinodeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.Rdns != nil {
		in, out := &in.Rdns, &out.Rdns
		*out = new(string)
		**out = **in
	}
	if in.RdnsRef != nil {
		in, out := &in.RdnsRef, &out.RdnsRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RdnsSelector != nil {
		in, out := &in.RdnsSelector, &out.RdnsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentInitParameters.
func (in *IPAssignmentInitParameters) DeepCopy() *IPAssignmentInitParameters {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentList) DeepCopyInto(out *IPAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentList.
func (in *IPAssignmentList) DeepCopy() *IPAssignmentList {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentObservation) DeepCopyInto(out *IPAssignmentObservation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(float64)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.Rdns != nil {
		in, out := &in.Rdns, &out.Rdns
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Reserved != nil {
		in, out := &in.Reserved, &out.Reserved
		*out = new(bool)
		**out = **in
	}
	if in.SubnetMask != nil {
		in, out := &in.SubnetMask, &out.SubnetMask
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VPCNAT11 != nil {
		in, out := &in.VPCNAT11, &out.VPCNAT11
		*out = make([]VPCNAT11Observation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentObservation.
func (in *IPAssignmentObservation) DeepCopy() *IPAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentParameters) DeepCopyInto(out *IPAssignmentParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.LinodeIDRef != nil {
		in, out := &in.LinodeIDRef, &out.LinodeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LinodeIDSelector != nil {
		in, out := &in.LinodeIDSelector, &out.LinodeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.Rdns != nil {
		in, out := &in.Rdns, &out.Rdns
		*out = new(string)
		**out = **in
	}
	if in.RdnsRef != nil {
		in, out := &in.RdnsRef, &out.RdnsRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RdnsSelector != nil {
		in, out := &in.RdnsSelector, &out.RdnsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentParameters.
func (in *IPAssignmentParameters) DeepCopy() *IPAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentSpec) DeepCopyInto(out *IPAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentSpec.
func (in *IPAssignmentSpec) DeepCopy() *IPAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAssignmentStatus) DeepCopyInto(out *IPAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAssignmentStatus.
func (in *IPAssignmentStatus) DeepCopy() *IPAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(IPAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCNAT11InitParameters) DeepCopyInto(out *VPCNAT11InitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCNAT11InitParameters.
func (in *VPCNAT11InitParameters) DeepCopy() *VPCNAT11InitParameters {
	if in == nil {
		return nil
	}
	out := new(VPCNAT11InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCNAT11Observation) DeepCopyInto(out *VPCNAT11Observation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(float64)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCNAT11Observation.
func (in *VPCNAT11Observation) DeepCopy() *VPCNAT11Observation {
	if in == nil {
		return nil
	}
	out := new(VPCNAT11Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCNAT11Parameters) DeepCopyInto(out *VPCNAT11Parameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCNAT11Parameters.
func (in *VPCNAT11Parameters) DeepCopy() *VPCNAT11Parameters {
	if in == nil {
		return nil
	}
	out := new(VPCNAT11Parameters)
	in.DeepCopyInto(out)
	return out
}
