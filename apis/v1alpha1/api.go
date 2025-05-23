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

// ApiSpec defines the desired state of Api.
//
// Represents an API.
type APISpec struct {

	// An API key selection expression. Supported only for WebSocket APIs. See API
	// Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty"`
	// Specifies how to interpret the base path of the API during import. Valid
	// values are ignore, prepend, and split. The default value is ignore. To learn
	// more, see Set the OpenAPI basePath Property (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html).
	// Supported only for HTTP APIs.
	Basepath *string `json:"basepath,omitempty"`
	// The OpenAPI definition. Supported only for HTTP APIs.
	Body *string `json:"body,omitempty"`
	// A CORS configuration. Supported only for HTTP APIs. See Configuring CORS
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
	// for more information.
	CORSConfiguration *CORS `json:"corsConfiguration,omitempty"`
	// This property is part of quick create. It specifies the credentials required
	// for the integration, if any. For a Lambda integration, three options are
	// available. To specify an IAM Role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To require that the caller's identity be passed
	// through from the request, specify arn:aws:iam::*:user/*. To use resource-based
	// permissions on supported AWS services, specify null. Currently, this property
	// is not used for HTTP integrations. Supported only for HTTP APIs.
	CredentialsARN *string `json:"credentialsARN,omitempty"`
	// The description of the API.
	Description *string `json:"description,omitempty"`
	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com
	// endpoint. To require that clients use a custom domain name to invoke your
	// API, disable the default endpoint.
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteAPIEndpoint,omitempty"`
	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation *bool `json:"disableSchemaValidation,omitempty"`
	// Specifies whether to rollback the API creation when a warning is encountered.
	// By default, API creation continues if a warning is encountered.
	FailOnWarnings *bool `json:"failOnWarnings,omitempty"`
	// The name of the API.
	Name *string `json:"name,omitempty"`
	// The API protocol.
	ProtocolType *string `json:"protocolType,omitempty"`
	// This property is part of quick create. If you don't specify a routeKey, a
	// default route of $default is created. The $default route acts as a catch-all
	// for any request made to your API, for a particular stage. The $default route
	// key can't be modified. You can add routes after creating the API, and you
	// can update the route keys of additional routes. Supported only for HTTP APIs.
	RouteKey *string `json:"routeKey,omitempty"`
	// The route selection expression for the API. For HTTP APIs, the routeSelectionExpression
	// must be ${request.method} ${request.path}. If not provided, this will be
	// the default for HTTP APIs. This property is required for WebSocket APIs.
	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty"`
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string `json:"tags,omitempty"`
	// This property is part of quick create. Quick create produces an API with
	// an integration, a default catch-all route, and a default stage which is configured
	// to automatically deploy changes. For HTTP integrations, specify a fully qualified
	// URL. For Lambda integrations, specify a function ARN. The type of the integration
	// will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	Target *string `json:"target,omitempty"`
	// A version identifier for the API.
	Version *string `json:"version,omitempty"`
}

// APIStatus defines the observed state of API
type APIStatus struct {
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
	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com.
	// The stage name is typically appended to this URI to form a complete path
	// to a deployed API stage.
	// +kubebuilder:validation:Optional
	APIEndpoint *string `json:"apiEndpoint,omitempty"`
	// Specifies whether an API is managed by API Gateway. You can't update or delete
	// a managed API by using API Gateway. A managed API can be deleted only through
	// the tooling or service that created it.
	// +kubebuilder:validation:Optional
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`
	// The API ID.
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiID,omitempty"`
	// The timestamp when the API was created.
	// +kubebuilder:validation:Optional
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	// +kubebuilder:validation:Optional
	ImportInfo []*string `json:"importInfo,omitempty"`
	// The warning messages reported when failonwarnings is turned on during API
	// import.
	// +kubebuilder:validation:Optional
	Warnings []*string `json:"warnings,omitempty"`
}

// API is the Schema for the APIS API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec,omitempty"`
	Status            APIStatus `json:"status,omitempty"`
}

// APIList contains a list of API
// +kubebuilder:object:root=true
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
