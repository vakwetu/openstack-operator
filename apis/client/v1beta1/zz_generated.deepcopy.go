//go:build !ignore_autogenerated

/*
Copyright 2022.

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClient) DeepCopyInto(out *OpenStackClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClient.
func (in *OpenStackClient) DeepCopy() *OpenStackClient {
	if in == nil {
		return nil
	}
	out := new(OpenStackClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientDefaults) DeepCopyInto(out *OpenStackClientDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientDefaults.
func (in *OpenStackClientDefaults) DeepCopy() *OpenStackClientDefaults {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientList) DeepCopyInto(out *OpenStackClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientList.
func (in *OpenStackClientList) DeepCopy() *OpenStackClientList {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientSpec) DeepCopyInto(out *OpenStackClientSpec) {
	*out = *in
	in.OpenStackClientSpecCore.DeepCopyInto(&out.OpenStackClientSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientSpec.
func (in *OpenStackClientSpec) DeepCopy() *OpenStackClientSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientSpecCore) DeepCopyInto(out *OpenStackClientSpecCore) {
	*out = *in
	if in.OpenStackConfigMap != nil {
		in, out := &in.OpenStackConfigMap, &out.OpenStackConfigMap
		*out = new(string)
		**out = **in
	}
	if in.OpenStackConfigSecret != nil {
		in, out := &in.OpenStackConfigSecret, &out.OpenStackConfigSecret
		*out = new(string)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	out.Ca = in.Ca
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientSpecCore.
func (in *OpenStackClientSpecCore) DeepCopy() *OpenStackClientSpecCore {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientStatus) DeepCopyInto(out *OpenStackClientStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientStatus.
func (in *OpenStackClientStatus) DeepCopy() *OpenStackClientStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientStatus)
	in.DeepCopyInto(out)
	return out
}
