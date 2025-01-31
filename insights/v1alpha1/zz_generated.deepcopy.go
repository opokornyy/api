//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataGather) DeepCopyInto(out *DataGather) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataGather.
func (in *DataGather) DeepCopy() *DataGather {
	if in == nil {
		return nil
	}
	out := new(DataGather)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataGather) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataGatherList) DeepCopyInto(out *DataGatherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataGather, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataGatherList.
func (in *DataGatherList) DeepCopy() *DataGatherList {
	if in == nil {
		return nil
	}
	out := new(DataGatherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataGatherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataGatherSpec) DeepCopyInto(out *DataGatherSpec) {
	*out = *in
	if in.Gatherers != nil {
		in, out := &in.Gatherers, &out.Gatherers
		*out = make([]GathererConfig, len(*in))
		copy(*out, *in)
	}
	out.StorageSpec = in.StorageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataGatherSpec.
func (in *DataGatherSpec) DeepCopy() *DataGatherSpec {
	if in == nil {
		return nil
	}
	out := new(DataGatherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataGatherStatus) DeepCopyInto(out *DataGatherStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Gatherers != nil {
		in, out := &in.Gatherers, &out.Gatherers
		*out = make([]GathererStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.FinishTime.DeepCopyInto(&out.FinishTime)
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	in.InsightsReport.DeepCopyInto(&out.InsightsReport)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataGatherStatus.
func (in *DataGatherStatus) DeepCopy() *DataGatherStatus {
	if in == nil {
		return nil
	}
	out := new(DataGatherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GathererConfig) DeepCopyInto(out *GathererConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GathererConfig.
func (in *GathererConfig) DeepCopy() *GathererConfig {
	if in == nil {
		return nil
	}
	out := new(GathererConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GathererStatus) DeepCopyInto(out *GathererStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LastGatherDuration = in.LastGatherDuration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GathererStatus.
func (in *GathererStatus) DeepCopy() *GathererStatus {
	if in == nil {
		return nil
	}
	out := new(GathererStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsightsReport) DeepCopyInto(out *InsightsReport) {
	*out = *in
	in.DownloadedAt.DeepCopyInto(&out.DownloadedAt)
	if in.HealthChecks != nil {
		in, out := &in.HealthChecks, &out.HealthChecks
		*out = make([]HealthCheck, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsightsReport.
func (in *InsightsReport) DeepCopy() *InsightsReport {
	if in == nil {
		return nil
	}
	out := new(InsightsReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
