// +build !ignore_autogenerated

/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBinding) DeepCopyInto(out *ProfileBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBinding.
func (in *ProfileBinding) DeepCopy() *ProfileBinding {
	if in == nil {
		return nil
	}
	out := new(ProfileBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBindingList) DeepCopyInto(out *ProfileBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProfileBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBindingList.
func (in *ProfileBindingList) DeepCopy() *ProfileBindingList {
	if in == nil {
		return nil
	}
	out := new(ProfileBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBindingSpec) DeepCopyInto(out *ProfileBindingSpec) {
	*out = *in
	out.ProfileRef = in.ProfileRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBindingSpec.
func (in *ProfileBindingSpec) DeepCopy() *ProfileBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ProfileBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileBindingStatus) DeepCopyInto(out *ProfileBindingStatus) {
	*out = *in
	if in.ActiveWorkloads != nil {
		in, out := &in.ActiveWorkloads, &out.ActiveWorkloads
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileBindingStatus.
func (in *ProfileBindingStatus) DeepCopy() *ProfileBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ProfileBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileRef) DeepCopyInto(out *ProfileRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileRef.
func (in *ProfileRef) DeepCopy() *ProfileRef {
	if in == nil {
		return nil
	}
	out := new(ProfileRef)
	in.DeepCopyInto(out)
	return out
}