// +build !ignore_autogenerated

// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package taxonomy

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	in.AdditionalProperties.DeepCopyInto(&out.AdditionalProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppInfo) DeepCopyInto(out *AppInfo) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppInfo.
func (in *AppInfo) DeepCopy() *AppInfo {
	if in == nil {
		return nil
	}
	out := new(AppInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
	in.AdditionalProperties.DeepCopyInto(&out.AdditionalProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Decision) DeepCopyInto(out *Decision) {
	*out = *in
	in.DeploymentRestrictions.DeepCopyInto(&out.DeploymentRestrictions)
	out.Policy = in.Policy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Decision.
func (in *Decision) DeepCopy() *Decision {
	if in == nil {
		return nil
	}
	out := new(Decision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DecisionPerCapabilityMap) DeepCopyInto(out *DecisionPerCapabilityMap) {
	{
		in := &in
		*out = make(DecisionPerCapabilityMap, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecisionPerCapabilityMap.
func (in DecisionPerCapabilityMap) DeepCopy() DecisionPerCapabilityMap {
	if in == nil {
		return nil
	}
	out := new(DecisionPerCapabilityMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecisionPolicy) DeepCopyInto(out *DecisionPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecisionPolicy.
func (in *DecisionPolicy) DeepCopy() *DecisionPolicy {
	if in == nil {
		return nil
	}
	out := new(DecisionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluationOutputStructure) DeepCopyInto(out *EvaluationOutputStructure) {
	*out = *in
	out.Config = in.Config.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluationOutputStructure.
func (in *EvaluationOutputStructure) DeepCopy() *EvaluationOutputStructure {
	if in == nil {
		return nil
	}
	out := new(EvaluationOutputStructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interface) DeepCopyInto(out *Interface) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Interface) DeepCopy() *Interface {
	if in == nil {
		return nil
	}
	out := new(Interface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyManagerRequestContext) DeepCopyInto(out *PolicyManagerRequestContext) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyManagerRequestContext.
func (in *PolicyManagerRequestContext) DeepCopy() *PolicyManagerRequestContext {
	if in == nil {
		return nil
	}
	out := new(PolicyManagerRequestContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restrictions) DeepCopyInto(out *Restrictions) {
	*out = *in
	out.Clusters = in.Clusters.DeepCopy()
	out.Modules = in.Modules.DeepCopy()
	out.StorageAccounts = in.StorageAccounts.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restrictions.
func (in *Restrictions) DeepCopy() *Restrictions {
	if in == nil {
		return nil
	}
	out := new(Restrictions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tags) DeepCopyInto(out *Tags) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in *Tags) DeepCopy() *Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return out
}
