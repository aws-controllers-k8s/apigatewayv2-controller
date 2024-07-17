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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Represents an API mapping.
type APIMapping struct {
	// The identifier.
	APIID *string `json:"apiID,omitempty"`
	// The identifier.
	APIMappingID *string `json:"apiMappingID,omitempty"`
	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	APIMappingKey *string `json:"apiMappingKey,omitempty"`
	// A string with a length between [1-128].
	Stage *string `json:"stage,omitempty"`
}

// Represents an API.
type API_SDK struct {
	APIEndpoint       *string `json:"apiEndpoint,omitempty"`
	APIGatewayManaged *bool   `json:"apiGatewayManaged,omitempty"`
	// The identifier.
	APIID *string `json:"apiID,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty"`
	// Represents a CORS configuration. Supported only for HTTP APIs. See Configuring
	// CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
	// for more information.
	CORSConfiguration *CORS        `json:"corsConfiguration,omitempty"`
	CreatedDate       *metav1.Time `json:"createdDate,omitempty"`
	// A string with a length between [0-1024].
	Description               *string   `json:"description,omitempty"`
	DisableExecuteAPIEndpoint *bool     `json:"disableExecuteAPIEndpoint,omitempty"`
	DisableSchemaValidation   *bool     `json:"disableSchemaValidation,omitempty"`
	ImportInfo                []*string `json:"importInfo,omitempty"`
	// A string with a length between [1-128].
	Name *string `json:"name,omitempty"`
	// Represents a protocol type.
	ProtocolType *string `json:"protocolType,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty"`
	// Represents a collection of tags associated with the resource.
	Tags map[string]*string `json:"tags,omitempty"`
	// A string with a length between [1-64].
	Version  *string   `json:"version,omitempty"`
	Warnings []*string `json:"warnings,omitempty"`
}

// Settings for logging access in a stage.
type AccessLogSettings struct {
	// Represents an Amazon Resource Name (ARN).
	DestinationARN *string `json:"destinationARN,omitempty"`
	// A string with a length between [1-1024].
	Format *string `json:"format,omitempty"`
}

// Represents an authorizer.
type Authorizer_SDK struct {
	// Represents an Amazon Resource Name (ARN).
	AuthorizerCredentialsARN *string `json:"authorizerCredentialsARN,omitempty"`
	// The identifier.
	AuthorizerID *string `json:"authorizerID,omitempty"`
	// A string with a length between [1-64].
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty"`
	// An integer with a value between [0-3600].
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTTLInSeconds,omitempty"`
	// The authorizer type. Specify REQUEST for a Lambda function using incoming
	// request parameters. Specify JWT to use JSON Web Tokens (supported only for
	// HTTP APIs).
	AuthorizerType *string `json:"authorizerType,omitempty"`
	// A string representation of a URI with a length between [1-2048].
	AuthorizerURI         *string `json:"authorizerURI,omitempty"`
	EnableSimpleResponses *bool   `json:"enableSimpleResponses,omitempty"`
	// The identity source for which authorization is requested. For the REQUEST
	// authorizer, this is required when authorization caching is enabled. The value
	// is a comma-separated string of one or more mapping expressions of the specified
	// request parameters. For example, if an Auth header, a Name query string parameter
	// are defined as identity sources, this value is $method.request.header.Auth,
	// $method.request.querystring.Name. These parameters will be used to derive
	// the authorization caching key and to perform runtime validation of the REQUEST
	// authorizer by verifying all of the identity-related request parameters are
	// present, not null and non-empty. Only when this is true does the authorizer
	// invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized
	// response without calling the Lambda function. The valid value is a string
	// of comma-separated mapping expressions of the specified request parameters.
	// When the authorization caching is not enabled, this property is optional.
	IdentitySource []*string `json:"identitySource,omitempty"`
	// A string with a length between [0-1024].
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty"`
	// Represents the configuration of a JWT authorizer. Required for the JWT authorizer
	// type. Supported only for HTTP APIs.
	JWTConfiguration *JWTConfiguration `json:"jwtConfiguration,omitempty"`
	// A string with a length between [1-128].
	Name *string `json:"name,omitempty"`
}

