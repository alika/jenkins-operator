/*
Copyright 2018 Samsung SDS Cloud Native Computing Team.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the JenkinsInstance resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JenkinsInstanceSpec defines the desired state of JenkinsInstance
type JenkinsInstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// JenkinsInstanceStatus defines the observed state of JenkinsInstance
type JenkinsInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JenkinsInstance
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=jenkinsinstances
type JenkinsInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JenkinsInstanceSpec   `json:"spec,omitempty"`
	Status JenkinsInstanceStatus `json:"status,omitempty"`
}
