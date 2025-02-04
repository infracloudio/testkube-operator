//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactRequest) DeepCopyInto(out *ArtifactRequest) {
	*out = *in
	if in.Dirs != nil {
		in, out := &in.Dirs, &out.Dirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactRequest.
func (in *ArtifactRequest) DeepCopy() *ArtifactRequest {
	if in == nil {
		return nil
	}
	out := new(ArtifactRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssertionResult) DeepCopyInto(out *AssertionResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssertionResult.
func (in *AssertionResult) DeepCopy() *AssertionResult {
	if in == nil {
		return nil
	}
	out := new(AssertionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvReference) DeepCopyInto(out *EnvReference) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvReference.
func (in *EnvReference) DeepCopy() *EnvReference {
	if in == nil {
		return nil
	}
	out := new(EnvReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Execution) DeepCopyInto(out *Execution) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]Variable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(TestContent)
		(*in).DeepCopyInto(*out)
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.ExecutionResult != nil {
		in, out := &in.ExecutionResult, &out.ExecutionResult
		*out = new(ExecutionResult)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Uploads != nil {
		in, out := &in.Uploads, &out.Uploads
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ArtifactRequest != nil {
		in, out := &in.ArtifactRequest, &out.ArtifactRequest
		*out = new(ArtifactRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.RunningContext != nil {
		in, out := &in.RunningContext, &out.RunningContext
		*out = new(RunningContext)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Execution.
func (in *Execution) DeepCopy() *Execution {
	if in == nil {
		return nil
	}
	out := new(Execution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionRequest) DeepCopyInto(out *ExecutionRequest) {
	*out = *in
	if in.ExecutionLabels != nil {
		in, out := &in.ExecutionLabels, &out.ExecutionLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]Variable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecretEnvs != nil {
		in, out := &in.SecretEnvs, &out.SecretEnvs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ArtifactRequest != nil {
		in, out := &in.ArtifactRequest, &out.ArtifactRequest
		*out = new(ArtifactRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvConfigMaps != nil {
		in, out := &in.EnvConfigMaps, &out.EnvConfigMaps
		*out = make([]EnvReference, len(*in))
		copy(*out, *in)
	}
	if in.EnvSecrets != nil {
		in, out := &in.EnvSecrets, &out.EnvSecrets
		*out = make([]EnvReference, len(*in))
		copy(*out, *in)
	}
	if in.RunningContext != nil {
		in, out := &in.RunningContext, &out.RunningContext
		*out = new(RunningContext)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionRequest.
func (in *ExecutionRequest) DeepCopy() *ExecutionRequest {
	if in == nil {
		return nil
	}
	out := new(ExecutionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionResult) DeepCopyInto(out *ExecutionResult) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ExecutionStatus)
		**out = **in
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]ExecutionStepResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reports != nil {
		in, out := &in.Reports, &out.Reports
		*out = new(ExecutionResultReports)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionResult.
func (in *ExecutionResult) DeepCopy() *ExecutionResult {
	if in == nil {
		return nil
	}
	out := new(ExecutionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionResultReports) DeepCopyInto(out *ExecutionResultReports) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionResultReports.
func (in *ExecutionResultReports) DeepCopy() *ExecutionResultReports {
	if in == nil {
		return nil
	}
	out := new(ExecutionResultReports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionStepResult) DeepCopyInto(out *ExecutionStepResult) {
	*out = *in
	if in.AssertionResults != nil {
		in, out := &in.AssertionResults, &out.AssertionResults
		*out = make([]AssertionResult, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionStepResult.
func (in *ExecutionStepResult) DeepCopy() *ExecutionStepResult {
	if in == nil {
		return nil
	}
	out := new(ExecutionStepResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	if in.UsernameSecret != nil {
		in, out := &in.UsernameSecret, &out.UsernameSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.TokenSecret != nil {
		in, out := &in.TokenSecret, &out.TokenSecret
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunningContext) DeepCopyInto(out *RunningContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunningContext.
func (in *RunningContext) DeepCopy() *RunningContext {
	if in == nil {
		return nil
	}
	out := new(RunningContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestContent) DeepCopyInto(out *TestContent) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(Repository)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestContent.
func (in *TestContent) DeepCopy() *TestContent {
	if in == nil {
		return nil
	}
	out := new(TestContent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecution) DeepCopyInto(out *TestExecution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecution.
func (in *TestExecution) DeepCopy() *TestExecution {
	if in == nil {
		return nil
	}
	out := new(TestExecution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestExecution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecutionList) DeepCopyInto(out *TestExecutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestExecution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecutionList.
func (in *TestExecutionList) DeepCopy() *TestExecutionList {
	if in == nil {
		return nil
	}
	out := new(TestExecutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestExecutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecutionSpec) DeepCopyInto(out *TestExecutionSpec) {
	*out = *in
	if in.Test != nil {
		in, out := &in.Test, &out.Test
		*out = new(ObjectRef)
		**out = **in
	}
	if in.ExecutionRequest != nil {
		in, out := &in.ExecutionRequest, &out.ExecutionRequest
		*out = new(ExecutionRequest)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecutionSpec.
func (in *TestExecutionSpec) DeepCopy() *TestExecutionSpec {
	if in == nil {
		return nil
	}
	out := new(TestExecutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestExecutionStatus) DeepCopyInto(out *TestExecutionStatus) {
	*out = *in
	if in.LatestExecution != nil {
		in, out := &in.LatestExecution, &out.LatestExecution
		*out = new(Execution)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestExecutionStatus.
func (in *TestExecutionStatus) DeepCopy() *TestExecutionStatus {
	if in == nil {
		return nil
	}
	out := new(TestExecutionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}