// Represents a CORS configuration. Supported only for HTTP APIs. See Configuring
// CORS (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html)
// for more information.
type CORS struct {
	AllowCredentials *bool `json:"allowCredentials,omitempty"`
	// Represents a collection of allowed headers. Supported only for HTTP APIs.
	AllowHeaders []*string `json:"allowHeaders,omitempty"`
	// Represents a collection of methods. Supported only for HTTP APIs.
	AllowMethods []*string `json:"allowMethods,omitempty"`
	// Represents a collection of origins. Supported only for HTTP APIs.
	AllowOrigins []*string `json:"allowOrigins,omitempty"`
	// Represents a collection of allowed headers. Supported only for HTTP APIs.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty"`
	// An integer with a value between -1 and 86400. Supported only for HTTP APIs.
	MaxAge *int64 `json:"maxAge,omitempty"`
}

// An immutable representation of an API that can be called by users. A Deployment
// must be associated with a Stage for it to be callable over the internet.
type Deployment_SDK struct {
	AutoDeployed *bool        `json:"autoDeployed,omitempty"`
	CreatedDate  *metav1.Time `json:"createdDate,omitempty"`
	// The identifier.
	DeploymentID *string `json:"deploymentID,omitempty"`
	// Represents a deployment status.
	DeploymentStatus        *string `json:"deploymentStatus,omitempty"`
	DeploymentStatusMessage *string `json:"deploymentStatusMessage,omitempty"`
	// A string with a length between [0-1024].
	Description *string `json:"description,omitempty"`
}

// The domain name configuration.
type DomainNameConfiguration struct {
	// Represents an Amazon Resource Name (ARN).
	CertificateARN *string `json:"certificateARN,omitempty"`
	// A string with a length between [1-128].
	CertificateName *string `json:"certificateName,omitempty"`
	// Represents an endpoint type.
	EndpointType *string `json:"endpointType,omitempty"`
	// Represents an Amazon Resource Name (ARN).
	OwnershipVerificationCertificateARN *string `json:"ownershipVerificationCertificateARN,omitempty"`
	// The Transport Layer Security (TLS) version of the security policy for this
	// domain name. The valid values are TLS_1_0 and TLS_1_2.
	SecurityPolicy *string `json:"securityPolicy,omitempty"`
}

