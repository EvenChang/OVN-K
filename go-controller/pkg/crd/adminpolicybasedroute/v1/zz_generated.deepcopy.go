//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminPolicyBasedExternalRoute) DeepCopyInto(out *AdminPolicyBasedExternalRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminPolicyBasedExternalRoute.
func (in *AdminPolicyBasedExternalRoute) DeepCopy() *AdminPolicyBasedExternalRoute {
	if in == nil {
		return nil
	}
	out := new(AdminPolicyBasedExternalRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminPolicyBasedExternalRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminPolicyBasedExternalRouteList) DeepCopyInto(out *AdminPolicyBasedExternalRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminPolicyBasedExternalRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminPolicyBasedExternalRouteList.
func (in *AdminPolicyBasedExternalRouteList) DeepCopy() *AdminPolicyBasedExternalRouteList {
	if in == nil {
		return nil
	}
	out := new(AdminPolicyBasedExternalRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminPolicyBasedExternalRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminPolicyBasedExternalRouteSpec) DeepCopyInto(out *AdminPolicyBasedExternalRouteSpec) {
	*out = *in
	in.From.DeepCopyInto(&out.From)
	in.NextHops.DeepCopyInto(&out.NextHops)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminPolicyBasedExternalRouteSpec.
func (in *AdminPolicyBasedExternalRouteSpec) DeepCopy() *AdminPolicyBasedExternalRouteSpec {
	if in == nil {
		return nil
	}
	out := new(AdminPolicyBasedExternalRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminPolicyBasedRouteStatus) DeepCopyInto(out *AdminPolicyBasedRouteStatus) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminPolicyBasedRouteStatus.
func (in *AdminPolicyBasedRouteStatus) DeepCopy() *AdminPolicyBasedRouteStatus {
	if in == nil {
		return nil
	}
	out := new(AdminPolicyBasedRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicHop) DeepCopyInto(out *DynamicHop) {
	*out = *in
	in.PodSelector.DeepCopyInto(&out.PodSelector)
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicHop.
func (in *DynamicHop) DeepCopy() *DynamicHop {
	if in == nil {
		return nil
	}
	out := new(DynamicHop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNetworkSource) DeepCopyInto(out *ExternalNetworkSource) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNetworkSource.
func (in *ExternalNetworkSource) DeepCopy() *ExternalNetworkSource {
	if in == nil {
		return nil
	}
	out := new(ExternalNetworkSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalNextHops) DeepCopyInto(out *ExternalNextHops) {
	*out = *in
	if in.StaticHops != nil {
		in, out := &in.StaticHops, &out.StaticHops
		*out = make([]*StaticHop, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StaticHop)
				**out = **in
			}
		}
	}
	if in.DynamicHops != nil {
		in, out := &in.DynamicHops, &out.DynamicHops
		*out = make([]*DynamicHop, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DynamicHop)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalNextHops.
func (in *ExternalNextHops) DeepCopy() *ExternalNextHops {
	if in == nil {
		return nil
	}
	out := new(ExternalNextHops)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticHop) DeepCopyInto(out *StaticHop) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticHop.
func (in *StaticHop) DeepCopy() *StaticHop {
	if in == nil {
		return nil
	}
	out := new(StaticHop)
	in.DeepCopyInto(out)
	return out
}
