// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentSpec defines the desired state of Deployment.
//
// An immutable representation of an API that can be called by users. A Deployment
// must be associated with a Stage for it to be callable over the internet.
type DeploymentSpec struct {

	// The API identifier.
	APIID  *string                                  `json:"apiID,omitempty"`
	APIRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"apiRef,omitempty"`
	// The description for the deployment resource.
	Description *string `json:"description,omitempty"`
	// The name of the Stage resource for the Deployment resource to create.
	StageName *string `json:"stageName,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment
type DeploymentStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Specifies whether a deployment was automatically released.
	// +kubebuilder:validation:Optional
	AutoDeployed *bool `json:"autoDeployed,omitempty"`
	// The date and time when the Deployment resource was created.
	// +kubebuilder:validation:Optional
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The identifier for the deployment.
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentID,omitempty"`
	// The status of the deployment: PENDING, FAILED, or SUCCEEDED.
	// +kubebuilder:validation:Optional
	DeploymentStatus *string `json:"deploymentStatus,omitempty"`
	// May contain additional feedback on the status of an API deployment.
	// +kubebuilder:validation:Optional
	DeploymentStatusMessage *string `json:"deploymentStatusMessage,omitempty"`
}

// Deployment is the Schema for the Deployments API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec,omitempty"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// DeploymentList contains a list of Deployment
// +kubebuilder:object:root=true
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