// Represents a domain name.
type DomainName_SDK struct {
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	APIMappingSelectionExpression *string `json:"apiMappingSelectionExpression,omitempty"`
	// A string with a length between [1-512].
	DomainName *string `json:"domainName,omitempty"`
	// The domain name configurations.
	DomainNameConfigurations []*DomainNameConfiguration `json:"domainNameConfigurations,omitempty"`
	// If specified, API Gateway performs two-way authentication between the client
	// and the server. Clients must present a trusted certificate to access your
	// API.
	MutualTLSAuthentication *MutualTLSAuthentication `json:"mutualTLSAuthentication,omitempty"`
	// Represents a collection of tags associated with the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// Represents an integration response.
type IntegrationResponse struct {
	// Specifies how to handle response payload content type conversions. Supported
	// only for WebSocket APIs.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`
	// The identifier.
	IntegrationResponseID *string `json:"integrationResponseID,omitempty"`
	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	IntegrationResponseKey *string `json:"integrationResponseKey,omitempty"`
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
	ResponseParameters map[string]*string `json:"responseParameters,omitempty"`
	// A mapping of identifier keys to templates. The value is an actual template
	// script. The key is typically a SelectionKey which is chosen based on evaluating
	// a selection expression.
	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`
}

// Represents an integration.
type Integration_SDK struct {
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`
	// A string with a length between [1-1024].
	ConnectionID *string `json:"connectionID,omitempty"`
	// Represents a connection type.
	ConnectionType *string `json:"connectionType,omitempty"`
	// Specifies how to handle response payload content type conversions. Supported
	// only for WebSocket APIs.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`
	// Represents an Amazon Resource Name (ARN).
	CredentialsARN *string `json:"credentialsARN,omitempty"`
	// A string with a length between [0-1024].
	Description *string `json:"description,omitempty"`
	// The identifier.
	IntegrationID *string `json:"integrationID,omitempty"`
	// A string with a length between [1-64].
	IntegrationMethod *string `json:"integrationMethod,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty"`
	// A string with a length between [1-128].
	IntegrationSubtype *string `json:"integrationSubtype,omitempty"`
	// Represents an API method integration type.
	IntegrationType *string `json:"integrationType,omitempty"`
	// A string representation of a URI with a length between [1-2048].
	IntegrationURI *string `json:"integrationURI,omitempty"`
	// Represents passthrough behavior for an integration response. Supported only
	// for WebSocket APIs.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty"`
	// A string with a length between [1-64].
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
	// A mapping of identifier keys to templates. The value is an actual template
	// script. The key is typically a SelectionKey which is chosen based on evaluating
	// a selection expression.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty"`
	// Supported only for HTTP APIs. You use response parameters to transform the
	// HTTP response from a backend integration before returning the response to
	// clients.
	ResponseParameters map[string]map[string]*string `json:"responseParameters,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`
	// An integer with a value between [50-30000].
	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty"`
	// The TLS configuration for a private integration. If you specify a TLS configuration,
	// private integration traffic uses the HTTPS protocol. Supported only for HTTP
	// APIs.
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
}

// Represents the configuration of a JWT authorizer. Required for the JWT authorizer
// type. Supported only for HTTP APIs.
type JWTConfiguration struct {
	Audience []*string `json:"audience,omitempty"`
	// A string representation of a URI with a length between [1-2048].
	Issuer *string `json:"issuer,omitempty"`
}

// Represents a data model for an API. Supported only for WebSocket APIs. See
// Create Models and Mapping Templates for Request and Response Mappings (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html).
type Model struct {
	// A string with a length between [0-1024].
	Description *string `json:"description,omitempty"`
	// The identifier.
	ModelID *string `json:"modelID,omitempty"`
	// A string with a length between [1-128].
	Name *string `json:"name,omitempty"`
	// A string with a length between [0-32768].
	Schema *string `json:"schema,omitempty"`
}

// If specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthentication struct {
	// A string representation of a URI with a length between [1-2048].
	TruststoreURI *string `json:"truststoreURI,omitempty"`
	// A string with a length between [1-64].
	TruststoreVersion  *string   `json:"truststoreVersion,omitempty"`
	TruststoreWarnings []*string `json:"truststoreWarnings,omitempty"`
}

// If specified, API Gateway performs two-way authentication between the client
// and the server. Clients must present a trusted certificate to access your
// API.
type MutualTLSAuthenticationInput struct {
	// A string representation of a URI with a length between [1-2048].
	TruststoreURI *string `json:"truststoreURI,omitempty"`
	// A string with a length between [1-64].
	TruststoreVersion *string `json:"truststoreVersion,omitempty"`
}

// Validation constraints imposed on parameters of a request (path, query string,
// headers).
type ParameterConstraints struct {
	Required *bool `json:"required,omitempty"`
}

// Represents a route response.
type RouteResponse struct {
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty"`
	// The route models.
	ResponseModels map[string]*string `json:"responseModels,omitempty"`
	// The route parameters.
	ResponseParameters map[string]*ParameterConstraints `json:"responseParameters,omitempty"`
	// The identifier.
	RouteResponseID *string `json:"routeResponseID,omitempty"`
	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteResponseKey *string `json:"routeResponseKey,omitempty"`
}

// Represents a collection of route settings.
type RouteSettings struct {
	DataTraceEnabled       *bool `json:"dataTraceEnabled,omitempty"`
	DetailedMetricsEnabled *bool `json:"detailedMetricsEnabled,omitempty"`
	// The logging level.
	LoggingLevel         *string  `json:"loggingLevel,omitempty"`
	ThrottlingBurstLimit *int64   `json:"throttlingBurstLimit,omitempty"`
	ThrottlingRateLimit  *float64 `json:"throttlingRateLimit,omitempty"`
}

