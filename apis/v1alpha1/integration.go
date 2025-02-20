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

// IntegrationSpec defines the desired state of Integration.
//
// Represents an integration.
type IntegrationSpec struct {

	// The API identifier.
	APIID  *string                                  `json:"apiID,omitempty"`
	APIRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"apiRef,omitempty"`
	// The ID of the VPC link for a private integration. Supported only for HTTP
	// APIs.
	ConnectionID  *string                                  `json:"connectionID,omitempty"`
	ConnectionRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"connectionRef,omitempty"`
	// The type of the network connection to the integration endpoint. Specify INTERNET
	// for connections through the public routable internet or VPC_LINK for private
	// connections between API Gateway and resources in a VPC. The default value
	// is INTERNET.
	ConnectionType *string `json:"connectionType,omitempty"`
	// Supported only for WebSocket APIs. Specifies how to handle response payload
	// content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT,
	// with the following behaviors:
	//
	// CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string
	// to the corresponding binary blob.
	//
	// CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded
	// string.
	//
	// If this property is not defined, the response payload will be passed through
	// from the integration response to the route response or method response without
	// modification.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`
	// Specifies the credentials required for the integration, if any. For AWS integrations,
	// three options are available. To specify an IAM Role for API Gateway to assume,
	// use the role's Amazon Resource Name (ARN). To require that the caller's identity
	// be passed through from the request, specify the string arn:aws:iam::*:user/*.
	// To use resource-based permissions on supported AWS services, specify null.
	CredentialsARN *string `json:"credentialsARN,omitempty"`
	// The description of the integration.
	Description *string `json:"description,omitempty"`
	// Specifies the integration's HTTP method type.
	IntegrationMethod *string `json:"integrationMethod,omitempty"`
	// Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service
	// action to invoke. To learn more, see Integration subtype reference (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html).
	IntegrationSubtype *string `json:"integrationSubtype,omitempty"`
	// The integration type of an integration. One of the following:
	//
	// AWS: for integrating the route or method request with an AWS service action,
	// including the Lambda function-invoking action. With the Lambda function-invoking
	// action, this is referred to as the Lambda custom integration. With any other
	// AWS service action, this is known as AWS integration. Supported only for
	// WebSocket APIs.
	//
	// AWS_PROXY: for integrating the route or method request with a Lambda function
	// or other AWS service action. This integration is also referred to as a Lambda
	// proxy integration.
	//
	// HTTP: for integrating the route or method request with an HTTP endpoint.
	// This integration is also referred to as the HTTP custom integration. Supported
	// only for WebSocket APIs.
	//
	// HTTP_PROXY: for integrating the route or method request with an HTTP endpoint,
	// with the client request passed through as-is. This is also referred to as
	// HTTP proxy integration. For HTTP API private integrations, use an HTTP_PROXY
	// integration.
	//
	// MOCK: for integrating the route or method request with API Gateway as a "loopback"
	// endpoint without invoking any backend. Supported only for WebSocket APIs.
	// +kubebuilder:validation:Required
	IntegrationType *string `json:"integrationType"`
	// For a Lambda integration, specify the URI of a Lambda function.
	//
	// For an HTTP integration, specify a fully-qualified URL.
	//
	// For an HTTP API private integration, specify the ARN of an Application Load
	// Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	// If you specify the ARN of an AWS Cloud Map service, API Gateway uses DiscoverInstances
	// to identify resources. You can use query parameters to target specific resources.
	// To learn more, see DiscoverInstances (https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html).
	// For private integrations, all resources must be owned by the same AWS account.
	IntegrationURI *string `json:"integrationURI,omitempty"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type
	// header in the request, and the available mapping templates specified as the
	// requestTemplates property on the Integration resource. There are three valid
	// values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported only for WebSocket
	// APIs.
	//
	// WHEN_NO_MATCH passes the request body for unmapped content types through
	// to the integration backend without transformation.
	//
	// NEVER rejects unmapped content types with an HTTP 415 Unsupported Media Type
	// response.
	//
	// WHEN_NO_TEMPLATES allows pass-through when the integration has no content
	// types mapped to templates. However, if there is at least one content type
	// defined, unmapped content types will be rejected with the same HTTP 415 Unsupported
	// Media Type response.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty"`
	// Specifies the format of the payload sent to an integration. Required for
	// HTTP APIs.
	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty"`
	// For WebSocket APIs, a key-value map specifying request parameters that are
	// passed from the method request to the backend. The key is an integration
	// request parameter name and the associated value is a method request parameter
	// value or static value that must be enclosed within single quotes and pre-encoded
	// as required by the backend. The method request parameter value must match
	// the pattern of method.request.{location}.{name} , where {location} is querystring,
	// path, or header; and {name} must be a valid and unique method request parameter
	// name.
	//
	// For HTTP API integrations with a specified integrationSubtype, request parameters
	// are a key-value map specifying parameters that are passed to AWS_PROXY integrations.
	// You can provide static values, or map request data, stage variables, or context
	// variables that are evaluated at runtime. To learn more, see Working with
	// AWS service integrations for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services.html).
	//
	// For HTTP API integrations without a specified integrationSubtype request
	// parameters are a key-value map specifying how to transform HTTP requests
	// before sending them to the backend. The key should follow the pattern <action>:<header|querystring|path>.<location>
	// where action can be append, overwrite or remove. For values, you can provide
	// static values, or map request data, stage variables, or context variables
	// that are evaluated at runtime. To learn more, see Transforming API requests
	// and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).
	RequestParameters map[string]*string `json:"requestParameters,omitempty"`
	// Represents a map of Velocity templates that are applied on the request payload
	// based on the value of the Content-Type header sent by the client. The content
	// type value is the key in this map, and the template (as a String) is the
	// value. Supported only for WebSocket APIs.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty"`
	// Supported only for HTTP APIs. You use response parameters to transform the
	// HTTP response from a backend integration before returning the response to
	// clients. Specify a key-value map from a selection key to response parameters.
	// The selection key must be a valid HTTP status code within the range of 200-599.
	// Response parameters are a key-value map. The key must match pattern <action>:<header>.<location>
	// or overwrite.statuscode. The action can be append, overwrite or remove. The
	// value can be a static value, or map to response data, stage variables, or
	// context variables that are evaluated at runtime. To learn more, see Transforming
	// API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).
	ResponseParameters map[string]map[string]*string `json:"responseParameters,omitempty"`
	// The template selection expression for the integration.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`
	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
	// between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is
	// 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty"`
	// The TLS configuration for a private integration. If you specify a TLS configuration,
	// private integration traffic uses the HTTPS protocol. Supported only for HTTP
	// APIs.
	TLSConfig *TLSConfigInput `json:"tlsConfig,omitempty"`
}

// IntegrationStatus defines the observed state of Integration
type IntegrationStatus struct {
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
	// Specifies whether an integration is managed by API Gateway. If you created
	// an API using using quick create, the resulting integration is managed by
	// API Gateway. You can update a managed integration, but you can't delete it.
	// +kubebuilder:validation:Optional
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`
	// Represents the identifier of an integration.
	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationID,omitempty"`
	// The integration response selection expression for the integration. Supported
	// only for WebSocket APIs. See Integration Response Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions).
	// +kubebuilder:validation:Optional
	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty"`
}

// Integration is the Schema for the Integrations API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationSpec   `json:"spec,omitempty"`
	Status            IntegrationStatus `json:"status,omitempty"`
}

// IntegrationList contains a list of Integration
// +kubebuilder:object:root=true
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