// Represents a route.
type Route_SDK struct {
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`
	APIKeyRequired    *bool `json:"apiKeyRequired,omitempty"`
	// A list of authorization scopes configured on a route. The scopes are used
	// with a JWT authorizer to authorize the method invocation. The authorization
	// works by matching the route scopes against the scopes parsed from the access
	// token in the incoming request. The method invocation is authorized if any
	// route scope matches a claimed scope in the access token. Otherwise, the invocation
	// is not authorized. When the route scope is configured, the client must provide
	// an access token instead of an identity token for authorization purposes.
	AuthorizationScopes []*string `json:"authorizationScopes,omitempty"`
	// The authorization type. For WebSocket APIs, valid values are NONE for open
	// access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda
	// authorizer. For HTTP APIs, valid values are NONE for open access, JWT for
	// using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM
	// for using a Lambda authorizer.
	AuthorizationType *string `json:"authorizationType,omitempty"`
	// The identifier.
	AuthorizerID *string `json:"authorizerID,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty"`
	// A string with a length between [1-64].
	OperationName *string `json:"operationName,omitempty"`
	// The route models.
	RequestModels map[string]*string `json:"requestModels,omitempty"`
	// The route parameters.
	RequestParameters map[string]*ParameterConstraints `json:"requestParameters,omitempty"`
	// The identifier.
	RouteID *string `json:"routeID,omitempty"`
	// After evaluating a selection expression, the result is compared against one
	// or more selection keys to find a matching key. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for a list of expressions and each expression's associated selection key
	// type.
	RouteKey *string `json:"routeKey,omitempty"`
	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty"`
	// A string with a length between [1-128].
	Target *string `json:"target,omitempty"`
}

// Represents an API stage.
type Stage_SDK struct {
	// Settings for logging access in a stage.
	AccessLogSettings *AccessLogSettings `json:"accessLogSettings,omitempty"`
	APIGatewayManaged *bool              `json:"apiGatewayManaged,omitempty"`
	AutoDeploy        *bool              `json:"autoDeploy,omitempty"`
	// The identifier.
	ClientCertificateID *string      `json:"clientCertificateID,omitempty"`
	CreatedDate         *metav1.Time `json:"createdDate,omitempty"`
	// Represents a collection of route settings.
	DefaultRouteSettings *RouteSettings `json:"defaultRouteSettings,omitempty"`
	// The identifier.
	DeploymentID *string `json:"deploymentID,omitempty"`
	// A string with a length between [0-1024].
	Description                 *string      `json:"description,omitempty"`
	LastDeploymentStatusMessage *string      `json:"lastDeploymentStatusMessage,omitempty"`
	LastUpdatedDate             *metav1.Time `json:"lastUpdatedDate,omitempty"`
	// The route settings map.
	RouteSettings map[string]*RouteSettings `json:"routeSettings,omitempty"`
	// A string with a length between [1-128].
	StageName *string `json:"stageName,omitempty"`
	// The stage variable map.
	StageVariables map[string]*string `json:"stageVariables,omitempty"`
	// Represents a collection of tags associated with the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// The TLS configuration for a private integration. If you specify a TLS configuration,
// private integration traffic uses the HTTPS protocol. Supported only for HTTP
// APIs.
type TLSConfig struct {
	// A string with a length between [1-512].
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty"`
}

// The TLS configuration for a private integration. If you specify a TLS configuration,
// private integration traffic uses the HTTPS protocol. Supported only for HTTP
// APIs.
type TLSConfigInput struct {
	// A string with a length between [1-512].
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty"`
}

// Represents a VPC link.
type VPCLink_SDK struct {
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	// A string with a length between [1-128].
	Name *string `json:"name,omitempty"`
	// A list of security group IDs for the VPC link.
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	// A list of subnet IDs to include in the VPC link.
	SubnetIDs []*string `json:"subnetIDs,omitempty"`
	// Represents a collection of tags associated with the resource.
	Tags map[string]*string `json:"tags,omitempty"`
	// The identifier.
	VPCLinkID *string `json:"vpcLinkID,omitempty"`
	// The status of the VPC link.
	VPCLinkStatus *string `json:"vpcLinkStatus,omitempty"`
	// A string with a length between [0-1024].
	VPCLinkStatusMessage *string `json:"vpcLinkStatusMessage,omitempty"`
	// The version of the VPC link.
	VPCLinkVersion *string `json:"vpcLinkVersion,omitempty"`
}
